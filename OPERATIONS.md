
# Marlin Operations

## Deployment

One static binary plus one YAML document:

```console
$ make build
$ install -m 0755 bin/marlind /usr/local/bin/marlind
$ install -m 0644 marlin.yaml.example /etc/marlin/marlin.yaml
$ systemctl start marlin
```

The daemon talks to the target and draft endpoints over plain HTTP inside
the trusted network.

## Health checks

```
GET /health
{"status":"ok","drafts":2,"pressure":0.41}
```

Wire `/health` to the LB. `pressure` is the fraction of the target cache
budget currently reserved by drafts - a sustained value over 0.9 means
admission is blocking most proposals.

## Operator commands

```console
$ marlinctl -addr localhost:8590 status      # drafts + pressure + rates
$ marlinctl -addr localhost:8590 drafts      # per-draft acceptance stats
$ marlinctl -addr localhost:8590 plans       # recent plan hashes
$ marlinctl -addr localhost:8590 tune        # force a tuning pass
```

## Tuning the ensemble

- `weight` controls budget split. Start with equal weights.
- `max_proposal_len` is the ceiling per draft - the autotuner moves the
  effective length between `min_draft_len` and this cap.
- After a target model swap, reset statistics (restart the daemon) - the
  old acceptance baselines are meaningless.

## Monitoring

Export per-draft: acceptance rate (window + baseline), tuned length,
admission rejections. Alert on:

- acceptance rate below 0.25 for 10 minutes (draft is dead weight)
- pressure above 0.9 (cache starvation)
- autotuner pinned at `max_draft_len` for an hour (proposals too long)
