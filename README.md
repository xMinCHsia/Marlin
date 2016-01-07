
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
