---
page_title: "IAM User - Utho"
subcategory: "Account / IAM"
description: |-
  Create and manage Utho IAM sub-users with granular permissions.
---

# utho_iam_user

Creates and manages a Utho IAM sub-user. Sub-users are team members invited to access your Utho account with specific permissions per service. Each permission can be granted at read, write, or delete level.

~> **Note:** After creation the sub-user receives an email invitation. Their status will be `Pending` until they accept.

## Example Usage

### Read-only sub-user

```hcl
resource "utho_iam_user" "readonly" {
  fullname  = "John Doe"
  email     = "john@mycompany.com"
  mobilecc  = "91"
  mobile    = "9876543210"
  resources = "all"
  permissions = "compute_read,kubernetes_read,database_read,objectstorage_read,loadbalancer_read,vpc_read,firewall_read,dns_read,monitoring_read"
}
```

### Full access sub-user

```hcl
resource "utho_iam_user" "admin" {
  fullname  = "Jane Smith"
  email     = "jane@mycompany.com"
  mobilecc  = "91"
  mobile    = "9876543211"
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
    "monitoring_read", "monitoring_write", "monitoring_delete",
  ])
}
```

### DevOps sub-user (compute + K8s + monitoring)

```hcl
resource "utho_iam_user" "devops" {
  fullname  = "DevOps Engineer"
  email     = "devops@mycompany.com"
  mobilecc  = "91"
  mobile    = "9571054173"
  resources = "all"
  permissions = join(",", [
    "compute_read", "compute_write", "compute_delete",
    "kubernetes_read", "kubernetes_write", "kubernetes_delete",
    "autoscaling_read", "autoscaling_write", "autoscaling_delete",
    "monitoring_read", "monitoring_write", "monitoring_delete",
    "loadbalancer_read", "loadbalancer_write", "loadbalancer_delete",
    "vpc_read", "vpc_write", "vpc_delete",
    "firewall_read", "firewall_write", "firewall_delete",
    "sshkey_read", "sshkey_write", "sshkey_delete",
  ])
}
```

### Update permissions

```hcl
resource "utho_iam_user" "devops" {
  fullname  = "DevOps Engineer"
  email     = "devops@mycompany.com"
  mobilecc  = "91"
  mobile    = "9571054173"
  resources = "all"
  # Add database access
  permissions = join(",", [
    "compute_read", "compute_write", "compute_delete",
    "kubernetes_read", "kubernetes_write", "kubernetes_delete",
    "database_read", "database_write",   # ← added
    "monitoring_read",
  ])
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `fullname`    | String | Yes      | Full name of the sub-user. Changing this forces a new resource. |
| `email`       | String | Yes      | Email address. An invitation is sent to this address. Changing this forces a new resource. |
| `mobilecc`    | String | Yes      | Mobile country code (e.g. `91` for India). Changing this forces a new resource. |
| `mobile`      | String | Yes      | Mobile number. Changing this forces a new resource. |
| `permissions` | String | Yes      | Comma-separated permission list. See [Permissions](#permissions). Updatable. |
| `resources`   | String | No       | Resource scope: `all` or specific resource IDs. Default: `all`. Updatable. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique sub-user ID. |
| `status`     | String | Invitation status: `Pending` or `Active`. |
| `date_added` | String | Timestamp when the user was invited. |

## Permissions

Permissions follow the format `{service}_{action}` where action is `read`, `write`, or `delete`.

### Available Services

| Service             | Permissions |
|---------------------|-------------|
| `compute`           | `compute_read`, `compute_write`, `compute_delete` |
| `kubernetes`        | `kubernetes_read`, `kubernetes_write`, `kubernetes_delete` |
| `autoscaling`       | `autoscaling_read`, `autoscaling_write`, `autoscaling_delete` |
| `database`          | `database_read`, `database_write`, `database_delete` |
| `objectstorage`     | `objectstorage_read`, `objectstorage_write`, `objectstorage_delete` |
| `loadbalancer`      | `loadbalancer_read`, `loadbalancer_write`, `loadbalancer_delete` |
| `vpc`               | `vpc_read`, `vpc_write`, `vpc_delete` |
| `firewall`          | `firewall_read`, `firewall_write`, `firewall_delete` |
| `dns`               | `dns_read`, `dns_write`, `dns_delete` |
| `ssl`               | `ssl_read`, `ssl_write`, `ssl_delete` |
| `monitoring`        | `monitoring_read`, `monitoring_write`, `monitoring_delete` |
| `vpn`               | `vpn_read`, `vpn_write`, `vpn_delete` |
| `snapshot`          | `snapshot_read`, `snapshot_write`, `snapshot_delete` |
| `backup`            | `backup_read`, `backup_write`, `backup_delete` |
| `sshkey`            | `sshkey_read`, `sshkey_write`, `sshkey_delete` |
| `api`               | `api_read`, `api_write`, `api_delete` |
| `billing`           | `billing_read`, `billing_write`, `billing_delete` |
| `user`              | `user_read`, `user_write`, `user_delete` |
| `ebs`               | `ebs_read`, `ebs_write`, `ebs_delete` |
| `ipsec`             | `ipsec_read`, `ipsec_write`, `ipsec_delete` |
| `container_registry`| `container_registry_read`, `container_registry_write`, `container_registry_delete` |

## Notes

- After creation the sub-user receives an email invitation — status will be `Pending` until accepted.
- Only `permissions` and `resources` are updatable in place.
- Changing `fullname`, `email`, `mobilecc`, or `mobile` destroys and recreates the user.
- Use `join(",", [...])` in Terraform to build the permissions string cleanly.