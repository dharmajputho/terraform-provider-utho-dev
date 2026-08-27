---
page_title: "utho_cloud_iso Resource - utho"
subcategory: "Compute"
description: |-
  Mount an ISO image on a Utho Cloud instance and boot from it.
---

# utho_cloud_iso

Mounts a custom ISO image on a Utho Cloud instance and boots from it.
This is useful for installing custom operating systems or running
live environments.

~> **Note:** This is a one-time boot action. Once mounted, the instance
boots from the ISO. Destroying this resource removes it from Terraform
state only — there is no API call to unmount the ISO.

## Example Usage

### Mount a custom ISO

```hcl
resource "utho_cloud_iso" "custom_os" {
  cloud_id = utho_cloud.bare.id
  iso      = "ATUjo-194454.iso"
}
```

### Mount ISO on an existing instance

```hcl
resource "utho_cloud_iso" "recovery" {
  cloud_id = "1671990"
  iso      = "rescue-disk-2026.iso"
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `cloud_id` | String | Yes      | The ID of the cloud instance. Changing this forces a new resource. |
| `iso`      | String | Yes      | ISO filename as listed in your Utho ISO library (e.g. `ATUjo-194454.iso`). Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cloud_id}:{iso}`. |

## Notes

- The ISO must be uploaded to your Utho account before using this resource.
- Changing `cloud_id` or `iso` destroys the current record and creates a new one,
  which re-mounts the new ISO.
- After mounting, power-cycle the instance to boot from the ISO.
