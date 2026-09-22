---
page_title: "Load Balancers - Utho"
subcategory: "Networking / Load Balancing"
description: |-
  List all load balancers in your Utho account.
---

# utho_loadbalancers

Fetches all load balancers in your account. Use this to discover existing load balancer IDs, IPs, and DNS names without importing them into Terraform state.

## Example Usage

### List all load balancers

```hcl
data "utho_loadbalancers" "all" {}

output "loadbalancers" {
  value = [
    for lb in data.utho_loadbalancers.all.loadbalancers :
    { id = lb.id, name = lb.name, ip = lb.ip, dns = lb.dns }
  ]
}
```

### Find a load balancer by name

```hcl
data "utho_loadbalancers" "all" {}

locals {
  prod_lb = one([
    for lb in data.utho_loadbalancers.all.loadbalancers :
    lb if lb.name == "production-lb"
  ])
}

output "prod_lb_ip"  { value = local.prod_lb.ip }
output "prod_lb_dns" { value = local.prod_lb.dns }
```

### List active load balancers only

```hcl
data "utho_loadbalancers" "all" {}

output "active_lbs" {
  value = [
    for lb in data.utho_loadbalancers.all.loadbalancers :
    { id = lb.id, name = lb.name, type = lb.type, dc = lb.dcslug }
    if lb.status == "Active"
  ]
}
```

## Attribute Reference

### Top-level

| Attribute       | Type | Description |
|-----------------|------|-------------|
| `loadbalancers` | List | List of all load balancers in your account. |

### loadbalancers

| Attribute        | Type   | Description |
|------------------|--------|-------------|
| `id`             | String | Load balancer ID. Use this as `loadbalancer_id` in frontend and backend resources. |
| `name`           | String | Load balancer name. |
| `ip`             | String | Public IP address. Point your DNS A record here. |
| `dns`            | String | DNS hostname assigned to the load balancer. |
| `type`           | String | Load balancer type: `network` or `application`. |
| `dcslug`         | String | Data center slug. |
| `status`         | String | Current status (e.g. `Active`). |
| `enable_publicip`| String | Whether public IP is enabled: `true` or `false`. |
| `created_at`     | String | Creation timestamp. |