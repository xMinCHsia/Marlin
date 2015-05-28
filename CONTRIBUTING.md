
# Contributing

Marlin is inference infrastructure - changes shift real latency and
memory on the serving path.

- `make test` before every push (mirrors CI: build + race tests).
- Stdlib only in this module; the YAML parser is the only dependency.
- One logical change per commit; `area: change` present-tense messages.
- API or plan-format changes require: CHANGELOG entry, PROTOCOL update
