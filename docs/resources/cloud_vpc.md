---
page_title: "Cloud VPC Attachment - Utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Attach or detach a VPC subnet from a Utho Cloud instance.
---

# utho_cloud_vpc

Attaches or detaches a VPC subnet from an existing cloud instance. Use this to add an instance to a private network post-deployment without recreating it.

~> **Note:** You can also attach a VPC subnet at creation time using the `vpc` argument in `utho_cloud`. Use `utho_cloud_vpc` to change the VPC attachment on an already-running instance.

## Example Usage

### Attach instance to a VPC subnet

```hcl
resource "utho_cloud_vpc" "attach" {
  cloud_id  = utho_cloud.web.id
  vpc_id    = utho_subnet.private.id
}
```

### Use existing subnet from data source

```hcl
data "utho_vpcs" "mumbai" {
  dcslug = "inmumbaizone2"
}

locals {
  private_subnet = one(flatten([
    for vpc in data.utho_vpcs.mumbai.vpcs : [
      for s in vpc.subnets :
      s if s.subnet_type == "private" && vpc.name == "production"
    ]
  ]))
}

resource "utho_cloud_vpc" "attach" {
  cloud_id = utho_cloud.backend.id
  vpc_id   = local.private_subnet.id
}
```

### Move instance to different subnet

```hcl
# Detach from old subnet — destroy the existing utho_cloud_vpc resource
# Then create a new one pointing to the new subnet

resource "utho_cloud_vpc" "new_subnet" {
  cloud_id = utho_cloud.web.id
  vpc_id   = utho_subnet.new_private.id
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `cloud_id` | String | Yes      | Cloud instance ID. Changing this forces a new resource. |
| `vpc_id`   | String | Yes      | VPC subnet numeric ID. Use [utho_vpcs](../data-sources/vpcs) or [utho_vpc_subnets](../data-sources/vpc_subnets) to find valid subnet IDs. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Attachment ID. |

## Notes

- Use the subnet's numeric `id` — not the UUID.
- The instance and subnet must be in the same data center.
- Destroying this resource detaches the instance from the VPC subnet.
- Use `data.utho_vpcs` or `data.utho_vpc_subnets` to discover existing subnet IDs.
