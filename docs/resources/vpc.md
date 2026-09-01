---
page_title: "VPC - Utho"
subcategory: "Networking / VPC"
description: |-
  Create and manage Utho VPC networks.
---

# utho_vpc

Creates and manages a Utho Virtual Private Cloud (VPC) network. A VPC is an
isolated private network where you can deploy cloud instances, subnets,
NAT gateways, and route tables.

## Example Usage

### Basic VPC

```hcl
resource "utho_vpc" "main" {
  name    = "production-vpc"
  network = "10.0.0.0"
  size    = "20"
  dcslug  = "inmumbaizone2"
  planid  = "1008"
}

output "vpc_id" {
  value = utho_vpc.main.id
}
```

### VPC with subnets

```hcl
resource "utho_vpc" "main" {
  name    = "production-vpc"
  network = "10.0.0.0"
  size    = "16"
  dcslug  = "inmumbaizone2"
  planid  = "1008"
}

resource "utho_subnet" "public" {
  name            = "public-subnet"
  vpc_id          = utho_vpc.main.id
  network         = "10.0.1.0"
  size            = 24
  type            = "public"
  assign_publicip = 1
}

resource "utho_subnet" "private" {
  name            = "private-subnet"
  vpc_id          = utho_vpc.main.id
  network         = "10.0.2.0"
  size            = 24
  type            = "private"
  assign_publicip = 0
}
```

### Full VPC stack

```hcl
resource "utho_vpc" "main" {
  name    = "production-vpc"
  network = "10.0.0.0"
  size    = "16"
  dcslug  = "inmumbaizone2"
  planid  = "1008"
}

resource "utho_subnet" "public" {
  name            = "public-subnet"
  vpc_id          = utho_vpc.main.id
  network         = "10.0.1.0"
  size            = 24
  type            = "public"
  assign_publicip = 1
}

resource "utho_elastic_ip" "nat_ip" {
  dcslug       = "inmumbaizone2"
  billingcycle = "monthly"
}

resource "utho_nat_gateway" "main" {
  name      = "main-nat"
  subnet_id = utho_subnet.public.id
  public_ip = utho_elastic_ip.nat_ip.ip
  dcslug    = "inmumbaizone2"
}

resource "utho_route_table" "main" {
  name    = "main-routes"
  vpc_id  = utho_vpc.main.id
  dcslug  = "inmumbaizone2"
}

resource "utho_route" "internet" {
  route_table_id         = utho_route_table.main.id
  destination_cidr_block = "0.0.0.0/0"
  route_type             = "igw"
  target                 = "igw"
}

resource "utho_cloud" "app" {
  count    = 3
  hostname = "app-${count.index + 1}.mhc"
  dcslug   = "inmumbaizone2"
  planid   = "10360"
  billingcycle = "hourly"
  auth     = "option1"
  root_password = "StrongP@ssw0rd!"
  image    = "ubuntu-22.04-x86_64"
  vpc      = utho_subnet.public.id
}
```

## Argument Reference

| Argument  | Type   | Required | Description |
|-----------|--------|----------|-------------|
| `name`    | String | Yes      | VPC name. |
| `network` | String | Yes      | Base network address (e.g. `10.0.0.0`). Changing this forces a new resource. |
| `size`    | String | Yes      | CIDR prefix length (e.g. `20` for /20). Changing this forces a new resource. |
| `dcslug`  | String | Yes      | Data center slug (e.g. `inmumbaizone2`). Changing this forces a new resource. |
| `planid`  | String | Yes      | VPC plan ID. Changing this forces a new resource. |

## Attribute Reference

| Attribute   | Type   | Description |
|-------------|--------|-------------|
| `id`        | String | Unique VPC ID (UUID). |
| `is_default`| String | Whether this is the default VPC (`0` or `1`). |
| `total`     | Number | Total number of IP addresses in the VPC. |
| `available` | Number | Number of available IP addresses. |

## Notes

- VPCs cannot be updated in place — destroy and recreate to change network, size, or datacenter.
- Delete all subnets, NAT gateways, and route tables inside a VPC before deleting the VPC itself.
