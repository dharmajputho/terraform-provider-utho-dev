---
page_title: "Load Balancer - Utho"
subcategory: "Networking / Load Balancing"
description: |-
  Create and manage Utho Load Balancers.
---

# utho_loadbalancer

Creates and manages a Utho Load Balancer. Supports both `network` and `application` types. Manage frontends using `utho_loadbalancer_frontend`, backends using `utho_loadbalancer_backend`, and ACL rules using `utho_loadbalancer_acl`.

## Example Usage

### Basic network load balancer

```hcl
resource "utho_loadbalancer" "main" {
  name            = "main-lb"
  type            = "network"
  dcslug          = "inmumbaizone2"
  enable_publicip = "true"
}

output "lb_ip" {
  value = utho_loadbalancer.main.ip
}
```

### Application load balancer inside a VPC

```hcl
resource "utho_loadbalancer" "app" {
  name            = "app-lb"
  type            = "application"
  dcslug          = "inmumbaizone2"
  vpc             = utho_subnet.public.id
  enable_publicip = "true"
  firewall        = utho_firewall.web.id
}
```

### Full stack — LB with frontend and backends

```hcl
resource "utho_loadbalancer" "main" {
  name            = "main-lb"
  type            = "network"
  dcslug          = "inmumbaizone2"
  enable_publicip = "true"
}

resource "utho_loadbalancer_frontend" "http" {
  loadbalancer_id = utho_loadbalancer.main.id
  name            = "http"
  algorithm       = "roundrobin"
  proto           = "http"
  port            = "80"
  cookie          = "0"
  redirecthttps   = "0"
  certificate_id  = "0"
}

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

## Argument Reference

| Argument          | Type   | Required | Description |
|-------------------|--------|----------|-------------|
| `name`            | String | Yes      | Load balancer name. |
| `type`            | String | Yes      | Load balancer type: `network` or `application`. Changing this forces a new resource. |
| `dcslug`          | String | Yes      | Data center slug. Changing this forces a new resource. |
| `enable_publicip` | String | Yes      | Assign a public IP: `true` or `false`. |
| `vpc`             | String | No       | Subnet ID to deploy inside a VPC. Changing this forces a new resource. |
| `firewall`        | String | No       | Security group ID to attach. Changing this forces a new resource. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique load balancer ID. |
| `ip`         | String | Public IP address. |
| `dns`        | String | DNS hostname. |
| `status`     | String | Current status. |
| `created_at` | String | Creation timestamp. |
