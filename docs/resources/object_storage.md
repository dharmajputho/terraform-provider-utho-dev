---
page_title: "Object Storage - Utho"
subcategory: "Storage / Object Storage"
description: |-
  Create and manage Utho Object Storage buckets (S3-compatible).
---

# utho_object_storage

Creates and manages a Utho Object Storage bucket. Buckets are S3-compatible — you can use any AWS S3 SDK, CLI tool, or library to read and write objects. Terraform manages the bucket itself; your application code manages the objects inside it.

Common uses: application file uploads, static assets, database backups, Terraform remote state, logs.

## Example Usage

### Private bucket for application uploads

```hcl
resource "utho_object_storage" "uploads" {
  name   = "myapp-uploads"
  dcslug = "innoida"
  access = "private"
}

# Use these in your application config
output "s3_endpoint"   { value = "https://innoida.utho.io" }
output "s3_access_key" { value = utho_object_storage.uploads.access_key; sensitive = true }
output "s3_secret_key" { value = utho_object_storage.uploads.secret_key; sensitive = true }
output "s3_bucket"     { value = utho_object_storage.uploads.name }
```

### Public static assets bucket

For hosting static files (images, CSS, JS) that your website needs to serve publicly.

```hcl
resource "utho_object_storage" "assets" {
  name   = "myapp-assets"
  dcslug = "innoida"
  access = "public"
}
```

### Terraform remote state backend

Store your Terraform state in object storage for team collaboration.

```hcl
resource "utho_object_storage" "tfstate" {
  name   = "mycompany-terraform-state"
  dcslug = "innoida"
  access = "private"
}

# Add this to a separate backend.tf file and run terraform init
# terraform {
#   backend "s3" {
#     bucket                      = "mycompany-terraform-state"
#     key                         = "prod/terraform.tfstate"
#     endpoint                    = "https://innoida.utho.io"
#     access_key                  = "<access_key>"
#     secret_key                  = "<secret_key>"
#     region                      = "us-east-1"
#     skip_credentials_validation = true
#     skip_metadata_api_check     = true
#     force_path_style            = true
#   }
# }
```

### Versioned bucket for backup storage

Enable versioning so every overwrite creates a new version, protecting against accidental deletion.

```hcl
resource "utho_object_storage" "backups" {
  name            = "myapp-backups"
  dcslug          = "innoida"
  access          = "private"
  version_enabled = true
}
```

### Bucket with a dedicated access key and permissions

Create a restricted key with only the access this service needs.

```hcl
resource "utho_object_storage" "data" {
  name   = "myapp-data"
  dcslug = "innoida"
  access = "private"
}

resource "utho_object_storage_key" "app" {
  dcslug = "innoida"
  name   = "myapp-service-key"
}

resource "utho_object_storage_permission" "app" {
  dcslug      = "innoida"
  bucket_name = utho_object_storage.data.name
  access_key  = utho_object_storage_key.app.access_key
  permission  = "full"
}
```

## Argument Reference

| Argument          | Type   | Required | Description |
|-------------------|--------|----------|-------------|
| `name`            | String | Yes      | Bucket name. Must be globally unique. Changing this forces a new resource. |
| `dcslug`          | String | Yes      | Data center. Currently only `innoida` is supported for object storage. Changing this forces a new resource. |
| `access`          | String | No       | Bucket visibility: `private` (default), `public` (public read), or `upload` (public upload). |
| `version_enabled` | Bool   | No       | Enable object versioning. Protects against overwrites and deletions. Default: `false`. |
| `billingcycle`    | String | No       | Billing cycle: `monthly`. Changing this forces a new resource. |

## Attribute Reference

| Attribute      | Type   | Description |
|----------------|--------|-------------|
| `id`           | String | Bucket name (used as unique identifier). |
| `access_key`   | String | S3 access key for this bucket. **Sensitive.** |
| `secret_key`   | String | S3 secret key for this bucket. **Sensitive.** |
| `status`       | String | Bucket status (`Active`). |
| `plan_gb`      | Number | Storage allocation in GB (100 GB minimum). |
| `used_gb`      | String | Current storage usage in GB. |
| `created_at`   | String | Creation timestamp. |

## Connecting with AWS SDK

```python
import boto3

s3 = boto3.client(
    "s3",
    endpoint_url="https://innoida.utho.io",
    aws_access_key_id=ACCESS_KEY,
    aws_secret_access_key=SECRET_KEY,
)

s3.upload_file("myfile.txt", "myapp-uploads", "myfile.txt")
```

## Billing

- ₹500/month minimum per bucket, covering up to 100 GB
- Usage above 100 GB is billed per GB (overage rate)
- Writes are never blocked regardless of usage

## Notes

- Object storage endpoint: `https://innoida.utho.io`
- Compatible with AWS SDK v2, boto3, s3cmd, rclone, MinIO client, and any S3-compatible tool
- Terraform manages the bucket — use the S3 SDK to manage objects inside it
- The `access_key` and `secret_key` in the bucket response are the bucket's own credentials — different from keys created via `utho_object_storage_key`
