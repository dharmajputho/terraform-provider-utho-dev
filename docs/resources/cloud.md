---
page_title: "Cloud Instance - Utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Create and manage Utho Cloud instances (virtual machines).
---

# utho_cloud

Creates and manages a Utho Cloud instance — a virtual machine running in one of Utho's data centers. You can deploy from standard OS images, marketplace stacks, snapshots, backups, or custom ISOs.

Once created, the instance ID is used by other resources like `utho_cloud_snapshot`, `utho_cloud_ebs`, and `utho_firewall_server`.

## Example Usage

### Minimal instance with password login

The simplest possible setup — one Ubuntu server with a root password.

```hcl
resource "utho_cloud" "web" {
  hostname        = "web-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option1"
  root_password   = var.root_password
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
}

output "ip" { value = utho_cloud.web.ip }
```

### Instance with SSH key authentication (recommended)

Using SSH keys is more secure than password auth. First import the key with `utho_ssh_key`, then reference it here.

```hcl
resource "utho_ssh_key" "deploy" {
  name   = "deploy-key"
  sshkey = file("~/.ssh/id_ed25519.pub")
}

resource "utho_cloud" "app" {
  hostname        = "app-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option2"
  sshkeys         = utho_ssh_key.deploy.id
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
}
```

### Scale horizontally with count

Create multiple identical servers in one block. Each gets a unique hostname.

```hcl
resource "utho_cloud" "worker" {
  count = 5

  hostname        = "worker-${count.index + 1}.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option2"
  sshkeys         = utho_ssh_key.deploy.id
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
}

output "worker_ips" {
  value = utho_cloud.worker[*].ip
}
```

### Instance inside a private VPC

Attach the instance to a VPC subnet for private networking. Set `enable_publicip = "false"` for fully private instances that communicate only over the VPC.

```hcl
resource "utho_cloud" "backend" {
  hostname        = "backend-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10355"
  billingcycle    = "monthly"
  auth            = "option2"
  sshkeys         = utho_ssh_key.deploy.id
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "false"
  vpc             = utho_subnet.private.id
  firewall        = utho_firewall.backend.id
}
```

### Instance with additional EBS storage

Attach extra block volumes at creation time for databases, logs, or large datasets.

```hcl
resource "utho_cloud" "db" {
  hostname        = "db-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10355"
  billingcycle    = "monthly"
  auth            = "option2"
  sshkeys         = utho_ssh_key.deploy.id
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "false"

  ebs = [
    { id = "1", disk = 100, type = "nvme" },
    { id = "2", disk = 500, type = "ssd"  },
  ]
}
```

### Restore from a snapshot

Useful for deploying pre-configured golden images or recovering from a snapshot.

```hcl
resource "utho_cloud" "restored" {
  hostname        = "restored-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option1"
  root_password   = var.root_password
  snapshotid      = "snap-7781"
  enable_publicip = "true"
}
```

## Argument Reference

### Required

| Argument       | Type   | Description |
|----------------|--------|-------------|
| `hostname`     | String | Hostname for the instance. |
| `dcslug`       | String | Data center. See [Data Centers](#data-centers). |
| `planid`       | String | Plan ID for the instance size (CPU/RAM/disk). |
| `billingcycle` | String | `hourly`, `monthly`, or `12month`. |
| `auth`         | String | `option1` = root password, `option2` = SSH key. |

### Authentication — one required

| Argument       | Type   | When Required |
|----------------|--------|---------------|
| `root_password` | String | When `auth = "option1"`. **Sensitive.** |
| `sshkeys`      | String | When `auth = "option2"`. SSH key ID from `utho_ssh_key`. |

### Image source — one required

| Argument     | Type   | Description |
|--------------|--------|-------------|
| `image`      | String | OS image slug (e.g. `ubuntu-22.04-x86_64`). |
| `snapshotid` | String | Deploy from an existing snapshot. |
| `backupid`   | String | Deploy from an existing backup. |
| `iso`        | String | Deploy from a custom ISO. |
| `stack`      | String | Deploy from a marketplace or custom stack. |

### Optional

| Argument          | Type   | Description |
|-------------------|--------|-------------|
| `enable_publicip` | String | `"true"` or `"false"`. Default: `"true"`. |
| `vpc`             | String | VPC subnet ID. Attaches the instance to a private network. |
| `firewall`        | String | Security group ID. Attaches at creation time. |
| `cpumodel`        | String | CPU preference: `amd` or `intel`. |
| `enablebackup`    | String | Enable automated backups: `"true"` or `"false"`. |
| `support`         | String | `unmanaged` or `managed`. |
| `delete_ebs`      | Bool   | Delete attached EBS volumes on destroy. Default: `false`. |
| `ebs`             | List   | EBS volumes to attach at creation. See [EBS Block](#ebs-block). |

### EBS Block

```hcl
ebs = [
  { id = "1", disk = 100, type = "nvme" }
]
```

| Argument | Type   | Description |
|----------|--------|-------------|
| `id`     | String | Sequential identifier (`"1"`, `"2"`, etc.). |
| `disk`   | Number | Disk size in GB. |
| `type`   | String | `nvme` (faster) or `ssd`. |

## Attribute Reference

| Attribute      | Type   | Description |
|----------------|--------|-------------|
| `id`           | String | Unique instance ID. Used to reference this instance in other resources. |
| `ip`           | String | Primary public IP address. |
| `status`       | String | Instance status (e.g. `Active`). |
| `power_status` | String | Power state (`Running`, `Not_Running`). |
| `created_at`   | String | Creation timestamp (UTC). |

## Import

Bring an existing instance under Terraform management without recreating it.

```bash
terraform import utho_cloud.web 1671990
```

After importing, add the required arguments to your `.tf` file and run `terraform plan` to sync state:

```hcl
resource "utho_cloud" "web" {
  hostname     = "web-01.mhc"
  dcslug       = "inmumbaizone2"
  planid       = "10360"
  billingcycle = "hourly"
  auth         = "option2"
  sshkeys      = utho_ssh_key.deploy.id
}
```

## Data Centers

| Slug | Location |
|------|----------|
| `innoida` | Delhi (Noida), India |
| `inmumbaizone2` | Mumbai, India |
| `inbangalore` | Bangalore, India |
