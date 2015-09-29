
# Marlin Protocol

## Orchestration surface

```
GET  /health                       # liveness + pressure
GET  /v1/status                    # ensemble summary
GET  /v1/drafts                    # per-draft acceptance stats
GET  /v1/plans                     # recent plan hashes
POST /v1/tune                      # force a tuning pass
```

## Plan lifecycle

A request produces one plan:

```
POST /v1/plans
{"request_id":"r-42","budget":12}
```

Response:

```json
