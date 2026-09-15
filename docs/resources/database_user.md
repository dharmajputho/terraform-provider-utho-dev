---
page_title: "Database User - Utho"
subcategory: "Database"
description: |-
  Create and manage users in a Utho database cluster.
---

# utho_database_user

Creates a user in a Utho Managed Database cluster. A strong password is auto-generated if not provided — save it immediately as it cannot be retrieved later.

## Example Usage

### Create a user with auto-generated password

```hcl
resource "utho_database_user" "app" {
  cluster_id = utho_database.postgres.id
  name       = "appuser"
}

output "db_password" {
  value     = utho_database_user.app.generated_password
  sensitive = true
}
```

### Create a user with a specific password

```hcl
resource "utho_database_user" "app" {
  cluster_id = utho_database.postgres.id
  name       = "appuser"
  password   = var.db_password
}
```

## Argument Reference

| Argument     | Type   | Required | Description |
|--------------|--------|----------|-------------|
| `cluster_id` | String | Yes      | Database cluster ID. Changing this forces a new resource. |
| `name`       | String | Yes      | Username. Changing this forces a new resource. |
| `password`   | String | No       | Password. Leave empty to auto-generate. **Sensitive.** Changing this forces a new resource. |

## Attribute Reference

| Attribute            | Type   | Description |
|----------------------|--------|-------------|
| `id`                 | String | Identifier in the format `{cluster_id}:{name}`. |
| `generated_password` | String | Auto-generated password. **Sensitive — shown only at creation time.** |

~> **Important:** Save `generated_password` immediately after creation. It cannot be retrieved later. Mark the output as `sensitive = true`.
