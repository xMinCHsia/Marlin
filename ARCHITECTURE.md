
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

