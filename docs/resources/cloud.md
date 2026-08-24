---
page_title: "utho_cloud Resource"
description: |-
  Create and manage Utho Cloud instances.
---

# utho_cloud

Create and manage Utho Cloud instances.

## Example Usage

```hcl
resource "utho_cloud" "server" {
  hostname        = "my-server.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option1"
  root_password   = "StrongP@ssw0rd!"
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
  cpumodel        = "amd"
  enablebackup    = "false"
  support         = "unmanaged"
}

output "server_ip" {
  value = utho_cloud.server.ip
}
```

## Multiple Servers

```hcl
resource "utho_cloud" "servers" {
  count = 3

  hostname        = "server-${count.index + 1}.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option1"
  root_password   = "StrongP@ssw0rd!"
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
}
```

## Argument Reference

### Required

- `hostname` - Hostname for the instance. Must be unique.
- `dcslug` - Data center slug (e.g. `innoida`, `inmumbaizone2`).
- `planid` - Plan ID for the instance size.
- `billingcycle` - Billing cycle: `hourly`, `monthly`, `12month`.
- `auth` - Auth method: `option1` (password) or `option2` (SSH key).

### Optional

- `root_password` - Root password. Required when `auth = option1`. Sensitive.
- `sshkeys` - SSH key ID. Required when `auth = option2`.
- `image` - OS image slug (e.g. `ubuntu-22.04-x86_64`).
- `enable_publicip` - Enable public IP: `true` or `false`.
- `cpumodel` - CPU model: `amd` or `intel`.
- `enablebackup` - Enable backup: `true` or `false`.
- `support` - Support type: `unmanaged` or `managed`.
- `firewall` - Firewall/security group ID.
- `vpc` - VPC ID for deploying inside a VPC.
- `snapshotid` - Snapshot ID to deploy from snapshot.
- `backupid` - Backup ID to restore from backup.
- `iso` - ISO ID to deploy from custom ISO.
- `stack` - Stack ID to deploy from marketplace stack.
- `delete_ebs` - Delete EBS volumes on destroy. Default: `false`.
- `ebs` - List of EBS volumes to attach. See [EBS](#ebs) below.

### EBS

- `id` - EBS volume index (e.g. `1`, `2`).
- `disk` - Disk size in GB.
- `type` - Disk type: `nvme` or `ssd`.

## Attribute Reference

- `id` - Unique cloud instance ID assigned by Utho.
- `ip` - Public IP address of the instance.
- `status` - Current status (e.g. `Active`).
- `power_status` - Power status (e.g. `Running`).
- `created_at` - Creation timestamp.

## Import

Existing cloud instances can be imported using the cloud ID:

```bash
terraform import utho_cloud.server 1671914
```