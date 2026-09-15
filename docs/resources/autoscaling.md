---
page_title: "Auto Scaling Group - Utho"
subcategory: "Compute / Auto Scaling"
description: |-
  Create and manage Utho Auto Scaling groups.
---

# utho_autoscaling

Creates and manages a Utho Auto Scaling group. Auto Scaling automatically adjusts the number of cloud instances based on CPU/RAM metrics or a defined schedule. Supports load balancer and security group attachments at creation time.

## Example Usage

### Basic auto scaling group with a stack

```hcl
resource "utho_autoscaling" "web" {
  name             = "web-asg"
  dcslug           = "inmumbaizone2"
  planid           = "10314"
  planname         = "basic"
  os_disk_size     = 80
  minsize          = "1"
  maxsize          = "5"
  desiredsize      = "2"
  public_ip_enabled = 1
  stack            = "6669726"
  stackid          = "6669726"
  stackimage       = "ubuntu-22.04-x86_64"
  cpumodel         = "amd"
}
```

### Auto scaling group with policies

```hcl
resource "utho_autoscaling" "app" {
  name             = "app-asg"
  dcslug           = "inmumbaizone2"
  planid           = "10314"
  planname         = "basic"
  os_disk_size     = 80
  minsize          = "1"
  maxsize          = "10"
  desiredsize      = "2"
  public_ip_enabled = 1
  stack            = "6669726"
  stackid          = "6669726"
  stackimage       = "ubuntu-22.04-x86_64"
  load_balancers   = utho_loadbalancer.main.id
  security_groups  = utho_firewall.web.id

  policies = [
    {
      name     = "scale-up-cpu"
      type     = "cpu"
      compare  = "above"
      value    = "80"
      adjust   = 2
      period   = "5m"
      cooldown = "300"
    },
    {
      name     = "scale-down-cpu"
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

### Auto scaling group with schedule

```hcl
resource "utho_autoscaling" "app" {
  name             = "app-asg"
  dcslug           = "inmumbaizone2"
  planid           = "10314"
  planname         = "basic"
  os_disk_size     = 80
  minsize          = "1"
  maxsize          = "10"
  desiredsize      = "2"
  public_ip_enabled = 1
  stack            = "6669726"
  stackid          = "6669726"
  stackimage       = "ubuntu-22.04-x86_64"

  schedules = [
    {
      name                = "peak-hours"
      desiredsize         = "5"
      timezone            = "Asia/Kolkata"
      recurrence          = "Every day 09:00"
      recurrence_duration = "Every day"
      recurrence_week     = ""
      selected_time       = "09:00"
      selected_date       = "2026-09-15"
      start_date          = "2026-09-15T09:00:00.000+05:30"
    }
  ]
}
```

### Auto scaling inside a VPC

```hcl
resource "utho_autoscaling" "private" {
  name              = "private-asg"
  dcslug            = "inmumbaizone2"
  planid            = "10314"
  planname          = "basic"
  os_disk_size      = 80
  minsize           = "2"
  maxsize           = "8"
  desiredsize       = "2"
  public_ip_enabled = 0
  vpc               = utho_subnet.private.id
  stack             = "6669726"
  stackid           = "6669726"
  stackimage        = "ubuntu-22.04-x86_64"
}
```

## Argument Reference

### Required

| Argument           | Type   | Description |
|--------------------|--------|-------------|
| `name`             | String | Auto scaling group name. Changing this forces a new resource. |
| `dcslug`           | String | Data center slug. Changing this forces a new resource. |
| `planid`           | String | Plan ID for instance size. Changing this forces a new resource. |
| `planname`         | String | Plan name (e.g. `basic`). Changing this forces a new resource. |
| `os_disk_size`     | Number | OS disk size in GB. Changing this forces a new resource. |
| `minsize`          | String | Minimum number of instances. |
| `maxsize`          | String | Maximum number of instances. |
| `desiredsize`      | String | Desired number of instances at launch. |
| `public_ip_enabled`| Number | Assign public IPs to instances: `1` or `0`. |

### Required — one of

| Argument     | Type   | Description |
|--------------|--------|-------------|
| `stack`      | String | Stack ID to deploy instances from. |
| `stackid`    | String | Stack ID (same as `stack`). |
| `stackimage` | String | Stack image slug (e.g. `ubuntu-22.04-x86_64`). |
| `snapshotid` | String | Snapshot ID to deploy instances from (alternative to stack). |

### Optional

| Argument         | Type   | Description |
|------------------|--------|-------------|
| `vpc`            | String | VPC subnet ID. Changing this forces a new resource. |
| `load_balancers` | String | Load balancer ID to attach. |
| `security_groups`| String | Security group ID to attach. |
| `target_groups`  | String | Target group ID to attach (alternative to `load_balancers`). |
| `backupid`       | String | Backup ID to deploy from. Changing this forces a new resource. |
| `cpumodel`       | String | CPU model: `amd` or `intel`. Changing this forces a new resource. |
| `policies`       | List   | Scaling policies. See [Policy Block](#policy-block). |
| `schedules`      | List   | Scheduled scaling. See [Schedule Block](#schedule-block). |

### Policy Block

| Argument   | Type   | Description |
|------------|--------|-------------|
| `name`     | String | Policy name. |
| `type`     | String | Metric: `cpu` or `ram`. |
| `compare`  | String | Trigger when metric is `above` or `below` the threshold. |
| `value`    | String | Threshold value (e.g. `80` for 80%). |
| `adjust`   | Number | Instances to add (positive) or remove (negative). |
| `period`   | String | Evaluation period (e.g. `5m`). |
| `cooldown` | String | Cooldown period in seconds after a scaling action. |

### Schedule Block

| Argument              | Type   | Description |
|-----------------------|--------|-------------|
| `name`                | String | Schedule name. |
| `desiredsize`         | String | Desired instance count at the scheduled time. |
| `timezone`            | String | Timezone (e.g. `Asia/Kolkata`). |
| `recurrence`          | String | Recurrence expression (e.g. `Every day 09:00`). |
| `recurrence_duration` | String | Duration type (e.g. `Every day`). |
| `recurrence_week`     | String | Week day for weekly recurrence (optional). |
| `selected_time`       | String | Time in `HH:MM` format. |
| `selected_date`       | String | Date in `YYYY-MM-DD` format. |
| `start_date`          | String | Start datetime in ISO 8601 format. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique auto scaling group ID. |
| `status`     | String | Current status. |
| `created_at` | String | Creation timestamp. |