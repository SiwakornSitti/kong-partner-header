# Kong Partner Header (Go)

A Kong Gateway Go PDK plugin that sets the upstream request header `Partner` to the configured `partner_value`. It does not inspect or validate `x-devops-api-key`.

Requires self-managed Kong Gateway 3.4 or later. Kong's Go PDK runs as a plugin-server executable.

## Build

```sh
make build
```

For a local Kong container, build a Linux executable matching its architecture:

```sh
make build-kong
```

Override the architecture if needed, for example `make build-kong KONG_GOARCH=amd64`.

Install the binary in the Kong Gateway image or filesystem. Add the plugin server to `kong.conf`:

```ini
pluginserver_names = partner-header
pluginserver_partner_header_query_cmd = /usr/local/bin/partner-header -dump
plugins = bundled,partner-header
```

Restart or roll out Kong after installation. Enable `partner-header` on the intended Route or Service with the `partner_value` configuration field. Set the real partner identifier there; `XXX` is only a placeholder.
