
# Changelog

All notable changes to Marlin are documented here.

## [Unreleased]

### Added
- config defaults applied on load for minimal configuration documents

### Changed
- sampler clamps negative scores via a shared helper
- ensemble hoists the weight total out of the composition loop

## [2.1.0] - 2026-03-12

### Added
- `marlinctl plans` with plan hash listing
- drift baseline snapshots per draft

### Changed
- admission backoff exponentiation bounds the wait by request timeout

## [2.0.0] - 2025-10-07

### Added
- draft-length autotuner (`internal/tuning`)
- per-draft statistics endpoint

### Changed
- proposal plans now carry a deterministic summary hash (breaking format)

## [1.2.0] - 2025-04-15

### Added
- sliding acceptance window with drift detection
- `POST /v1/tune` force pass

## [1.0.0] - 2024-06-20

### Added
- first stable orchestration surface
- ensemble composition with weight split

## [0.3.0] - 2023-11-08

### Added
- KV-cache admission control with exponential backoff
- `/v1/status` summary endpoint

## [0.2.0] - 2023-03-17

### Added
- rejection statistics and per-draft counters
- draft endpoint contract

## [0.1.0] - 2022-09-14

### Added
- initial orchestrator daemon and CLI scaffold
- proposal batch composition

<!-- draft note 791 -->
