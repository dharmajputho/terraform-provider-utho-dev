---
page_title: "Managed Database - Utho"
subcategory: "Database"
description: |-
  Create and manage Utho Managed Database clusters (PostgreSQL, Redis).
---

# utho_database

Creates and manages a Utho Managed Database cluster. Supports PostgreSQL (`pg`) and Redis (`redis`) engines. Includes built-in high availability via replica nodes, automated backups, and point-in-time recovery.

## Example Usage

### PostgreSQL cluster

```hcl
resource "utho_database" "postgres" {
  cluster_name    = "production-pg"
  dcslug          = "inmumbaizone2"
  engine          = "pg"
  version         = "17"
  size            = "10157"
  network_type    = "public"
  billing         = "monthly"
  pitr_enabled    = "1"
  replica_count   = "1"
}

output "connection_host" {
  value = "public-primary-pg-inmumbaizone2-${utho_database.postgres.id}.db.onutho.com"
}

output "db_password" {
  value     = utho_database.postgres.default_pass
  sensitive = true
}
```

### PostgreSQL inside a VPC with security group

```hcl
resource "utho_database" "postgres" {
  cluster_name  = "private-pg"
  dcslug        = "inmumbaizone2"
  engine        = "pg"
  version       = "17"
  size          = "10157"
  network_type  = "private"
  billing       = "monthly"
  vpc           = utho_subnet.private.id
  firewall      = utho_firewall.db.id
  pitr_enabled  = "1"
  replica_count = "1"
}
```

### Add a database and user after cluster creation

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

| Argument        | Type   | Required | Description |
|-----------------|--------|----------|-------------|
| `cluster_name`  | String | Yes      | Cluster label. Changing this forces a new resource. |
| `dcslug`        | String | Yes      | Data center slug. Changing this forces a new resource. |
| `engine`        | String | Yes      | Database engine: `pg` (PostgreSQL) or `redis`. Changing this forces a new resource. |
| `version`       | String | Yes      | Database version (e.g. `17` for PostgreSQL 17). Changing this forces a new resource. |
| `size`          | String | Yes      | Plan ID for node size. Updatable via resize. |
| `network_type`  | String | Yes      | Network type: `public` or `private`. Changing this forces a new resource. |
| `billing`       | String | Yes      | Billing cycle: `monthly` or `hourly`. |
| `pitr_enabled`  | String | No       | Enable point-in-time recovery: `1` or `0`. |
| `replica_count` | String | No       | Number of replica nodes to deploy at creation. |
| `vpc`           | String | No       | VPC subnet ID. Changing this forces a new resource. |
| `firewall`      | String | No       | Security group ID to attach. |

## Attribute Reference

| Attribute       | Type   | Description |
|-----------------|--------|-------------|
| `id`            | String | Unique cluster ID. |
| `cloud_id`      | String | Primary node cloud ID (required for some operations). |
| `status`        | String | Cluster status. |
| `default_user`  | String | Default admin username. |
| `default_pass`  | String | Default admin password. **Sensitive.** |
| `default_dbname`| String | Default database name. |
| `port`          | String | Database port. |
| `created_at`    | String | Creation timestamp. |

## Connection String Format

```
postgres://{default_user}:{default_pass}@public-primary-pg-{dcslug}-{id}-{cloud_id}.db.onutho.com:{port}/{default_dbname}
```
