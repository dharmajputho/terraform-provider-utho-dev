---
page_title: "VPC Peering - Utho"
subcategory: "Networking / VPC"
description: |-
  Create and manage VPC peering connections between two Utho VPCs.
---

# utho_vpc_peering

Creates and manages a VPC peering connection between two Utho VPCs.
A peering connection allows instances in either VPC to communicate
with each other using private IP addresses.

## Example Usage

### Basic peering connection

```hcl
resource "utho_vpc_peering" "dev_prod" {
  name            = "dev-to-prod"
  dcslug          = "inmumbaizone2"
  requester_vpc_id = utho_vpc.dev.id
  accepter_vpc_id  = utho_vpc.prod.id
}

output "peering_id" {
  value = utho_vpc_peering.dev_prod.id
}

output "peering_status" {
  value = utho_vpc_peering.dev_prod.status
}
```

### Peer two existing VPCs by ID

```hcl
resource "utho_vpc_peering" "connection" {
  name             = "cross-team-peering"
  dcslug           = "inmumbaizone2"
  requester_vpc_id = "1534ffe8-964c-4ad7-85d0-b61eadb58e08"
  accepter_vpc_id  = "c066167d-5dbd-4aa2-8af2-07240f05a3b8"
}
```

### Full example with two VPCs and peering

```hcl
resource "utho_vpc" "app" {
  name    = "app-vpc"
  network = "10.1.0.0"
  size    = "20"
  dcslug  = "inmumbaizone2"
  planid  = "1008"
}

resource "utho_vpc" "data" {
  name    = "data-vpc"
  network = "10.2.0.0"
  size    = "20"
  dcslug  = "inmumbaizone2"
  planid  = "1008"
}

resource "utho_vpc_peering" "app_data" {
  name             = "app-to-data"
  dcslug           = "inmumbaizone2"
  requester_vpc_id = utho_vpc.app.id
  accepter_vpc_id  = utho_vpc.data.id
}
```

## Argument Reference

| Argument           | Type   | Required | Description |
|--------------------|--------|----------|-------------|
| `name`             | String | Yes      | Peering connection name. Changing this forces a new resource. |
| `dcslug`           | String | Yes      | Data center slug. Both VPCs must be in the same datacenter. Changing this forces a new resource. |
| `requester_vpc_id` | String | Yes      | ID of the VPC initiating the peering request. Changing this forces a new resource. |
| `accepter_vpc_id`  | String | Yes      | ID of the VPC accepting the peering request. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique peering connection ID (UUID). |
| `status`  | String | Peering connection status (e.g. `Active`). |

## Notes

- Both VPCs must be in the same datacenter.
- VPC CIDR ranges must not overlap.
- All fields require destroy and recreate on change.
- Deleting this resource removes the peering connection — instances
  in both VPCs will lose private connectivity.
