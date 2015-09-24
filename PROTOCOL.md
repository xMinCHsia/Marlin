
# Marlin Protocol

## Orchestration surface

```
GET  /health                       # liveness + pressure
GET  /v1/status                    # ensemble summary
GET  /v1/drafts                    # per-draft acceptance stats
GET  /v1/plans                     # recent plan hashes
