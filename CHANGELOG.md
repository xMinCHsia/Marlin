
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
