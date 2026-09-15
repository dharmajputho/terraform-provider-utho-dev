---
page_title: "Object Storage - Utho"
subcategory: "Storage / Object Storage"
description: |-
  Create and manage Utho Object Storage buckets (S3-compatible).
---

# utho_object_storage

Creates and manages a Utho Object Storage bucket. Buckets are S3-compatible and can be used as a remote backend for Terraform state, application file storage, backups, and more.

## Example Usage

### Basic private bucket

```hcl
resource "utho_object_storage" "assets" {
  name   = "my-app-assets"
  dcslug = "innoida"
}

output "endpoint" {
  value = "https://innoida.utho.io"
}

output "access_key" {
  value     = utho_object_storage.assets.access_key
  sensitive = true
}

output "secret_key" {
  value     = utho_object_storage.assets.secret_key
  sensitive = true
}
```

### Public read bucket

```hcl
resource "utho_object_storage" "public_assets" {
  name   = "my-public-assets"
  dcslug = "innoida"
  access = "public"
}
```

### Bucket with versioning enabled

```hcl
resource "utho_object_storage" "versioned" {
  name            = "my-versioned-bucket"
  dcslug          = "innoida"
  version_enabled = true
}
```

### Use as Terraform remote state backend

```hcl
resource "utho_object_storage" "tfstate" {
  name   = "terraform-state"
  dcslug = "innoida"
  access = "private"
}

# After creating, configure backend:
# terraform {
#   backend "s3" {
#     bucket     = "terraform-state"
#     key        = "prod/terraform.tfstate"
#     endpoint   = "https://innoida.utho.io"
#     access_key = "<access_key>"
#     secret_key = "<secret_key>"
#     region     = "us-east-1"
#   }
# }
```

## Argument Reference

| Argument          | Type   | Required | Description |
|-------------------|--------|----------|-------------|
| `name`            | String | Yes      | Bucket name. Must be unique. Changing this forces a new resource. |
| `dcslug`          | String | Yes      | Data center slug. Currently only `innoida` is supported. Changing this forces a new resource. |
| `access`          | String | No       | Bucket access policy: `private`, `public`, `upload`. Default: `private`. |
| `version_enabled` | Bool   | No       | Enable object versioning. Default: `false`. |
| `billingcycle`    | String | No       | Billing cycle. Default: `monthly`. Changing this forces a new resource. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Bucket name (used as unique identifier). |
| `access_key` | String | S3-compatible access key. **Sensitive.** |
| `secret_key` | String | S3-compatible secret key. **Sensitive.** |
| `status`     | String | Bucket status. |
| `plan_gb`    | Number | Plan storage allocation in GB. |
| `used_gb`    | String | Current storage usage in GB. |
| `created_at` | String | Creation timestamp. |

## Billing

Utho Object Storage is billed at ₹500/month per bucket, covering up to 100 GB. Usage above 100 GB is billed per GB on the overage rate. Writes are never blocked.

## Notes

- Object manipulation (upload, download, delete) is done via the S3 SDK using `access_key` and `secret_key` — not through Terraform.
- Compatible with any AWS S3 SDK, s3cmd, MinIO client, or rclone.
- S3 endpoint: `https://innoida.utho.io`
