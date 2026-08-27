---
page_title: "utho_cloud_snapshot Resource - utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Create, restore, and delete snapshots of a Utho Cloud instance.
---

# utho_cloud_snapshot

Creates a point-in-time snapshot of a Utho Cloud instance. Snapshots
capture the full disk state and can be used to restore the instance
or deploy new instances from that state.

~> **Note:** The Utho API does not return the snapshot ID in the create
response. After creation, retrieve the `snapshot_id` from the Utho dashboard
(**Cloud Instance → Snapshot tab**) and add it to your configuration
to enable restore and delete operations.

## Example Usage

### Create a snapshot before a deployment

```hcl
resource "utho_cloud_snapshot" "before_deploy" {
  cloud_id = utho_cloud.web.id
  name     = "before-v2-deploy"
}
```

### Restore a snapshot

After creation, note the snapshot ID from the dashboard, then set `restore = true`:

```hcl
resource "utho_cloud_snapshot" "before_deploy" {
  cloud_id    = utho_cloud.web.id
  name        = "before-v2-deploy"
  snapshot_id = "201852"
  restore     = true
}
```

### Delete a snapshot

Add `snapshot_id` then remove the resource block and run `terraform apply`:

```hcl
# Remove this block and run: terraform apply
resource "utho_cloud_snapshot" "before_deploy" {
  cloud_id    = utho_cloud.web.id
  name        = "before-v2-deploy"
  snapshot_id = "201852"
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `cloud_id`    | String | Yes      | The ID of the cloud instance to snapshot. Changing this forces a new resource. |
| `name`        | String | Yes      | Name of the snapshot. Changing this forces a new resource. |
| `snapshot_id` | String | No       | Snapshot ID. Retrieved from the Utho dashboard after creation. Required for restore and delete. |
| `restore`     | Bool   | No       | Set to `true` to restore the instance from this snapshot. Requires `snapshot_id` to be set. |

## Attribute Reference

| Attribute     | Type   | Description |
|---------------|--------|-------------|
| `id`          | String | Identifier in the format `{cloud_id}:{name}`. |
| `snapshot_id` | String | Snapshot ID set after retrieval from the dashboard. |

## Workflow

1. Create the resource — Terraform sends the snapshot creation request.
2. Wait for the snapshot to complete in the Utho dashboard.
3. Note the snapshot ID from **Cloud Instance → Snapshot tab**.
4. Add `snapshot_id` to your `.tf` file.
5. To restore: set `restore = true` and run `terraform apply`.
6. To delete: remove the resource block and run `terraform apply`.

~> **Warning:** Restoring a snapshot overwrites all current data on
the instance. This operation cannot be undone.
