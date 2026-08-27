---
page_title: "utho_clouds Data Source - utho"
subcategory: "Compute"
description: |-
  List all Utho Cloud instances in your account.
---

# utho_clouds

Lists all Utho Cloud instances in your account. Useful for referencing
existing instances, auditing your infrastructure, or using instance data
in other resources.

## Example Usage

### List all instances

```hcl
data "utho_clouds" "all" {}

output "all_instances" {
  value = data.utho_clouds.all.clouds
}
```

### Get IPs of all running instances

```hcl
data "utho_clouds" "all" {}

output "running_ips" {
  value = [
    for c in data.utho_clouds.all.clouds :
    c.ip if c.power_status == "Running"
  ]
}
```

### Find instances in a specific datacenter

```hcl
data "utho_clouds" "all" {}

locals {
  mumbai_instances = [
    for c in data.utho_clouds.all.clouds :
    c if c.dcslug == "inmumbaizone2"
  ]
}

output "mumbai_instances" {
  value = local.mumbai_instances
}
```

### Get instance IDs by hostname prefix

```hcl
data "utho_clouds" "all" {}

locals {
  worker_ids = [
    for c in data.utho_clouds.all.clouds :
    c.cloud_id if startswith(c.hostname, "worker-")
  ]
}
```

## Attribute Reference

This data source exports the following attributes:

### `clouds`

A list of cloud instances. Each element contains:

| Attribute      | Type   | Description |
|----------------|--------|-------------|
| `cloud_id`     | String | Unique cloud instance ID. |
| `hostname`     | String | Hostname of the instance. |
| `status`       | String | Instance status (e.g. `Active`, `Deleting`). |
| `power_status` | String | Power state (e.g. `Running`, `Not_Running`). |
| `ip`           | String | Primary public IP address. |
| `dcslug`       | String | Data center slug (e.g. `inmumbaizone2`). |
| `billingcycle` | String | Billing cycle (e.g. `hourly`, `monthly`). |
| `created_at`   | String | Creation timestamp (UTC). |
