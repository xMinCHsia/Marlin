
# Contributing

Marlin is inference infrastructure - changes shift real latency and
memory on the serving path.

- `make test` before every push (mirrors CI: build + race tests).
- Stdlib only in this module; the YAML parser is the only dependency.
