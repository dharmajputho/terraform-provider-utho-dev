---
page_title: "Object Storage Access Key - Utho"
subcategory: "Storage / Object Storage"
description: |-
  Create and manage access keys for Utho Object Storage.
---

# utho_object_storage_key

Creates and manages an S3-compatible access key for Utho Object Storage. Access keys are used to authenticate with buckets via the S3 API.

## Example Usage

### Create an access key

```hcl
resource "utho_object_storage_key" "app" {
  dcslug = "innoida"
  name   = "app-key"
}

output "access_key" {
  value     = utho_object_storage_key.app.access_key
  sensitive = true
}

output "secret_key" {
  value     = utho_object_storage_key.app.secret_key
  sensitive = true
}
```

### Create key and grant bucket access

```hcl
resource "utho_object_storage_key" "deploy" {
  dcslug = "innoida"
  name   = "deploy-key"
}

resource "utho_object_storage_permission" "access" {
  dcslug      = "innoida"
  bucket_name = utho_object_storage.assets.name
  access_key  = utho_object_storage_key.deploy.access_key
  permission  = "full"
}
```

### Disable an access key

```hcl
resource "utho_object_storage_key" "app" {
  dcslug = "innoida"
  name   = "app-key"
  status = "disable"
}
```

## Argument Reference

| Argument | Type   | Required | Description |
|----------|--------|----------|-------------|
| `dcslug` | String | Yes      | Data center slug. Changing this forces a new resource. |
| `name`   | String | Yes      | Access key name. Changing this forces a new resource. |
| `status` | String | No       | Key status: `enable` or `disable`. Default: `enable`. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | The access key string (used as identifier). |
| `access_key` | String | Generated S3 access key. **Sensitive.** |
| `secret_key` | String | Generated S3 secret key. **Sensitive — shown only at creation time.** |
| `status`     | String | Current key status. |

~> **Important:** The `secret_key` is only available at creation time. Save it immediately — it cannot be retrieved later.
