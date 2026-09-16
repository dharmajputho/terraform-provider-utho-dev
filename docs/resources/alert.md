---
page_title: "Monitoring Alert - Utho"
subcategory: "Monitoring"
description: |-
  Create and manage Utho monitoring alert rules.
---

# utho_alert

Creates and manages a monitoring alert rule. Alerts watch a metric (CPU, RAM, disk, bandwidth) on one or more cloud instances and notify contacts when the threshold is breached for a specified duration.

## Example Usage

### CPU alert — scale up trigger

```hcl
resource "utho_alert" "cpu_high" {
  name     = "cpu-above-80"
  ref_type = "cloud"
  type     = "cpu"
  compare  = "above"
  value    = "80"
  for      = "5m"
  contacts = utho_alert_contact.ops.id
  status   = "active"
  ref_ids  = utho_cloud.app[0].id
}
```

### RAM alert on multiple servers

```hcl
resource "utho_alert" "ram_high" {
  name     = "ram-above-90"
  ref_type = "cloud"
  type     = "ram"
  compare  = "above"
  value    = "90"
  for      = "10m"
  contacts = "${utho_alert_contact.ops.id},${utho_alert_contact.oncall.id}"
  status   = "active"
  ref_ids  = join(",", utho_cloud.app[*].id)
}
```

### Disk usage alert

```hcl
resource "utho_alert" "disk_full" {
  name     = "disk-above-90"
  ref_type = "cloud"
  type     = "disk"
  compare  = "above"
  value    = "90"
  for      = "30m"
  contacts = utho_alert_contact.ops.id
  status   = "active"
  ref_ids  = utho_cloud.db.id
}
```

### Bandwidth alert

```hcl
resource "utho_alert" "low_bandwidth" {
  name     = "bandwidth-below-40"
  ref_type = "cloud"
  type     = "bandwidth"
  compare  = "below"
  value    = "40"
  for      = "30m"
  contacts = utho_alert_contact.ops.id
  status   = "active"
  ref_ids  = utho_cloud.app[0].id
}
```

### Full monitoring setup

```hcl
resource "utho_alert_contact" "ops" {
  name         = "ops-team"
  email        = "ops@mycompany.com"
  mobilenumber = "9571054173"
  status       = "1"
}

resource "utho_alert" "cpu" {
  name     = "cpu-above-80"
  ref_type = "cloud"
  type     = "cpu"
  compare  = "above"
  value    = "80"
  for      = "5m"
  contacts = utho_alert_contact.ops.id
  status   = "active"
  ref_ids  = join(",", utho_cloud.app[*].id)
}

resource "utho_alert" "ram" {
  name     = "ram-above-85"
  ref_type = "cloud"
  type     = "ram"
  compare  = "above"
  value    = "85"
  for      = "5m"
  contacts = utho_alert_contact.ops.id
  status   = "active"
  ref_ids  = join(",", utho_cloud.app[*].id)
}

resource "utho_alert" "disk" {
  name     = "disk-above-90"
  ref_type = "cloud"
  type     = "disk"
  compare  = "above"
  value    = "90"
  for      = "15m"
  contacts = utho_alert_contact.ops.id
  status   = "active"
  ref_ids  = join(",", utho_cloud.app[*].id)
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `name`     | String | Yes      | Alert rule name. Updatable. |
| `ref_type` | String | Yes      | Resource type to monitor. Currently: `cloud`. Changing this forces a new resource. |
| `type`     | String | Yes      | Metric to monitor: `cpu`, `ram`, `disk`, `bandwidth`. Changing this forces a new resource. |
| `compare`  | String | Yes      | Comparison operator: `above` or `below`. Updatable. |
| `value`    | String | Yes      | Threshold percentage (e.g. `80` for 80%). Updatable. |
| `for`      | String | Yes      | Duration the threshold must be breached before alerting: `5m`, `10m`, `15m`, `30m`, `1h`. Updatable. |
| `contacts` | String | Yes      | Comma-separated contact IDs to notify (e.g. `426,508`). Updatable. |
| `status`   | String | Yes      | Alert status: `active` or `inactive`. Updatable. |
| `ref_ids`  | String | Yes      | Comma-separated cloud instance IDs to monitor (e.g. `1671990,990001511`). Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique alert rule ID. |

## Metric Types

| Value       | Description |
|-------------|-------------|
| `cpu`       | CPU usage percentage. |
| `ram`       | RAM usage percentage. |
| `disk`      | Disk usage percentage. |
| `bandwidth` | Network bandwidth usage percentage. |

## Duration Values

| Value  | Description |
|--------|-------------|
| `5m`   | 5 minutes   |
| `10m`  | 10 minutes  |
| `15m`  | 15 minutes  |
| `30m`  | 30 minutes  |
| `1h`   | 1 hour      |

## Notes

- `name`, `compare`, `value`, `for`, `contacts`, and `status` are all updatable in place.
- Changing `ref_type`, `type`, or `ref_ids` destroys and recreates the alert.
- Multiple contacts can be notified by passing comma-separated IDs: `contacts = "426,508"`.
- Multiple instances can be monitored by passing comma-separated IDs: `ref_ids = "1671990,990001511"`.