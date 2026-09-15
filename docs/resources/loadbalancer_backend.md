---
page_title: "Load Balancer Backend - Utho"
subcategory: "Networking / Load Balancing"
description: |-
  Add and manage backends on a Utho Load Balancer frontend.
---

# utho_loadbalancer_backend

Adds a backend server to a Utho Load Balancer frontend. Backends can be Utho cloud instances or custom IP addresses.

## Example Usage

### Add cloud instances as backends

```hcl
resource "utho_loadbalancer_backend" "app" {
  count           = 3
  loadbalancer_id = utho_loadbalancer.main.id
  frontend_id     = utho_loadbalancer_frontend.http.id
  backend_port    = "80"
  weight          = "1"
  type            = "cloud"
  cloudid         = utho_cloud.app[count.index].id
}
```

### Add a custom IP backend

```hcl
resource "utho_loadbalancer_backend" "external" {
  loadbalancer_id = utho_loadbalancer.main.id
  frontend_id     = utho_loadbalancer_frontend.http.id
  backend_port    = "80"
  weight          = "1"
  type            = "custom"
  ip              = "10.0.0.42"
}
```

## Argument Reference

| Argument          | Type   | Required | Description |
|-------------------|--------|----------|-------------|
| `loadbalancer_id` | String | Yes      | Load balancer ID. Changing this forces a new resource. |
| `frontend_id`     | String | Yes      | Frontend ID to attach this backend to. Changing this forces a new resource. |
| `backend_port`    | String | Yes      | Port the backend listens on. Changing this forces a new resource. |
| `weight`          | String | Yes      | Backend weight for load distribution (e.g. `1`). Changing this forces a new resource. |
| `type`            | String | Yes      | Backend type: `cloud` or `custom`. Changing this forces a new resource. |
| `cloudid`         | String | No       | Cloud instance ID. Required when `type = cloud`. Changing this forces a new resource. |
| `ip`              | String | No       | Custom IP address. Required when `type = custom`. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique backend ID. |
