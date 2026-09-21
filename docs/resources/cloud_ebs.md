---
page_title: "Cloud EBS - Utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Attach and manage EBS block volumes on a Utho Cloud instance.
---

# utho_cloud_ebs

Attaches and manages an EBS (Elastic Block Storage) volume on an existing Utho Cloud instance. Use EBS volumes for additional persistent storage — databases, large datasets, logs, or any data that must survive instance rebuilds.

~> **Note:** EBS volumes can also be attached at instance creation time using the `ebs` block inside `utho_cloud`. Use `utho_cloud_ebs` to attach volumes to already-running instances.

## Example Usage

### Attach a single EBS volume

```hcl
resource "utho_cloud_ebs" "data" {
  cloud_id = utho_cloud.db.id
  disk     = 100
  type     = "nvme"
}
```

### Attach multiple EBS volumes

```hcl
# Fast NVMe for database files
resource "utho_cloud_ebs" "db_data" {
  cloud_id = utho_cloud.db.id
  disk     = 500
  type     = "nvme"
}

# SSD for logs
resource "utho_cloud_ebs" "db_logs" {
  cloud_id = utho_cloud.db.id
  disk     = 100
  type     = "ssd"
}
```

### Full database server setup

```hcl
resource "utho_ssh_key" "deploy" {
  name   = "deploy-key"
  sshkey = file("~/.ssh/id_ed25519.pub")
}

resource "utho_cloud" "db" {
  hostname        = "db-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10355"   # EBS plan — no included disk
  billingcycle    = "monthly"
  auth            = "option2"
  sshkeys         = utho_ssh_key.deploy.id
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "false"
  cpumodel        = "amd"
}

# Root OS volume
resource "utho_cloud_ebs" "os" {
  cloud_id = utho_cloud.db.id
  disk     = 80
  type     = "nvme"
}

# Data volume for PostgreSQL
resource "utho_cloud_ebs" "pgdata" {
  cloud_id = utho_cloud.db.id
  disk     = 500
  type     = "nvme"
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `cloud_id` | String | Yes      | Cloud instance ID. Changing this forces a new resource. |
| `disk`     | Number | Yes      | Volume size in GB. Minimum: 20 GB. |
| `type`     | String | Yes      | Volume type: `nvme` (high performance) or `ssd` (standard). |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique EBS volume ID. |

## Volume Types

| Type   | Description | Best For |
|--------|-------------|----------|
| `nvme` | High Performance NVMe — fastest IOPS | Databases, high-traffic apps |
| `ssd`  | Standard SSD — balanced performance and cost | Logs, backups, general storage |

## Notes

- EBS volumes are persistent — they survive instance reboots and stops.
- Volumes are attached to the instance automatically. Mount them inside the OS using standard Linux disk tools (`lsblk`, `mount`, `fstab`).
- Destroying this resource detaches and deletes the volume permanently — back up your data first.
- EBS plans (`disk = "0"` in plan list) require at least one EBS volume for the OS disk.
- You can attach EBS volumes at creation time using the `ebs` block in `utho_cloud`, or post-creation using this resource.
