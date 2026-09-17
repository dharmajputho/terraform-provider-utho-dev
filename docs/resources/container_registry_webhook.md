---
page_title: "Container Registry Webhook - Utho"
subcategory: "Container Registry"
description: |-
  Create and manage webhooks for a Utho Container Registry.
---

# utho_container_registry_webhook

Creates and manages a webhook for a Utho Container Registry. Webhooks send HTTP notifications to an external endpoint when registry events occur (image push, scan completion, deletion, etc.).

## Example Usage

### Webhook for push and scan events

```hcl
resource "utho_container_registry_webhook" "notify" {
  project_name   = utho_container_registry.main.project_name
  name           = "ci-notify"
  address        = "https://hooks.mycompany.com/registry"
  payload_format = "CloudEvents"
  skip_cert_verify = false
  auth_header    = "Bearer ${var.webhook_secret}"

  event_types = [
    "PUSH_ARTIFACT",
    "SCANNING_COMPLETED",
    "DELETE_ARTIFACT",
  ]
}
```

### All events webhook

```hcl
resource "utho_container_registry_webhook" "all_events" {
  project_name   = utho_container_registry.main.project_name
  name           = "all-events"
  address        = "https://webhook.site/my-endpoint"
  payload_format = "Default"
  skip_cert_verify = true

  event_types = [
    "PUSH_ARTIFACT",
    "SCANNING_COMPLETED",
    "DELETE_ARTIFACT",
    "PULL_ARTIFACT",
    "TAG_RETENTION",
    "REPLICATION",
    "SCANNING_FAILED",
    "QUOTA_EXCEED",
  ]
}
```

## Argument Reference

| Argument          | Type   | Required | Description |
|-------------------|--------|----------|-------------|
| `project_name`    | String | Yes      | Registry project name. Changing this forces a new resource. |
| `name`            | String | Yes      | Webhook name. Changing this forces a new resource. |
| `address`         | String | Yes      | Webhook endpoint URL. Changing this forces a new resource. |
| `payload_format`  | String | Yes      | Payload format: `Default` or `CloudEvents`. Changing this forces a new resource. |
| `event_types`     | List   | Yes      | Event types to trigger. See [Event Types](#event-types). Changing this forces a new resource. |
| `skip_cert_verify`| Bool   | No       | Skip TLS certificate verification. Changing this forces a new resource. |
| `auth_header`     | String | No       | Authorization header (e.g. `Bearer token`). **Sensitive.** Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique webhook ID. |

## Event Types

| Value                | Description |
|----------------------|-------------|
| `PUSH_ARTIFACT`      | Triggered when an image is pushed. |
| `PULL_ARTIFACT`      | Triggered when an image is pulled. |
| `DELETE_ARTIFACT`    | Triggered when an artifact is deleted. |
| `SCANNING_COMPLETED` | Triggered when a vulnerability scan completes. |
| `SCANNING_FAILED`    | Triggered when a vulnerability scan fails. |
| `TAG_RETENTION`      | Triggered by tag retention policy runs. |
| `REPLICATION`        | Triggered by replication events. |
| `QUOTA_EXCEED`       | Triggered when storage quota is exceeded. |