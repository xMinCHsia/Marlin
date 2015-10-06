
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
{
  "request_id": "r-42",
  "target": "http://127.0.0.1:8000",
  "budget": 12,
  "steps": [
    { "draft_id": "draft-small", "tokens": ["the","quick"], "accepted": 2, "len": 2 }
  ],
  "summary_hash": "9f3ab2c1d4e5f607"
}
```

`accepted` is filled by the target verification pass. The hash covers the
ordered steps + request id, so identical inputs produce identical plans.

## Draft endpoint contract

Drafts expose `POST /propose` returning:

