---
page_title: "utho_cloud_ebs Resource - utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Attach and manage EBS volumes on a Utho Cloud instance.
---

# utho_cloud_ebs

Attaches an existing EBS (Elastic Block Storage) volume to a Utho Cloud
instance. The EBS volume must already exist in your Utho account.

Destroying this resource detaches the EBS volume. The volume itself is
not deleted.

## Example Usage

### Attach an EBS volume

```hcl
resource "utho_cloud_ebs" "data" {
  cloud_id = utho_cloud.app.id
  ebs_id   = "85829"
}
```

### Attach with bus and type configuration

```hcl
resource "utho_cloud_ebs" "data" {
  cloud_id  = utho_cloud.app.id
  ebs_id    = "85829"
  bus       = "virtio"
  disk_type = "Additional"
}
```

### Update bus and type after attachment

```hcl
resource "utho_cloud_ebs" "data" {
  cloud_id  = utho_cloud.app.id
  ebs_id    = "85829"
  bus       = "ide"
  disk_type = "Primary"
}
```

### Attach and power-cycle to detect new disk

```hcl
resource "utho_cloud_ebs" "data" {
  cloud_id = utho_cloud.app.id
  ebs_id   = "85829"
}

resource "utho_cloud_power" "apply_disk" {
  cloud_id   = utho_cloud.app.id
  action     = "powercycle"
  depends_on = [utho_cloud_ebs.data]
}
```

## Argument Reference

| Argument    | Type   | Required | Description |
|-------------|--------|----------|-------------|
| `cloud_id`  | String | Yes      | The ID of the cloud instance. Changing this forces a new resource. |
| `ebs_id`    | String | Yes      | The ID of the EBS volume to attach. Changing this forces a new resource. |
| `bus`       | String | No       | Disk bus type. Accepted values: `virtio`, `ide`. |
| `disk_type` | String | No       | Disk role. Accepted values: `Primary`, `Additional`. Only one `Primary` disk allowed per instance. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{ebs_id}:{cloud_id}`. |

## Notes

- Changing `cloud_id` or `ebs_id` destroys the current attachment and creates a new one.
- Changing `bus` or `disk_type` updates the existing attachment in place.
- The EBS volume is **not** deleted when this resource is destroyed — only the attachment is removed.
- After attaching, power-cycle the instance using `utho_cloud_power` so the guest OS detects the new disk.
