---
page_title: "Object Storage Permission - Utho"
subcategory: "Storage / Object Storage"
description: |-
  Grant permissions on a Utho Object Storage bucket to an access key.
---

# utho_object_storage_permission

Grants a specific permission level on an Object Storage bucket to an access key. Changing the permission updates the access key's rights on the bucket in place.

Destroying this resource revokes the permission by setting it to `none`.

## Example Usage

### Grant full access

```hcl
resource "utho_object_storage_key" "deploy" {
  dcslug = "innoida"
  name   = "deploy-key"
}

resource "utho_object_storage_permission" "deploy_full" {
  dcslug      = "innoida"
  bucket_name = utho_object_storage.assets.name
  access_key  = utho_object_storage_key.deploy.access_key
  permission  = "full"
}
```

### Grant read-only access

```hcl
resource "utho_object_storage_permission" "readonly" {
  dcslug      = "innoida"
  bucket_name = utho_object_storage.assets.name
  access_key  = utho_object_storage_key.reader.access_key
  permission  = "read"
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `dcslug`      | String | Yes      | Data center slug. Changing this forces a new resource. |
| `bucket_name` | String | Yes      | Bucket name to grant permission on. Changing this forces a new resource. |
| `access_key`  | String | Yes      | Access key to grant permission to. Changing this forces a new resource. |
| `permission`  | String | Yes      | Permission level: `read`, `write`, `full`, `none`. Changing this updates in place. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{bucket_name}:{access_key}:{permission}`. |

## Permission Levels

| Value   | Description |
|---------|-------------|
| `read`  | Read objects from the bucket only. |
| `write` | Write objects to the bucket only. |
| `full`  | Read and write — full access. |
| `none`  | No access to the bucket. |
