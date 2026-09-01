---
page_title: "Elastic IP - Utho"
subcategory: "Networking / VPC"
description: |-
  Allocate and manage Elastic IPs on Utho Cloud.
---

# utho_elastic_ip

Allocates and manages an Elastic IP address on Utho Cloud. An Elastic IP
is a static public IP that persists independently of any specific cloud
instance. Use it with NAT Gateways or assign it to instances directly.

Destroying this resource releases the Elastic IP back to Utho.

## Example Usage

### Allocate an Elastic IP

```hcl
resource "utho_elastic_ip" "main" {
  dcslug       = "inmumbaizone2"
  billingcycle = "monthly"
}

output "elastic_ip" {
  value       = utho_elastic_ip.main.ip
  description = "Allocated elastic IP address"
}
```

### Use with a NAT Gateway

```hcl
resource "utho_elastic_ip" "nat_ip" {
  dcslug       = "inmumbaizone2"
  billingcycle = "monthly"
}

resource "utho_nat_gateway" "main" {
  name      = "main-nat"
  subnet_id = utho_subnet.public.id
  public_ip = utho_elastic_ip.nat_ip.ip
  dcslug    = "inmumbaizone2"
}
```

### Multiple Elastic IPs

```hcl
resource "utho_elastic_ip" "ips" {
  count        = 3
  dcslug       = "inmumbaizone2"
  billingcycle = "monthly"
}

output "ip_addresses" {
  value = utho_elastic_ip.ips[*].ip
}
```

## Argument Reference

| Argument       | Type   | Required | Description |
|----------------|--------|----------|-------------|
| `dcslug`       | String | Yes      | Data center slug where the IP is allocated. Changing this forces a new resource. |
| `billingcycle` | String | Yes      | Billing cycle. Accepted values: `monthly`, `hourly`. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Elastic IP identifier. |
| `ip`      | String | Allocated public IP address. |

## Notes

- The IP address is assigned automatically — a specific address cannot be requested.
- Elastic IPs are billed even when not attached to any resource.
- Changing `dcslug` or `billingcycle` releases the current IP and allocates a new one.
- Always detach the Elastic IP from any resource before releasing it.
