---
page_title: "Database Connection Pool - Utho"
subcategory: "Database"
description: |-
  Create and manage connection pools in a Utho database cluster.
---

# utho_database_pool

Creates a PgBouncer connection pool on a Utho Managed PostgreSQL cluster. Connection pooling reduces the overhead of establishing new database connections from your application.

## Example Usage

### Transaction mode pool

```hcl
resource "utho_database_pool" "app" {
  cluster_id = utho_database.postgres.id
  cloud_id   = utho_database.postgres.cloud_id
  name       = "app-pool"
  db         = "appdb"
  user       = "appuser"
  mode       = "transaction"
  size       = 25
}
```

### Session mode pool

```hcl
resource "utho_database_pool" "session" {
  cluster_id = utho_database.postgres.id
  cloud_id   = utho_database.postgres.cloud_id
  name       = "session-pool"
  db         = "appdb"
  user       = "dbadmin"
  mode       = "session"
  size       = 10
}
```

## Argument Reference

| Argument     | Type   | Required | Description |
|--------------|--------|----------|-------------|
| `cluster_id` | String | Yes      | Database cluster ID. Changing this forces a new resource. |
| `cloud_id`   | String | Yes      | Primary node cloud ID. Get this from `utho_database.cloud_id`. Changing this forces a new resource. |
| `name`       | String | Yes      | Connection pool name. Changing this forces a new resource. |
| `db`         | String | Yes      | Database name to pool connections for. Changing this forces a new resource. |
| `user`       | String | Yes      | Database user for the pool. Changing this forces a new resource. |
| `mode`       | String | Yes      | Pool mode: `transaction`, `session`, or `statement`. Changing this forces a new resource. |
| `size`       | Number | Yes      | Maximum number of connections in the pool. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique connection pool ID assigned by Utho. |

## Pool Modes

| Mode          | Description |
|---------------|-------------|
| `transaction` | A server connection is assigned for the duration of a transaction. Recommended for most applications. |
| `session`     | A server connection is assigned for the duration of a client session. |
| `statement`   | A server connection is assigned per statement. Only for autocommit mode. |
