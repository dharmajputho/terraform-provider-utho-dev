---
page_title: "SSH Key - Utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Import and manage SSH keys on Utho Cloud.
---

# utho_ssh_key

Imports and manages an SSH key on Utho Cloud. Once imported, the key ID
can be referenced in `utho_cloud` resources to enable SSH key authentication
instead of root password authentication.

## Example Usage

### Import an SSH key

```hcl
resource "utho_ssh_key" "main" {
  name   = "my-deploy-key"
  sshkey = "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAI... user@host"
}
0
output "key_id" {
  value = utho_ssh_key.main.id
}
```

### Read key from file

```hcl
resource "utho_ssh_key" "main" {
  name   = "my-deploy-key"
  sshkey = file("~/.ssh/id_ed25519.pub")
}
```

### Use with a cloud instance

```hcl
resource "utho_ssh_key" "main" {
  name   = "my-deploy-key"
  sshkey = file("~/.ssh/id_ed25519.pub")
}

resource "utho_cloud" "app" {
  hostname        = "app-01.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option2"
  sshkeys         = utho_ssh_key.main.id
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
}
```

### Multiple servers sharing the same key

```hcl
resource "utho_ssh_key" "deploy" {
  name   = "deploy-key"
  sshkey = file("~/.ssh/id_ed25519.pub")
}

resource "utho_cloud" "servers" {
  count    = 5
  hostname = "server-${count.index + 1}.mhc"
  dcslug   = "inmumbaizone2"
  planid   = "10360"
  billingcycle = "hourly"
  auth     = "option2"
  sshkeys  = utho_ssh_key.deploy.id
  image    = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
}
```

## Argument Reference

| Argument | Type   | Required | Description |
|----------|--------|----------|-------------|
| `name`   | String | Yes      | Name of the SSH key. Changing this forces a new resource. |
| `sshkey` | String | Yes      | Public SSH key content. Accepted formats: `ssh-rsa`, `ssh-ed25519`, `ecdsa-sha2-nistp256`. Changing this forces a new resource. Sensitive — not shown in logs. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique SSH key ID assigned by Utho. Use this as `sshkeys` in `utho_cloud`. |
| `created_at` | String | Timestamp when the key was imported (UTC). |

## Notes

- Only the **public key** is imported — never the private key.
- SSH keys cannot be updated in place. Destroy and recreate to replace a key.
- If two keys have the same name, the most recently created one is used.
- After deleting an SSH key, any instances using it will retain SSH access
  until they are rebuilt — the key is only required at instance creation time.