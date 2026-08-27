---
page_title: "utho_cloud_power Resource - utho"
subcategory: "Compute"
description: |-
  Manage the power state of a Utho Cloud instance.
---

# utho_cloud_power

Manages the power state of a Utho Cloud instance. Supports power on,
power off, hard reboot, and power cycle operations.

~> **Note:** Power actions are imperative — they execute immediately when
`terraform apply` is run. Changing the `action` field triggers the new
action on the next apply.

## Example Usage

### Power off an instance

```hcl
resource "utho_cloud_power" "stop_web" {
  cloud_id = "1671990"
  action   = "poweroff"
}
```

### Reboot an instance

```hcl
resource "utho_cloud_power" "reboot_app" {
  cloud_id = utho_cloud.app.id
  action   = "hardreboot"
}
```

### Power cycle (force restart)

```hcl
resource "utho_cloud_power" "cycle_db" {
  cloud_id = "1671990"
  action   = "powercycle"
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `cloud_id` | String | Yes      | The ID of the cloud instance to manage. |
| `action`   | String | Yes      | Power action to perform. See [Actions](#actions) below. |

### Actions

| Value        | Description |
|--------------|-------------|
| `poweron`    | Start a stopped instance. |
| `poweroff`   | Gracefully shut down a running instance. |
| `hardreboot` | Force reboot the instance (equivalent to a hard reset). |
| `powercycle` | Power cycle the instance (power off then power on). |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cloud_id}-{action}`. |
