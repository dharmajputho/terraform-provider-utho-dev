---
page_title: "Cloud Instance Security Group - Utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Attach or detach a security group from a Utho Cloud instance.
---

# utho_cloud_firewall

Attaches an existing security group (firewall) to a Utho Cloud instance.
Multiple security groups can be attached to the same instance by creating
multiple `utho_cloud_firewall` resources.

Destroying this resource detaches the security group. Rules stop applying
within a few seconds of detachment.

## Example Usage

### Attach a single security group

```hcl
resource "utho_cloud_firewall" "web_sg" {
  cloud_id    = utho_cloud.web.id
  firewall_id = "23437038"
}
```

### Attach multiple security groups to the same instance

```hcl
resource "utho_cloud_firewall" "web_sg" {
  cloud_id    = utho_cloud.web.id
  firewall_id = "23437038"
}

resource "utho_cloud_firewall" "monitoring_sg" {
  cloud_id    = utho_cloud.web.id
  firewall_id = "23436544"
}
```

### Attach the same security group to multiple instances

```hcl
resource "utho_cloud" "workers" {
  count    = 3
  hostname = "worker-${count.index + 1}.mhc"
  # ...
}

resource "utho_cloud_firewall" "worker_sg" {
  count       = 3
  cloud_id    = utho_cloud.workers[count.index].id
  firewall_id = "23437038"
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `cloud_id`    | String | Yes      | The ID of the cloud instance. Changing this forces a new resource. |
| `firewall_id` | String | Yes      | The security group ID to attach. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{firewall_id}:{cloud_id}`. |

## Notes

- Changing either `cloud_id` or `firewall_id` destroys and recreates the attachment.
- To detach a security group, remove the resource block and run `terraform apply`.
- Security group rules apply within a few seconds of attachment.
- The security group itself is not deleted when this resource is destroyed.
