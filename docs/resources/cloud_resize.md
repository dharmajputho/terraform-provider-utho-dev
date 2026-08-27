---
page_title: "utho_cloud_resize Resource - utho"
subcategory: "Compute"
description: |-
  Resize a Utho Cloud instance — CPU/RAM only or full resize including disk.
---

# utho_cloud_resize

Resizes a Utho Cloud instance to a different plan. Supports two modes:
CPU/RAM only resize (`ramcpu`) which does not affect disk size,
and full resize (`full`) which upgrades CPU, RAM, and disk together.

~> **Note:** Resize operations are applied immediately. Downgrading to a
smaller plan may not be possible if the current disk usage exceeds the
target plan's disk allocation. Use `full` resize only when upgrading.

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
resource "utho_cloud_resize" "scale_up" {
  cloud_id    = utho_cloud.app.id
  plan_id     = "10400"   # change plan and run terraform apply
  resize_type = "ramcpu"
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `cloud_id`    | String | Yes      | The ID of the cloud instance to resize. Changing this forces a new resource. |
| `plan_id`     | String | Yes      | The target plan ID. Contact Utho support or check the dashboard for available plan IDs. |
| `resize_type` | String | Yes      | Resize mode. See [Resize Types](#resize-types) below. |

### Resize Types

| Value    | Description |
|----------|-------------|
| `ramcpu` | Resize CPU and RAM only. Disk size remains unchanged. Useful for quick vertical scaling without data risk. |
| `full`   | Resize CPU, RAM, and disk. Use when upgrading to a larger plan and additional storage is needed. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cloud_id}:{plan_id}`. |

## Notes

- Changing `plan_id` or `resize_type` applies the new resize when
  `terraform apply` is run.
- Resize operations may require a brief instance restart.
- Disk downgrades are not supported — use `ramcpu` when moving to a
  plan with a smaller disk.
