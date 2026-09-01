---
page_title: "Route - Utho"
subcategory: "Networking / VPC"
description: |-
  Create and manage individual routes inside a Utho route table.
---

# utho_route

Creates and manages an individual route inside a Utho route table.
Each route defines where traffic for a specific destination CIDR
should be directed.

## Example Usage

### Internet gateway route

```hcl
resource "utho_route" "internet" {
  route_table_id         = utho_route_table.main.id
  destination_cidr_block = "0.0.0.0/0"
  route_type             = "igw"
  target                 = "igw"
}
```

### Local VPC route

```hcl
resource "utho_route" "local" {
  route_table_id         = utho_route_table.main.id
  destination_cidr_block = "10.0.0.0/16"
  route_type             = "local"
  target                 = "local"
}
```

### Multiple routes in one route table

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

resource "utho_route" "vpc_local" {
  route_table_id         = utho_route_table.main.id
  destination_cidr_block = "10.0.0.0/16"
  route_type             = "local"
  target                 = "local"
}
```

## Argument Reference

| Argument                 | Type   | Required | Description |
|--------------------------|--------|----------|-------------|
| `route_table_id`         | String | Yes      | Route table ID to add this route to. Changing this forces a new resource. |
| `destination_cidr_block` | String | Yes      | Destination CIDR block (e.g. `0.0.0.0/0`, `10.0.0.0/16`). |
| `route_type`             | String | Yes      | Route type (e.g. `igw`, `local`). |
| `target`                 | String | Yes      | Route target (e.g. `igw`, `local`). |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique route ID (UUID). |

## Notes

- Changing `route_table_id` destroys and recreates the route.
- Changing `destination_cidr_block`, `route_type`, or `target`
  updates the route in place.
- Multiple routes in the same route table must have unique
  `destination_cidr_block` values.
