---
page_title: "utho_cloud_iso Resource - utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Mount an ISO image on a Utho Cloud instance and boot from it.
---

# utho_cloud_iso

Mounts a custom ISO image on a Utho Cloud instance and boots from it.
Useful for installing custom operating systems or running live environments.

~> **Note:** ISO mount is a one-time boot action. Destroying this resource
removes the record from Terraform state only — there is no API call to unmount.

## Example Usage

### Mount a custom ISO

```hcl
resource "utho_cloud_iso" "custom_os" {
  cloud_id = utho_cloud.bare.id
  iso      = "ATUjo-194454.iso"
}
```

### Mount and power-cycle to boot

```hcl
resource "utho_cloud_iso" "recovery" {
  cloud_id = "1671990"
  iso      = "rescue-disk-2026.iso"
}

resource "utho_cloud_power" "reboot_into_iso" {
  cloud_id   = "1671990"
  action     = "powercycle"
  depends_on = [utho_cloud_iso.recovery]
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
- After mounting, power-cycle the instance using `utho_cloud_power` to boot from the ISO.
- Changing `cloud_id` or `iso` destroys the record and re-executes the mount with the new values.
