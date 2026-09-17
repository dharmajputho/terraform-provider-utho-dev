---
page_title: "Managed Database - Utho"
subcategory: "Database"
description: |-
  Create and manage Utho Managed PostgreSQL database clusters.
---

# utho_database

Creates and manages a Utho Managed Database cluster. Utho handles provisioning, OS patching, automated backups, and high availability — you get a connection string and start using it.

Currently supports PostgreSQL (`pg`). After creating a cluster, add databases with `utho_database_db`, users with `utho_database_user`, and connection pools with `utho_database_pool`.

## Example Usage

### Minimal PostgreSQL cluster

The quickest way to get a managed database. Utho generates a default admin user and password.

```hcl
resource "utho_database" "pg" {
  cluster_name = "my-app-db"
  dcslug       = "inmumbaizone2"
  engine       = "pg"
  version      = "17"
  size         = "10157"
  network_type = "public"
  billing      = "monthly"
}

# Connection string output for your application
output "db_connection" {
  value     = "postgres://${utho_database.pg.default_user}:PASSWORD@public-primary-pg-inmumbaizone2-${utho_database.pg.id}.db.onutho.com:5432/"
  sensitive = false
}
output "db_password" {
  value     = utho_database.pg.default_pass
  sensitive = true
}
```

### Production cluster with HA replica and PITR

For production: one primary + one replica for high availability, point-in-time recovery enabled for data protection.

```hcl
resource "utho_database" "prod" {
  cluster_name  = "production-db"
  dcslug        = "inmumbaizone2"
  engine        = "pg"
  version       = "17"
  size          = "10157"
  network_type  = "public"
  billing       = "monthly"
  pitr_enabled  = "1"    # point-in-time recovery
  replica_count = "1"    # one standby replica
}
```

### Private cluster inside a VPC with security group

For maximum security — no public endpoint, only accessible from within your VPC. Combined with a security group that only allows your app servers.

```hcl
resource "utho_firewall" "db" { name = "db-sg" }

resource "utho_firewall_rule" "postgres" {
  firewall_id  = utho_firewall.db.id
  type         = "incoming"
  service      = "CUSTOM"
  protocol     = "tcp"
  port         = "5432"
  port_range   = "5432"
  addresses    = "10.0.0.0/8"
  source_range = "10.0.0.0/8"
}

resource "utho_database" "private" {
  cluster_name  = "private-db"
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

### Full database setup — cluster, database, user, and pool

```hcl
resource "utho_database" "main" {
  cluster_name = "app-db"
  dcslug       = "inmumbaizone2"
  engine       = "pg"
  version      = "17"
  size         = "10157"
  network_type = "public"
  billing      = "monthly"
}

resource "utho_database_db" "app" {
  cluster_id = utho_database.main.id
  name       = "appdb"
}

resource "utho_database_user" "app" {
  cluster_id = utho_database.main.id
  name       = "appuser"
}

resource "utho_database_pool" "app" {
  cluster_id = utho_database.main.id
  cloud_id   = utho_database.main.cloud_id
  name       = "app-pool"
  db         = utho_database_db.app.name
  user       = utho_database_user.app.name
  mode       = "transaction"
  size       = 25
}

output "connection_string" {
  value     = "postgres://${utho_database_user.app.name}:PASSWORD@public-primary-pg-inmumbaizone2-${utho_database.main.id}.db.onutho.com:5432/${utho_database_db.app.name}"
  sensitive = false
}
```

## Argument Reference

### Required

| Argument       | Type   | Description |
|----------------|--------|-------------|
| `cluster_name` | String | Cluster label. Changing this forces a new resource. |
| `dcslug`       | String | Data center. Changing this forces a new resource. |
| `engine`       | String | Database engine: `pg` (PostgreSQL). Changing this forces a new resource. |
| `version`      | String | Engine version (e.g. `17` for PostgreSQL 17). Changing this forces a new resource. |
| `size`         | String | Plan ID for node size (CPU/RAM/disk). |
| `network_type` | String | `public` or `private`. Changing this forces a new resource. |
| `billing`      | String | `monthly` or `hourly`. |

### Optional

| Argument        | Type   | Description |
|-----------------|--------|-------------|
| `pitr_enabled`  | String | Enable point-in-time recovery: `"1"` or `"0"`. Recommended for production. |
| `replica_count` | String | Number of standby replica nodes. Use `"1"` for high availability. |
| `vpc`           | String | VPC subnet ID. Required when `network_type = "private"`. Changing this forces a new resource. |
| `firewall`      | String | Security group ID to attach. |

## Attribute Reference

| Attribute        | Type   | Description |
|------------------|--------|-------------|
| `id`             | String | Unique cluster ID. |
| `cloud_id`       | String | Primary node cloud ID. Required by `utho_database_pool`. |
| `status`         | String | Cluster status (`Active`, `Pending`). |
| `default_user`   | String | Auto-created admin username (always `dbadmin`). |
| `default_pass`   | String | Auto-generated admin password. **Sensitive.** Save this immediately. |
| `default_dbname` | String | Default database name. |
| `port`           | String | Database port (`5432` for PostgreSQL). |
| `created_at`     | String | Creation timestamp. |

## Connection String Format

```
postgres://{default_user}:{default_pass}@public-primary-pg-{dcslug}-{id}-{cloud_id}.db.onutho.com:5432/{default_dbname}
```

Example:
```
postgres://dbadmin:mypassword@public-primary-pg-inmumbaizone2-189842-1666717.db.onutho.com:5432/defaultdb
```

## Notes

- Database clusters take 5–10 minutes to provision.
- The `default_pass` is shown in Terraform state. Encrypt your state file or use a remote backend with encryption at rest.
- For production, always enable `pitr_enabled = "1"` and `replica_count = "1"`.
- Use connection pooling (`utho_database_pool`) with PgBouncer to reduce connection overhead in high-traffic apps.
