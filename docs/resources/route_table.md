---
page_title: "Route Table - Utho"
subcategory: "Networking / VPC"
description: |-
  Create and manage route tables for Utho VPCs.
---

# utho_route_table

Creates and manages a route table for a Utho VPC. Route tables control
how network traffic is directed within your VPC. Individual routes are
managed using the `utho_route` resource.

When created, the route table is automatically attached to the VPC.
When destroyed, it is automatically detached before deletion.

## Example Usage

### Basic route table

```hcl
resource "utho_route_table" "main" {
  name    = "main-routes"
  vpc_id  = utho_vpc.main.id
  dcslug  = "inmumbaizone2"
}
```

### Route table with internet gateway route

```hcl
resource "utho_route_table" "main" {
  name   = "main-routes"
  vpc_id = utho_vpc.main.id
  dcslug = "inmumbaizone2"
}

resource "utho_route" "internet" {
  route_table_id         = utho_route_table.main.id
  destination_cidr_block = "0.0.0.0/0"
  route_type             = "igw"
  target                 = "igw"
}
```

### Route table with multiple routes

```hcl
resource "utho_route_table" "main" {
  name   = "main-routes"
  vpc_id = utho_vpc.main.id
  dcslug = "inmumbaizone2"
}

resource "utho_route" "internet" {
  route_table_id         = utho_route_table.main.id
  destination_cidr_block = "0.0.0.0/0"
  route_type             = "igw"
  target                 = "igw"
}

resource "utho_route" "internal" {
  route_table_id         = utho_route_table.main.id
  destination_cidr_block = "10.0.0.0/16"
  route_type             = "local"
  target                 = "local"
}
```

## Argument Reference

| Argument  | Type   | Required | Description |
|-----------|--------|----------|-------------|
| `name`    | String | Yes      | Route table name. |
| `vpc_id`  | String | Yes      | VPC ID to associate this route table with. Changing this forces a new resource. |
| `dcslug`  | String | Yes      | Data center slug. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique route table ID (UUID). |

## Notes

- The route table is automatically attached to the VPC on creation
  and detached on destruction.
- Add routes to the route table using the `utho_route` resource.
- Changing `vpc_id` or `dcslug` destroys and recreates the route table
  and all associated routes.
