---
page_title: "Auto Scaling Group - Utho"
subcategory: "Compute / Auto Scaling"
description: |-
  Create and manage Utho Auto Scaling groups.
---

# utho_autoscaling

Creates and manages a Utho Auto Scaling group. An auto scaling group automatically adjusts the number of cloud instances based on CPU or RAM thresholds, or on a time-based schedule. When load spikes, it adds instances; when load drops, it removes them.

Add scaling policies later with `utho_autoscaling_policy`, or define them inline at creation.

## Example Usage

### Basic auto scaling group

A simple group that scales between 1 and 5 instances based on CPU.

```hcl
resource "utho_autoscaling" "web" {
  name              = "web-asg"
  dcslug            = "inmumbaizone2"
  planid            = "10314"
  planname          = "basic"
  os_disk_size      = 80
  minsize           = "1"
  maxsize           = "5"
  desiredsize       = "2"
  public_ip_enabled = 1
  stack             = "6669726"
  stackid           = "6669726"
  stackimage        = "ubuntu-22.04-x86_64"
  cpumodel          = "amd"

  policies = [
    {
      name     = "scale-up"
      type     = "cpu"
      compare  = "above"
      value    = "80"
      adjust   = 2
      period   = "5m"
      cooldown = "300"
    },
    {
      name     = "scale-down"
      type     = "cpu"
      compare  = "below"
      value    = "20"
      adjust   = -1
      period   = "5m"
      cooldown = "300"
    }
  ]
}
```

### Auto scaling group behind a load balancer

New instances are automatically added to the load balancer as they come up.

```hcl
resource "utho_autoscaling" "app" {
  name              = "app-asg"
  dcslug            = "inmumbaizone2"
  planid            = "10314"
  planname          = "basic"
  os_disk_size      = 80
  minsize           = "2"
  maxsize           = "10"
  desiredsize       = "3"
  public_ip_enabled = 1
  load_balancers    = utho_loadbalancer.main.id
  security_groups   = utho_firewall.web.id
  stack             = "6669726"
  stackid           = "6669726"
  stackimage        = "ubuntu-22.04-x86_64"

  policies = [
    {
      name     = "scale-up-cpu"
      type     = "cpu"
      compare  = "above"
      value    = "75"
      adjust   = 2
      period   = "5m"
      cooldown = "300"
    }
  ]
}
```

### Auto scaling with scheduled scaling

Scale up before peak hours, scale down at night.

```hcl
resource "utho_autoscaling" "app" {
  name              = "app-asg"
  dcslug            = "inmumbaizone2"
  planid            = "10314"
  planname          = "basic"
  os_disk_size      = 80
  minsize           = "1"
  maxsize           = "10"
  desiredsize       = "2"
  public_ip_enabled = 1
  stack             = "6669726"
  stackid           = "6669726"
  stackimage        = "ubuntu-22.04-x86_64"

  schedules = [
    {
      name                = "morning-scale-up"
      desiredsize         = "6"
      timezone            = "Asia/Kolkata"
      recurrence          = "Every day 09:00"
      recurrence_duration = "Every day"
      recurrence_week     = ""
      selected_time       = "09:00"
      selected_date       = "2026-09-17"
      start_date          = "2026-09-17T09:00:00.000+05:30"
    },
    {
      name                = "night-scale-down"
      desiredsize         = "1"
      timezone            = "Asia/Kolkata"
      recurrence          = "Every day 23:00"
      recurrence_duration = "Every day"
      recurrence_week     = ""
      selected_time       = "23:00"
      selected_date       = "2026-09-17"
      start_date          = "2026-09-17T23:00:00.000+05:30"
    }
  ]
}
```

## Argument Reference

### Required

| Argument            | Type   | Description |
|---------------------|--------|-------------|
| `name`              | String | Group name. Changing this forces a new resource. |
| `dcslug`            | String | Data center. Changing this forces a new resource. |
| `planid`            | String | Plan ID for each instance. Changing this forces a new resource. |
| `planname`          | String | Plan name (e.g. `basic`). Changing this forces a new resource. |
| `os_disk_size`      | Number | OS disk size in GB. Changing this forces a new resource. |
| `minsize`           | String | Minimum number of instances. |
| `maxsize`           | String | Maximum number of instances. |
| `desiredsize`       | String | Starting instance count. |
| `public_ip_enabled` | Number | `1` to assign public IPs, `0` for private only. |

### Image source — one required

| Argument     | Description |
|--------------|-------------|
| `stack` + `stackid` + `stackimage` | Deploy from a marketplace or custom stack. |
| `snapshotid` | Deploy from a snapshot (alternative to stack). |

### Optional

| Argument         | Description |
|------------------|-------------|
| `vpc`            | VPC subnet ID. Changing this forces a new resource. |
| `load_balancers` | Load balancer ID to attach instances to. |
| `security_groups`| Security group ID. |
| `target_groups`  | Target group ID (alternative to `load_balancers`). |
| `cpumodel`       | `amd` or `intel`. Changing this forces a new resource. |
| `policies`       | Inline scaling policies. See [Policy Block](#policy-block). |
| `schedules`      | Scheduled scaling. See [Schedule Block](#schedule-block). |

### Policy Block

```hcl
policies = [
  {
    name     = "scale-up"
    type     = "cpu"       # or "ram"
    compare  = "above"     # or "below"
    value    = "80"        # percentage threshold
    adjust   = 2           # positive = add, negative = remove
    period   = "5m"        # evaluation window
    cooldown = "300"       # seconds before next scale action
  }
]
```

### Schedule Block

```hcl
schedules = [
  {
    name                = "peak-hours"
    desiredsize         = "5"
    timezone            = "Asia/Kolkata"
    recurrence          = "Every day 09:00"
    recurrence_duration = "Every day"
    recurrence_week     = ""
    selected_time       = "09:00"
    selected_date       = "2026-09-17"
    start_date          = "2026-09-17T09:00:00.000+05:30"
  }
]
```

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique auto scaling group ID. |
| `status`     | String | Group status. |
| `created_at` | String | Creation timestamp. |
