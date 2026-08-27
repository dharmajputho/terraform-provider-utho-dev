---
page_title: "utho_cloud_storage Resource - utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Add and manage general storage disks on a Utho Cloud instance.
---

# utho_cloud_storage

Adds a general storage disk to a Utho Cloud instance. Multiple disks
can be added to the same instance by creating multiple `utho_cloud_storage`
resources.

~> **Note:** The Utho API does not return the disk ID in the create response.
After creation, retrieve the `disk_id` from the Utho dashboard
(**Cloud Instance → Storage tab**) and add it to your configuration
to enable update and delete operations.

## Example Usage

### Add an additional disk

```hcl
resource "utho_cloud_storage" "extra" {
  cloud_id = utho_cloud.web.id
  size_gb  = 50
}
```

### Add a disk with bus and type configuration

```hcl
resource "utho_cloud_storage" "data" {
  cloud_id  = utho_cloud.web.id
  size_gb   = 100
  bus       = "virtio"
  disk_type = "Additional"
}
```

### Update disk configuration after noting the disk_id

```hcl
resource "utho_cloud_storage" "data" {
  cloud_id  = utho_cloud.web.id
  size_gb   = 100
  disk_id   = "89153"    # retrieved from dashboard after creation
  bus       = "ide"
  disk_type = "Additional"
}
```

## Argument Reference

| Argument    | Type   | Required | Description |
|-------------|--------|----------|-------------|
| `cloud_id`  | String | Yes      | The ID of the cloud instance. Changing this forces a new resource. |
| `size_gb`   | Number | Yes      | Disk size in GB. Changing this forces a new resource. |
| `disk_id`   | String | No       | Disk ID assigned by Utho. Required for update and delete. Retrieve from **Cloud Instance → Storage tab** after creation. |
| `bus`       | String | No       | Disk bus type. Accepted values: `virtio`, `ide`. Default: `virtio`. |
| `disk_type` | String | No       | Disk role. Accepted values: `Primary`, `Additional`. Only one `Primary` disk allowed per instance. Default: `Additional`. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cloud_id}:{size_gb}`. |
| `disk_id` | String | Disk ID. Set manually after retrieving from the dashboard. |

## Workflow

1. Create the resource — Terraform adds the disk to the instance.
2. Note the disk ID from **Cloud Instance → Storage tab** in the Utho dashboard.
3. Add `disk_id` to your `.tf` file.
4. Run `terraform apply` to sync the state.
5. Update `bus` or `disk_type` as needed — Terraform calls the update API.
6. Remove the resource block and run `terraform apply` to delete the disk.
