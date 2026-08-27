---
page_title: "utho_cloud Resource - utho"
subcategory: "Compute"
description: |-
  Create and manage Utho Cloud instances.
---

# utho_cloud

Creates and manages a Utho Cloud instance (virtual machine).

Supports deployment from OS images, snapshots, backups, ISOs, and marketplace stacks.
Supports password and SSH key authentication, VPC networking, EBS volumes, and firewall
attachment at creation time.

## Example Usage

### Basic instance with password authentication

```hcl
resource "utho_cloud" "web" {
  hostname      = "web-01.mhc"
  dcslug        = "inmumbaizone2"
  planid        = "10360"
  billingcycle  = "hourly"
  auth          = "option1"
  root_password = "StrongP@ssw0rd!"
  image         = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
  cpumodel      = "amd"
  support       = "unmanaged"
}
```

### Instance with SSH key authentication

```hcl
resource "utho_cloud" "app" {
  hostname        = "app-01.mhc"
  dcslug          = "innoida"
  planid          = "10315"
  billingcycle    = "monthly"
  auth            = "option2"
  sshkeys         = "74133628"
  image           = "ubuntu-25.04-x86_64"
  enable_publicip = "true"
  cpumodel        = "intel"
  firewall        = "23437038"
  support         = "unmanaged"
}
```

### Multiple instances with count

```hcl
resource "utho_cloud" "workers" {
  count = 5

  hostname        = "worker-${count.index + 1}.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option1"
  root_password   = "StrongP@ssw0rd!"
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
}

output "worker_ips" {
  value = utho_cloud.workers[*].ip
}
```

### Instance inside a VPC with EBS

```hcl
resource "utho_cloud" "db" {
  hostname        = "db-01.mhc"
  dcslug          = "innoida"
  planid          = "10355"
  billingcycle    = "12month"
  auth            = "option1"
  root_password   = "StrongP@ssw0rd!"
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "false"
  vpc             = "3a276678-6c71-4203-a788-905930a728be"
  firewall        = "23436544"

  ebs = [
    {
      id   = "1"
      disk = 80
      type = "nvme"
    }
  ]
}
```

### Deploy from snapshot

```hcl
resource "utho_cloud" "restored" {
  hostname        = "restored-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option1"
  root_password   = "StrongP@ssw0rd!"
  snapshotid      = "snap-7781"
  enable_publicip = "true"
}
```

## Argument Reference

### Required

| Argument       | Type   | Description |
|----------------|--------|-------------|
| `hostname`     | String | Hostname for the instance. Must be unique within your account. |
| `dcslug`       | String | Data center slug. See [Data Centers](#data-centers) below. |
| `planid`       | String | Plan ID for the instance size. |
| `billingcycle` | String | Billing cycle. Accepted values: `hourly`, `monthly`, `12month`. |
| `auth`         | String | Authentication method. `option1` = password, `option2` = SSH key. |

### Optional

| Argument          | Type   | Description |
|-------------------|--------|-------------|
| `root_password`   | String | Root password. Required when `auth = option1`. **Sensitive** — never shown in logs. |
| `sshkeys`         | String | SSH key ID. Required when `auth = option2`. |
| `image`           | String | OS image slug (e.g. `ubuntu-22.04-x86_64`). Required for standard deploys. |
| `enable_publicip` | String | Assign a public IP: `true` or `false`. Defaults to `true`. |
| `cpumodel`        | String | CPU architecture. Accepted values: `amd`, `intel`. |
| `enablebackup`    | String | Enable automated backups: `true` or `false`. |
| `support`         | String | Support level. Accepted values: `unmanaged`, `managed`. |
| `firewall`        | String | Security group ID to attach at creation time. |
| `vpc`             | String | VPC subnet ID for deploying the instance inside a private network. |
| `subnet_required` | String | Whether subnet attachment is required: `true` or `false`. |
| `snapshotid`      | String | Snapshot ID to deploy from an existing snapshot. |
| `backupid`        | String | Backup ID to restore from an existing backup. |
| `iso`             | String | ISO ID to deploy from a custom ISO image. |
| `stack`           | String | Stack ID to deploy from a marketplace or custom stack. |
| `delete_ebs`      | Bool   | Whether to delete attached EBS volumes when the instance is destroyed. Default: `false`. |
| `ebs`             | List   | List of EBS volumes to attach at creation time. See [EBS Block](#ebs-block) below. |

### EBS Block

The `ebs` block supports the following arguments:

| Argument | Type   | Required | Description |
|----------|--------|----------|-------------|
| `id`     | String | Yes      | EBS volume index (e.g. `"1"`, `"2"`). |
| `disk`   | Number | Yes      | Disk size in GB. |
| `type`   | String | Yes      | Disk type. Accepted values: `nvme`, `ssd`. |

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

| Attribute      | Type   | Description |
|----------------|--------|-------------|
| `id`           | String | Unique cloud instance ID assigned by Utho. |
| `ip`           | String | Primary public IP address of the instance. |
| `status`       | String | Current instance status (e.g. `Active`). |
| `power_status` | String | Current power state (e.g. `Running`, `Not_Running`). |
| `created_at`   | String | Timestamp when the instance was created (UTC). |

## Import

Existing cloud instances can be imported into Terraform state using
the instance ID from the Utho dashboard.

```bash
terraform import utho_cloud.web 1671990
```

After importing, add the required fields (`planid`, `auth`) to your
`.tf` file and run `terraform plan` to reconcile the state.

```hcl
resource "utho_cloud" "web" {
  hostname      = "web-01.mhc"
  dcslug        = "inmumbaizone2"
  planid        = "10360"   # add after import
  billingcycle  = "hourly"
  auth          = "option1" # add after import
}
```

## Data Centers

| Slug             | Location       |
|------------------|----------------|
| `innoida`        | Noida, India   |
| `inmumbaizone2`  | Mumbai, India  |
| `inbangalore`    | Bangalore, India |
