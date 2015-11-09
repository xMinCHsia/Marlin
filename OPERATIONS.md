
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

