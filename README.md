
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
