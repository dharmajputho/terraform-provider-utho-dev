---
page_title: "utho_cloud_public_ip Resource - utho"
subcategory: "Networking"
description: |-
  Assign or release an additional public IP address on a Utho Cloud instance.
---

# utho_cloud_public_ip

Assigns an additional public IPv4 address to a Utho Cloud instance.
Each instance can have multiple public IPs. The assigned IP is
automatically configured on the instance within a few seconds.

Destroying this resource releases the IP address.

## Example Usage

### Assign an additional public IP

```hcl
resource "utho_cloud_public_ip" "extra" {
  cloud_id = utho_cloud.web.id
}

output "extra_ip" {
  value       = utho_cloud_public_ip.extra.ip
  description = "Additional public IP assigned to the instance"
}
```

### Assign multiple public IPs

```hcl
resource "utho_cloud_public_ip" "ip_1" {
  cloud_id = utho_cloud.web.id
}

resource "utho_cloud_public_ip" "ip_2" {
  cloud_id = utho_cloud.web.id
}

output "all_ips" {
  value = [
    utho_cloud_public_ip.ip_1.ip,
    utho_cloud_public_ip.ip_2.ip,
  ]
}
```

### Use assigned IP in a DNS record

```hcl
resource "utho_cloud_public_ip" "api_ip" {
  cloud_id = utho_cloud.api.id
}

# Pass the IP to a DNS resource
output "api_ip_address" {
  value = utho_cloud_public_ip.api_ip.ip
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `cloud_id` | String | Yes      | The ID of the cloud instance to assign the IP to. Changing this forces a new resource. |

## Attribute Reference

| Attribute  | Type   | Description |
|------------|--------|-------------|
| `id`       | String | Identifier in the format `{cloud_id}:{ip}`. |
| `ip`       | String | The public IPv4 address assigned by Utho. |

## Notes

- The IP is assigned automatically by Utho — you cannot specify a
  particular IP address.
- Network configuration on the instance completes within 30–60 seconds
  of assignment.
- Releasing the IP (destroying the resource) takes 30–40 seconds to
  fully propagate.
