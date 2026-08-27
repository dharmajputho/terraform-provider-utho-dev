---
page_title: "utho_cloud_power Resource - utho"
subcategory: "Compute / Cloud Instances"
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
resource "utho_cloud_power" "stop" {
  cloud_id = utho_cloud.web.id
  action   = "poweroff"
}
```

### Hard reboot an instance

```hcl
resource "utho_cloud_power" "reboot" {
  cloud_id = utho_cloud.app.id
  action   = "hardreboot"
}
```

### Power cycle after attaching a VPC

```hcl
resource "utho_cloud_vpc" "private" {
  cloud_id  = utho_cloud.app.id
  subnet_id = "b0825dae-fd86-4d67-8238-c438774da894"
}

resource "utho_cloud_power" "apply_vpc" {
  cloud_id   = utho_cloud.app.id
  action     = "powercycle"
  depends_on = [utho_cloud_vpc.private]
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
| `powercycle` | Power off then power on the instance. Required after VPC or NIC changes. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cloud_id}-{action}`. |
