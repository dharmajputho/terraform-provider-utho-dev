---
page_title: "IAM User - Utho"
subcategory: "Account / IAM"
description: |-
  Invite and manage sub-users with granular permissions on your Utho account.
---

# utho_iam_user

Invites a sub-user to your Utho account with specific permissions. Sub-users are team members who can access and manage resources based on the permissions you grant — scoped per service and per action (read, write, delete).

The invited user receives an email invitation and must accept it to activate their account. Their status will be `Pending` until they accept.

## Example Usage

### DevOps engineer — compute and Kubernetes access

```hcl
resource "utho_iam_user" "devops" {
  fullname  = "Rahul Sharma"
  email     = "rahul@mycompany.com"
  mobilecc  = "91"
  mobile    = "9876543210"
  resources = "all"
  permissions = join(",", [
    "compute_read", "compute_write", "compute_delete",
    "kubernetes_read", "kubernetes_write", "kubernetes_delete",
    "autoscaling_read", "autoscaling_write", "autoscaling_delete",
    "loadbalancer_read", "loadbalancer_write", "loadbalancer_delete",
    "vpc_read", "vpc_write", "vpc_delete",
    "firewall_read", "firewall_write", "firewall_delete",
    "monitoring_read", "monitoring_write",
    "sshkey_read", "sshkey_write", "sshkey_delete",
  ])
}
```

### Read-only access for auditors

```hcl
resource "utho_iam_user" "auditor" {
  fullname  = "Security Auditor"
  email     = "audit@mycompany.com"
  mobilecc  = "91"
  mobile    = "9571054173"
  resources = "all"
  permissions = join(",", [
    "compute_read",
    "kubernetes_read",
    "database_read",
    "objectstorage_read",
    "loadbalancer_read",
    "vpc_read",
    "firewall_read",
    "dns_read",
    "monitoring_read",
    "billing_read",
    "api_read",
    "activity_read",
  ])
}
```

### Full admin access

```hcl
resource "utho_iam_user" "admin" {
  fullname  = "Platform Engineer"
  email     = "platform@mycompany.com"
  mobilecc  = "91"
  mobile    = "9000000000"
  resources = "all"
  permissions = join(",", [
    "compute_read", "compute_write", "compute_delete",
    "kubernetes_read", "kubernetes_write", "kubernetes_delete",
    "database_read", "database_write", "database_delete",
    "objectstorage_read", "objectstorage_write", "objectstorage_delete",
    "loadbalancer_read", "loadbalancer_write", "loadbalancer_delete",
    "vpc_read", "vpc_write", "vpc_delete",
    "firewall_read", "firewall_write", "firewall_delete",
    "dns_read", "dns_write", "dns_delete",
    "ssl_read", "ssl_write", "ssl_delete",
    "monitoring_read", "monitoring_write", "monitoring_delete",
    "vpn_read", "vpn_write", "vpn_delete",
    "snapshot_read", "snapshot_write", "snapshot_delete",
    "backup_read", "backup_write", "backup_delete",
    "sshkey_read", "sshkey_write", "sshkey_delete",
    "api_read", "api_write", "api_delete",
    "user_read", "user_write", "user_delete",
  ])
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `fullname`    | String | Yes      | User's full name. Changing this forces a new resource. |
| `email`       | String | Yes      | Email address — the invitation is sent here. Changing this forces a new resource. |
| `mobilecc`    | String | Yes      | Country code without `+` (e.g. `"91"` for India). Changing this forces a new resource. |
| `mobile`      | String | Yes      | Mobile number. Changing this forces a new resource. |
| `permissions` | String | Yes      | Comma-separated permissions. Format: `{service}_{action}`. Updatable in place. |
| `resources`   | String | No       | Resource scope: `"all"` or specific IDs. Default: `"all"`. Updatable. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique sub-user ID. Use this in `utho_project_member`. |
| `status`     | String | `Pending` (invitation not accepted) or `Active`. |
| `date_added` | String | Timestamp when the invitation was sent. |

## Available Permissions

Format: `{service}_{read|write|delete}`

| Service | Read | Write | Delete |
|---------|------|-------|--------|
| `compute` | ✓ | ✓ | ✓ |
| `kubernetes` | ✓ | ✓ | ✓ |
| `autoscaling` | ✓ | ✓ | ✓ |
| `database` | ✓ | ✓ | ✓ |
| `objectstorage` | ✓ | ✓ | ✓ |
| `loadbalancer` | ✓ | ✓ | ✓ |
| `vpc` | ✓ | ✓ | ✓ |
| `firewall` | ✓ | ✓ | ✓ |
| `dns` | ✓ | ✓ | ✓ |
| `ssl` | ✓ | ✓ | ✓ |
| `monitoring` | ✓ | ✓ | ✓ |
| `vpn` | ✓ | ✓ | ✓ |
| `snapshot` | ✓ | ✓ | ✓ |
| `backup` | ✓ | ✓ | ✓ |
| `sshkey` | ✓ | ✓ | ✓ |
| `api` | ✓ | ✓ | ✓ |
| `billing` | ✓ | ✓ | ✓ |
| `user` | ✓ | ✓ | ✓ |
| `ebs` | ✓ | ✓ | ✓ |
| `ipsec` | ✓ | ✓ | ✓ |
| `container_registry` | ✓ | ✓ | ✓ |

## Notes

- The invitation email expires — if the user doesn't accept, delete and recreate the resource to resend.
- Only `permissions` and `resources` are updatable without recreating the user.
- Changing `email`, `fullname`, `mobilecc`, or `mobile` destroys the user record and resends an invitation.
- Use `join(",", [...])` in Terraform to build the permissions string cleanly from a list.
