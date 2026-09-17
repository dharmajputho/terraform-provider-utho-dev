---
page_title: "Monitoring Alert - Utho"
subcategory: "Monitoring"
description: |-
  Create and manage metric-based monitoring alerts for cloud instances.
---

# utho_alert

Creates and manages a monitoring alert rule. Alerts watch a metric (CPU, RAM, disk, bandwidth) on one or more cloud instances. When the metric crosses a threshold for a specified duration, Utho sends an email and SMS notification to the configured contacts.

## Example Usage

### CPU spike alert

Alert your team when CPU stays above 80% for 5 minutes — a sign your servers are overloaded.

```hcl
resource "utho_alert_contact" "oncall" {
  name         = "on-call"
  email        = "oncall@mycompany.com"
  mobilenumber = "9876543210"
  status       = "1"
}

resource "utho_alert" "cpu_high" {
  name     = "cpu-above-80"
  ref_type = "cloud"
  type     = "cpu"
  compare  = "above"
  value    = "80"
  for      = "5m"
  contacts = utho_alert_contact.oncall.id
  status   = "active"
  ref_ids  = utho_cloud.app[0].id
}
```

### Monitor all servers in a fleet

Watch all your app servers at once. Use `join(",", ...)` to build a comma-separated list of IDs.

```hcl
resource "utho_alert" "cpu_fleet" {
  name     = "fleet-cpu-alert"
  ref_type = "cloud"
  type     = "cpu"
  compare  = "above"
  value    = "85"
  for      = "10m"
  contacts = utho_alert_contact.oncall.id
  status   = "active"
  ref_ids  = join(",", utho_cloud.app[*].id)
}
```

### Full monitoring setup — CPU, RAM, and disk

```hcl
resource "utho_alert_contact" "ops" {
  name         = "ops-team"
  email        = "ops@mycompany.com"
  mobilenumber = "9571054173"
  status       = "1"
}

resource "utho_alert" "cpu" {
  name     = "cpu-high"
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
  name     = "ram-high"
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
  name     = "disk-full"
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

### Notify multiple contacts

Pass a comma-separated list of contact IDs.

```hcl
resource "utho_alert" "critical" {
  name     = "critical-cpu"
  ref_type = "cloud"
  type     = "cpu"
  compare  = "above"
  value    = "95"
  for      = "5m"
  contacts = "${utho_alert_contact.ops.id},${utho_alert_contact.oncall.id}"
  status   = "active"
  ref_ids  = utho_cloud.prod.id
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `name`     | String | Yes      | Alert name. Updatable. |
| `ref_type` | String | Yes      | Resource type: `cloud`. Changing this forces a new resource. |
| `type`     | String | Yes      | Metric: `cpu`, `ram`, `disk`, `bandwidth`. Changing this forces a new resource. |
| `compare`  | String | Yes      | `above` or `below` the threshold. Updatable. |
| `value`    | String | Yes      | Threshold percentage (e.g. `"80"` for 80%). Updatable. |
| `for`      | String | Yes      | How long the threshold must be breached: `5m`, `10m`, `15m`, `30m`, `1h`. Updatable. |
| `contacts` | String | Yes      | Comma-separated contact IDs. Updatable. |
| `status`   | String | Yes      | `active` or `inactive`. Updatable. |
| `ref_ids`  | String | Yes      | Comma-separated cloud instance IDs to monitor. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique alert rule ID. |

## Recommended Thresholds

| Metric      | Warning | Critical |
|-------------|---------|----------|
| CPU         | 70%     | 85%      |
| RAM         | 75%     | 90%      |
| Disk        | 80%     | 90%      |
| Bandwidth   | — (use `below` for underutilization) | |
