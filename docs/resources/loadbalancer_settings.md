---
page_title: "Load Balancer Settings - Utho"
subcategory: "Networking / Load Balancing"
description: |-
  Configure advanced settings for a Utho Load Balancer.
---

# utho_loadbalancer_settings

Configures advanced timeout, connection, and performance settings for a Utho Load Balancer.

## Example Usage

```hcl
resource "utho_loadbalancer_settings" "main" {
  loadbalancer_id        = utho_loadbalancer.main.id
  timeout_connect        = "5"
  timeout_client         = "60"
  timeout_server         = "60"
  timeout_http_request   = "10"
  timeout_http_keepalive = "15"
  timeout_tunnel         = "3600"
  max_connections        = "10000"
  http2                  = "1"
  compression            = "1"
  hsts                   = "1"
}
```

## Argument Reference

| Argument                 | Type   | Required | Description |
|--------------------------|--------|----------|-------------|
| `loadbalancer_id`        | String | Yes      | Load balancer ID. Changing this forces a new resource. |
| `timeout_connect`        | String | No       | Connection timeout in seconds. Default: `5`. |
| `timeout_client`         | String | No       | Client inactivity timeout in seconds. Default: `60`. |
| `timeout_server`         | String | No       | Server response timeout in seconds. Default: `60`. |
| `timeout_http_request`   | String | No       | HTTP request timeout in seconds. Default: `10`. |
| `timeout_http_keepalive` | String | No       | HTTP keep-alive timeout in seconds. Default: `15`. |
| `timeout_tunnel`         | String | No       | Tunnel timeout in seconds (for WebSocket/TCP). Default: `3600`. |
| `max_connections`        | String | No       | Maximum concurrent connections. Default: `10000`. |
| `http2`                  | String | No       | Enable HTTP/2: `1` or `0`. |
| `compression`            | String | No       | Enable response compression: `1` or `0`. |
| `hsts`                   | String | No       | Enable HTTP Strict Transport Security: `1` or `0`. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Same as `loadbalancer_id`. |
