
# Marlin Architecture

## Positioning

Marlin is an orchestrator, not an inference engine. It sits in front of a
target LLM endpoint and a set of draft endpoints, decides *what* to propose,
*when* to admit a proposal, and *how long* the proposals should be - then
lets the target verify.

## The economics of speculation

Speculative decoding wins when the draft's acceptance rate is high and the
draft is cheap. Two failure modes ruin the bet:

1. **Acceptance drift** - a draft model that used to mirror the target
   (fine-tuned checkpoints drift apart) stops being accepted. Marlin
   detects this from a sliding acceptance window and backs the draft off.
2. **Cache pressure** - each draft proposal occupies KV-cache slots that
   the target engine would otherwise use for real work. When the target
   engine runs out of cache, it recomputes (or worse, degrades). Marlin
   prices every proposal's footprint and admits drafts only under the
   headroom budget.

## Components

| Component | Responsibility |
|---|---|
| `marlind` | orchestrator daemon, HTTP surface |
| `marlinctl` | operator CLI |
| `internal/admission` | KV-cache-aware admission control |
| `internal/draft` | ensemble composition + sampling |
| `internal/accept` | rejection statistics + drift detection |
| `internal/kvcache` | cache memory tracker |
| `internal/tuning` | draft-length autotuner |
| `pkg/plan` | deterministic execution plans |

## Request flow

1. A request arrives with a token budget (`budget`).
2. `ensemble.Compose` splits the budget across draft models by weight,
   respecting each draft's `max_proposal_len`.
3. `admission.Admit` prices the total footprint; rejected drafts back
   off exponentially and the request proceeds with fewer candidates.
4. The target engine verifies proposals; `reject.Record` tallies verdicts
   per draft.
5. `statistics.Window` updates acceptance rates; drift over 15% triggers
   `tuning` to shrink that draft's length one step per probe interval.
6. The whole exchange is captured as a `plan` with a deterministic hash
   for replay and auditing.

