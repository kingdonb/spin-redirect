# spin-redirect

This is a simple HTTP redirect component written in Go.

This is not a Spin application in itself but a component that can be used in applications to redirect a route.

## Configuration

The `spin-redirect` component can be configured to address your needs in different scenarios. `spin-redirect` tries to load configuration data from multiple places in the specified order:

1. Spin component configuration
2. Environment variables

The following table outlines available configuration values:

| Key            | Description                                                                                                             | Default Value |
|----------------|-------------------------------------------------------------------------------------------------------------------------|---------------|
| `destination`  | Where should the component redirect to                                                                                  | `/`           |
| `statuscode`   | What HTTP status code should be used when redirecting                                                                   | `302`         |
| `include_path` | Whether to include the original request path on the destination redirect; see [wildcard redirects](#wildcard-redirects) | `false`       |
| `trim_prefix`  | A specific prefix portion of the original request path to trim (when `include_path` is `true`)                          |               |

The `spin-redirect` component tries to look up the config value in the Spin component configuration using the keys shown in the table above (lower case). If desired key is not present, it transforms the key to upper case (e.g., `DESTINATION`) and checks environment variables.

### Valid redirection status codes

The `spin-redirect` component supports the following HTTP status codes to perform a redirect:

- `301` Moved Permanently
- `302` Found (Moved Temporarily)
- `303` See Other: Only supported for `PUT` and `POST` requests
- `307` Temporary Redirect
- `308` Permanent Redirect

## Example usage

The following snippet shows how to add and configure `spin-redirect` in your `spin.toml` using environment variables

```toml
spin_manifest_version = "1"
description = ""
name = "test"
trigger = { type = "http", base = "/" }
version = "0.1.0"

# Redirect / to /index.html using HTTP status code 301
[[component]]
id = "redirect-sample"
source = "path/to/redirect.wasm"
environment = { DESTINATION = "/index.html", STATUSCODE = "301" }

[component.trigger]
route = "/"
```

Alternatively, you can use component configuration to configure `spin-redirect` as shown below:

```toml
spin_manifest_version = "1"
description = ""
name = "test"
trigger = { type = "http", base = "/" }
version = "0.1.0"

# Redirect / to /index.html using HTTP status code 301
[[component]]
id = "redirect-sample"
source = "path/to/redirect.wasm"

[component.config]
destination="/index.html"
statuscode="301"

[component.trigger]
route = "/"

```

### Wildcard redirects

Wildcard redirects can be enabled by setting `include_path` to `true`.  In the example below,
all requests to the application will be redirected to `/foo`, with the original request path included,
e.g. `/bar` will redirect to `/foo/bar`, `/baz` will redirect to `/foo/baz` and so on.

```toml
spin_manifest_version = 2

[application]
name = "wildcard-redirect"
version = "0.1.0"

[[trigger.http]]
id = "trigger-redirect-to-foo"
component = "redirect-to-foo"
route = "/..."

[component.redirect-to-foo]
source = "modules/redirect.wasm"
environment = { DESTINATION = "/foo", INCLUDE_PATH = "true" }
```

A prefix portion of the original request path can be trimmed via the `trim_prefix` configuration.

In this example, the `/v1/` prefix is trimmed, such that requests to `/v1/bar` will redirect to `/v2/bar` and so on.

```toml
spin_manifest_version = 2

[application]
name = "wildcard-redirect"
version = "0.1.0"

[[trigger.http]]
id = "trigger-redirect-v1-to-v2"
component = "redirect-v1-to-v2"
route = "/v1/..."

[component.redirect-v1-to-v2]
source = "modules/redirect.wasm"
environment = { DESTINATION = "/v2", INCLUDE_PATH = "true", TRIM_PREFIX = "/v1/" }
```
