---
page_title: "Cloud Instance Resize - Utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Resize a Utho Cloud instance — CPU/RAM only or full resize including disk.
---

# utho_cloud_resize

Resizes a Utho Cloud instance to a different plan. Supports two modes:
CPU/RAM only (`ramcpu`) and full resize including disk (`full`).

~> **Note:** Resize operations apply immediately. Disk downgrades are not
supported — use `ramcpu` when moving to a plan with a smaller disk allocation.

## Example Usage

### Upgrade CPU and RAM only

```hcl
resource "utho_cloud_resize" "scale_up" {
  cloud_id    = utho_cloud.app.id
  plan_id     = "10315"
  resize_type = "ramcpu"
}
```

### Full upgrade including disk

```hcl
resource "utho_cloud_resize" "full_upgrade" {
  cloud_id    = utho_cloud.app.id
  plan_id     = "10355"
  resize_type = "full"
}
```

### Re-resize to a different plan

```hcl
resource "utho_cloud_resize" "upgrade" {
  cloud_id    = utho_cloud.app.id
  plan_id     = "10400"    # update and run terraform apply
  resize_type = "ramcpu"
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `cloud_id`    | String | Yes      | The ID of the cloud instance to resize. Changing this forces a new resource. |
| `plan_id`     | String | Yes      | The target plan ID. |
| `resize_type` | String | Yes      | Resize mode. Accepted values: `ramcpu`, `full`. See below. |

### Resize Types

| Value    | Description |
|----------|-------------|
| `ramcpu` | Resize CPU and RAM only. Disk size is unchanged. Use for quick vertical scaling. |
| `full`   | Resize CPU, RAM, and disk. Use when upgrading to a plan with more storage. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cloud_id}:{plan_id}`. |
