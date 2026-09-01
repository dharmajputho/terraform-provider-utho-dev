---
page_title: "Provider: Utho"
description: |-
  The Utho provider lets you manage Utho Cloud infrastructure using Terraform.
---

# Utho Provider

The **Utho** provider allows you to create, manage, and destroy infrastructure
on [Utho Cloud](https://utho.com) using Terraform. It supports cloud instances,
networking, storage, snapshots, power management, VPC, and more.

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

The Utho provider requires an API key. Generate one from the
[Utho Dashboard](https://console.utho.com) under **Settings → API Keys**.

### Option 1 — Provider block

```hcl
provider "utho" {
  api_key = "your-api-key"
}
```

### Option 2 — Environment variable

```bash
export UTHO_API_KEY="your-api-key"
```

~> **Security note:** Never commit API keys to version control.

## Provider Arguments

| Argument  | Type   | Required | Description |
|-----------|--------|----------|-------------|
| `api_key` | String | Yes      | Utho API key. Can also be set via the `UTHO_API_KEY` environment variable. |

## Resources

### Compute / Cloud Instances

| Resource | Description |
|----------|-------------|
| [utho_cloud](resources/cloud) | Create and manage cloud instances |
| [utho_cloud_power](resources/cloud_power) | Manage power state |
| [utho_cloud_resize](resources/cloud_resize) | Resize CPU, RAM, and disk |
| [utho_cloud_snapshot](resources/cloud_snapshot) | Create, restore, and delete snapshots |
| [utho_cloud_iso](resources/cloud_iso) | Mount an ISO and boot from it |
| [utho_cloud_firewall](resources/cloud_firewall) | Attach or detach security groups |
| [utho_cloud_storage](resources/cloud_storage) | Add and manage general storage disks |
| [utho_cloud_ebs](resources/cloud_ebs) | Attach and manage EBS volumes |
| [utho_cloud_public_ip](resources/cloud_public_ip) | Assign or release additional public IPs |
| [utho_cloud_vpc](resources/cloud_vpc) | Attach or detach VPC subnets from instances |

### Networking / VPC

| Resource | Description |
|----------|-------------|
| [utho_vpc](resources/vpc) | Create and manage VPC networks |
| [utho_subnet](resources/subnet) | Create and manage subnets inside a VPC |
| [utho_nat_gateway](resources/nat_gateway) | Create and manage NAT Gateways |
| [utho_route_table](resources/route_table) | Create and manage route tables |
| [utho_route](resources/route) | Create and manage individual routes |
| [utho_elastic_ip](resources/elastic_ip) | Allocate and manage Elastic IPs |
| [utho_vpc_peering](resources/vpc_peering) | Create and manage VPC peering connections |

## Data Sources

### Compute / Cloud Instances

| Data Source | Description |
|-------------|-------------|
| [utho_clouds](data-sources/clouds) | List all cloud instances in your account |