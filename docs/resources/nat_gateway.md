---
page_title: "NAT Gateway - Utho"
subcategory: "Networking / VPC"
description: |-
  Create and manage NAT Gateways for Utho VPC subnets.
---

# utho_nat_gateway

Creates and manages a NAT Gateway for a Utho VPC subnet. A NAT Gateway
allows instances in a private subnet to access the internet without
being directly reachable from the internet.

When this resource is destroyed, Terraform automatically detaches the
NAT Gateway from the subnet before deleting it.

## Example Usage

### NAT Gateway with an Elastic IP

```hcl
resource "utho_elastic_ip" "nat_ip" {
  dcslug       = "inmumbaizone2"
  billingcycle = "monthly"
}

resource "utho_nat_gateway" "main" {
  name      = "main-nat-gateway"
  subnet_id = utho_subnet.public.id
  public_ip = utho_elastic_ip.nat_ip.ip
  dcslug    = "inmumbaizone2"
}

output "nat_gateway_id" {
  value = utho_nat_gateway.main.id
}
```

### Full private subnet with NAT access

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

# Private instances route outbound traffic through the NAT
resource "utho_cloud" "worker" {
  count    = 2
  hostname = "worker-${count.index + 1}.mhc"
  dcslug   = "inmumbaizone2"
  planid   = "10360"
  billingcycle = "hourly"
  auth     = "option1"
  root_password = "StrongP@ssw0rd!"
  image    = "ubuntu-22.04-x86_64"
  vpc      = utho_subnet.private.id
}
```

## Argument Reference

| Argument    | Type   | Required | Description |
|-------------|--------|----------|-------------|
| `name`      | String | Yes      | NAT Gateway name. Changing this forces a new resource. |
| `subnet_id` | String | Yes      | Subnet ID to attach the NAT Gateway to. Changing this forces a new resource. |
| `public_ip` | String | Yes      | Public IP address for outbound traffic. Use an `utho_elastic_ip`. Changing this forces a new resource. |
| `dcslug`    | String | Yes      | Data center slug. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique NAT Gateway ID (UUID). |

## Notes

- The NAT Gateway is automatically detached from the subnet before deletion.
- Use an `utho_elastic_ip` resource for the `public_ip` to ensure the IP
  is managed by Terraform.
- All fields require destroy and recreate on change.
