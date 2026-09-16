---
page_title: "Alert Contact - Utho"
subcategory: "Monitoring"
description: |-
  Create and manage alert contacts for Utho monitoring alerts.
---

# utho_alert_contact

Creates and manages an alert contact. Contacts receive notifications via email and SMS when a monitoring alert is triggered. Multiple contacts can be assigned to a single alert.

## Example Usage

### Create a contact

```hcl
resource "utho_alert_contact" "ops" {
  name         = "ops-team"
  email        = "ops@mycompany.com"
  mobilenumber = "9571054173"
  status       = "1"
}

output "contact_id" {
  value = utho_alert_contact.ops.id
}
```

### Multiple contacts for different teams

```hcl
resource "utho_alert_contact" "infra" {
  name         = "infra-team"
  email        = "infra@mycompany.com"
  mobilenumber = "9571054173"
  status       = "1"
}

resource "utho_alert_contact" "oncall" {
  name         = "on-call-engineer"
  email        = "oncall@mycompany.com"
  mobilenumber = "9876543210"
  status       = "1"
}
```

### Use contacts in an alert

```hcl
resource "utho_alert_contact" "ops" {
  name         = "ops-team"
  email        = "ops@mycompany.com"
  mobilenumber = "9571054173"
  status       = "1"
}

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

## Argument Reference

| Argument       | Type   | Required | Description |
|----------------|--------|----------|-------------|
| `name`         | String | Yes      | Contact name. Updatable. |
| `email`        | String | Yes      | Email address for alert notifications. Updatable. |
| `mobilenumber` | String | Yes      | Mobile number for SMS notifications. Updatable. |
| `status`       | String | No       | Contact status: `1` (active) or `0` (inactive). Default: `1`. Updatable. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique contact ID. Use this in `utho_alert.contacts`. |

## Notes

- All fields are updatable in place.
- A contact must be active (`status = 1`) to receive alert notifications.
- Use the contact `id` in `utho_alert.contacts` as a comma-separated list.