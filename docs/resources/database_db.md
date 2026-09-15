---
page_title: "Database - Utho"
subcategory: "Database"
description: |-
  Create and manage databases inside a Utho database cluster.
---

# utho_database_db

Creates a database inside a Utho Managed Database cluster.

## Example Usage

### Create databases

```hcl
resource "utho_database_db" "app" {
  cluster_id = utho_database.postgres.id
  name       = "appdb"
}

resource "utho_database_db" "analytics" {
  cluster_id = utho_database.postgres.id
  name       = "analyticsdb"
}
```

### Create database and assign user permissions

```hcl
resource "utho_database_db" "app" {
  cluster_id = utho_database.postgres.id
  name       = "appdb"
}

resource "utho_database_user" "app" {
  cluster_id = utho_database.postgres.id
  name       = "appuser"
}
```

## Argument Reference

| Argument     | Type   | Required | Description |
|--------------|--------|----------|-------------|
| `cluster_id` | String | Yes      | Database cluster ID. Changing this forces a new resource. |
| `name`       | String | Yes      | Database name. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{cluster_id}:{name}`. |
