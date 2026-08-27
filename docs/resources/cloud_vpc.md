---
page_title: "utho_cloud_vpc Resource - utho"
subcategory: "Compute / Cloud Instances"
description: |-
  Attach or detach a VPC subnet from a Utho Cloud instance.
---

# utho_cloud_vpc

Attaches a VPC subnet to a Utho Cloud instance, assigning a private IP
address from the subnet's CIDR range. Enables private communication
between instances without traversing the public internet.

Destroying this resource detaches the subnet and releases the private IP.

~> **Important:** After attaching or detaching a VPC subnet, power-cycle
the instance using `utho_cloud_power` with `action = "powercycle"` so the
guest OS picks up the new or removed private network interface.

## Example Usage

### Attach a VPC subnet

```hcl
resource "utho_cloud_vpc" "private" {
  cloud_id  = utho_cloud.app.id
  subnet_id = "b0825dae-fd86-4d67-8238-c438774da894"
}

output "private_ip" {
  value = utho_cloud_vpc.private.private_ip
}
```

### Attach VPC and power-cycle to apply the NIC

```hcl
resource "utho_cloud_vpc" "private" {
  cloud_id  = utho_cloud.app.id
  subnet_id = "b0825dae-fd86-4d67-8238-c438774da894"
}

resource "utho_cloud_power" "apply_vpc" {
  cloud_id   = utho_cloud.app.id
  action     = "powercycle"
  depends_on = [utho_cloud_vpc.private]
}
```

### Connect multiple instances to the same private network

```hcl
resource "utho_cloud" "nodes" {
  count    = 3
  hostname = "node-${count.index + 1}.mhc"
  # ...
}

resource "utho_cloud_vpc" "node_network" {
  count     = 3
  cloud_id  = utho_cloud.nodes[count.index].id
  subnet_id = "b0825dae-fd86-4d67-8238-c438774da894"
}

output "private_ips" {
  value = utho_cloud_vpc.node_network[*].private_ip
}
```

## Argument Reference

| Argument    | Type   | Required | Description |
|-------------|--------|----------|-------------|
| `cloud_id`  | String | Yes      | The ID of the cloud instance. Changing this forces a new resource. |
| `subnet_id` | String | Yes      | The UUID of the VPC subnet (e.g. `b0825dae-fd86-4d67-8238-c438774da894`). Changing this forces a new resource. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Identifier in the format `{cloud_id}:{subnet_id}`. |
| `private_ip` | String | Private IP address assigned from the subnet's CIDR range. |

## Notes

- The private IP is assigned automatically from the subnet's address pool.
- Changing `cloud_id` or `subnet_id` destroys the current attachment and creates a new one.
- Multiple subnets can be attached to the same instance.
- Always power-cycle the instance after attaching or detaching to apply NIC changes.
