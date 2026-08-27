---
page_title: "Provider: Utho"
description: |-
  The Utho provider lets you manage Utho Cloud infrastructure using Terraform.
---

# Utho Provider

The **Utho** provider allows you to create, manage, and destroy infrastructure
on [Utho Cloud](https://utho.com) using Terraform. It supports cloud instances,
networking, storage, snapshots, power management, and more.

## Example Usage

```hcl
terraform {
  required_providers {
    utho = {
      source  = "dharmajputho/utho-dev"
      version = "~> 0.1"
    }
  }
}

provider "utho" {
  api_key = var.utho_api_key
}

variable "utho_api_key" {
  description = "Utho API key"
  type        = string
  sensitive   = true
}
```

## Authentication

The Utho provider requires an API key. You can generate one from the
[Utho Dashboard](https://console.utho.com) under **Settings → API Keys**.

### Option 1 — Provider block (recommended for CI/CD)

```hcl
provider "utho" {
  api_key = "your-api-key"
}
```

### Option 2 — Environment variable (recommended for local dev)

```bash
export UTHO_API_KEY="your-api-key"
```

Then omit `api_key` from the provider block:

```hcl
provider "utho" {}
```

~> **Security note:** Never commit API keys to version control.
Use environment variables or a secrets manager such as HashiCorp Vault.

## Provider Arguments

| Argument  | Type   | Required | Description |
|-----------|--------|----------|-------------|
| `api_key` | String | Yes      | Utho API key. Can also be set via the `UTHO_API_KEY` environment variable. |

## Resources

| Resource | Description |
|----------|-------------|
| [utho_cloud](resources/cloud.md) | Create and manage cloud instances |
| [utho_cloud_power](resources/cloud_power.md) | Manage power state of a cloud instance |
| [utho_cloud_firewall](resources/cloud_firewall.md) | Attach or detach security groups |
| [utho_cloud_storage](resources/cloud_storage.md) | Add and manage general storage disks |
| [utho_cloud_ebs](resources/cloud_ebs.md) | Attach and manage EBS volumes |
| [utho_cloud_snapshot](resources/cloud_snapshot.md) | Create, restore, and delete snapshots |
| [utho_cloud_iso](resources/cloud_iso.md) | Mount an ISO and boot from it |
| [utho_cloud_resize](resources/cloud_resize.md) | Resize CPU, RAM, and disk |
| [utho_cloud_public_ip](resources/cloud_public_ip.md) | Assign or release additional public IPs |
| [utho_cloud_vpc](resources/cloud_vpc.md) | Attach or detach VPC subnets |

## Data Sources

| Data Source | Description |
|-------------|-------------|
| [utho_clouds](data-sources/clouds.md) | List all cloud instances in your account |
