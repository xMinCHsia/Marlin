
# Marlin

> Speculative decoding orchestrator for LLM serving: a draft-model ensemble
> with rejection sampling, KV-cache-aware admission control and
> self-tuning draft lengths - the substrate for faster, cheaper inference.

![Go](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-blue?style=flat-square)
![Deps](https://img.shields.io/badge/dependencies-1-orange?style=flat-square)
![Platform](https://img.shields.io/badge/platform-linux%20%7C%20macOS-lightgrey?style=flat-square)

![Marlin request flow](assets/marlin-flow.svg)

## Table of contents

- [What Marlin actually does](#what-marlin-actually-does)
- [Features](#features)
- [Quick start](#quick-start)
- [How it works](#how-it-works)
- [Configuration](#configuration)
- [Operations](#operations)
- [FAQ](#faq)
- [License](#license)

## What Marlin actually does

Speculative decoding is a bet: most tokens are predictable, so a cheap
draft model can propose several, and the expensive target only verifies.
Marlin orchestrates that bet across a *draft ensemble* - multiple draft
models with different strengths - instead of a single draft model.

Three subsystems make the bet stay profitable:

1. **The ensemble** - draft candidates are sampled from several models and
   merged into a proposal batch. Composition is a tuning parameter.
2. **Rejection statistics** - acceptance rate per draft model over a
   sliding window, with *drift detection*: a draft that used to mirror
   the target and suddenly does not.
3. **KV-cache admission control** - drafts consume the same KV-cache the
   target needs; Marlin prices each proposal and admits drafts only while
   the target engine stays under its memory budget.

## Features

| | |
|---|---|
| 🎛 **Ensemble composition** | weight-split budget allocation across drafts, capped per draft |
| 🎯 **Rejection sampling** | per-token verdicts with sliding-window statistics |
| 📉 **Drift detection** | baseline-vs-window acceptance drop flags stale drafts |
| 🧠 **Autotuner** | draft lengths move with acceptance, one step per probe |
| 💾 **Cache admission** | exponential backoff when the target cache budget is tight |
| 📋 **Deterministic plans** | every request leaves a hashed, replayable plan |

## Quick start

```console
$ make build
$ marlind -config marlin.yaml.example

        marlind v2.1.0
        target : http://127.0.0.1:8000 (8 GiB kv budget)
        drafts : draft-small(8) draft-code(6)
        listening on 127.0.0.1:8590
```

Then from another shell:

```console
$ marlinctl -addr localhost:8590 status
200 {
  "status": "ok",
  "drafts": 2,
  "pressure": 0.41
}
```

## How it works

1. A request arrives with a token budget.
2. `ensemble.Compose` splits the budget across draft models by weight.
3. `admission.Admit` prices the total KV footprint; rejected drafts back
   off exponentially.
4. The target verifies proposals; verdicts are tallied per draft.
5. The acceptance window updates; drift over 15% shrinks that draft.
6. The whole exchange lands in a hashed, replayable plan.

## Configuration

`marlin.yaml.example` ships in the repo root. The important knobs:

| Key | Default | Meaning |
|---|---|---|
| `target.max_kv_bytes` | 8 GiB | cache budget reserved for the target |
| `drafts[].weight` | 1.0 | share of the proposal budget |
| `drafts[].max_proposal_len` | 8 | ceiling for that draft's proposals |
| `admission.headroom_ratio` | 0.2 | safety margin kept free |
| `tuning.enabled` | true | let draft lengths self-adjust |

## Operations

