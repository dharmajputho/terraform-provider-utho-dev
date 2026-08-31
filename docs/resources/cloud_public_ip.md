---
page_title: "Cloud Instance Public IP - Utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Assign or release an additional public IP address on a Utho Cloud instance.
---

# utho_cloud_public_ip

Assigns an additional public IPv4 address to a Utho Cloud instance.
Each instance can have multiple public IPs. The IP is automatically
configured on the instance within a few seconds.

Destroying this resource releases the IP address.

## Example Usage

### Assign an additional public IP

```hcl
resource "utho_cloud_public_ip" "extra" {
  cloud_id = utho_cloud.web.id
}

output "extra_ip" {
  value = utho_cloud_public_ip.extra.ip
}
```

### Assign multiple public IPs to the same instance

```hcl
resource "utho_cloud_public_ip" "ip_1" {
  cloud_id = utho_cloud.web.id
}

resource "utho_cloud_public_ip" "ip_2" {
  cloud_id = utho_cloud.web.id
}

output "additional_ips" {
  value = [
    utho_cloud_public_ip.ip_1.ip,
    utho_cloud_public_ip.ip_2.ip,
  ]
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `cloud_id` | String | Yes      | The ID of the cloud instance. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cloud_id}:{ip}`. |
| `ip`      | String | The additional public IPv4 address assigned by Utho. |

## Notes

- The IP is assigned automatically — a specific address cannot be requested.
- Network configuration completes within 30–60 seconds of assignment.
- Releasing the IP takes 30–40 seconds to fully propagate.
- This resource manages **additional** IPs only. The primary IP is managed by `utho_cloud`.
