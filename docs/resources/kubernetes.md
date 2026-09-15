---
page_title: "Kubernetes Cluster - Utho"
subcategory: "Compute / Kubernetes"
description: |-
  Create and manage Utho Kubernetes clusters.
---

# utho_kubernetes

Creates and manages a Utho Managed Kubernetes cluster. Node pools are defined inline at creation time. Additional node pools can be added using `utho_kubernetes_node_pool`.

## Example Usage

### Basic cluster

```hcl
resource "utho_kubernetes" "main" {
  dcslug          = "inmumbaizone2"
  cluster_label   = "production-cluster"
  cluster_version = "1.30.0-utho"
  network_type    = "public"

  nodepools = [
    {
      label     = "worker-pool"
      size      = "10355"
      count     = "3"
      min_nodes = "1"
      max_nodes = "10"
      disk_size = "30"
      disk_type = "nvme"
    }
  ]
}

output "cluster_id" {
  value = utho_kubernetes.main.id
}

output "cluster_dns" {
  value = utho_kubernetes.main.dns
}
```

### Cluster inside a VPC

```hcl
resource "utho_kubernetes" "private" {
  dcslug          = "inmumbaizone2"
  cluster_label   = "private-cluster"
  cluster_version = "1.30.0-utho"
  network_type    = "publicprivate"
  vpc             = utho_subnet.private.id
  cpumodel        = "intel"

  nodepools = [
    {
      label     = "app-pool"
      size      = "10355"
      count     = "2"
      min_nodes = "2"
      max_nodes = "5"
    }
  ]
}
```

### Get kubeconfig after creation

```hcl
data "utho_kubernetes_config" "main" {
  cluster_id = utho_kubernetes.main.id
}

output "kubeconfig" {
  value     = data.utho_kubernetes_config.main.raw_config
  sensitive = true
}
```

## Argument Reference

| Argument           | Type   | Required | Description |
|--------------------|--------|----------|-------------|
| `dcslug`           | String | Yes      | Data center slug. Changing this forces a new resource. |
| `cluster_label`    | String | Yes      | Cluster name. Changing this forces a new resource. |
| `cluster_version`  | String | Yes      | Kubernetes version (e.g. `1.30.0-utho`). Changing this forces a new resource. |
| `network_type`     | String | Yes      | Network type: `public`, `private`, or `publicprivate`. Changing this forces a new resource. |
| `nodepools`        | List   | Yes      | Initial node pools. See [Node Pool Block](#node-pool-block). |
| `vpc`              | String | No       | VPC subnet ID. Changing this forces a new resource. |
| `cpumodel`         | String | No       | CPU model: `amd` or `intel`. Changing this forces a new resource. |

### Node Pool Block

| Argument    | Type   | Required | Description |
|-------------|--------|----------|-------------|
| `label`     | String | Yes      | Node pool label. |
| `size`      | String | Yes      | Plan ID for worker node size. |
| `count`     | String | Yes      | Number of worker nodes. |
| `min_nodes` | String | Yes      | Minimum nodes for autoscaling. |
| `max_nodes` | String | Yes      | Maximum nodes for autoscaling. |
| `disk_size` | String | No       | Additional EBS disk size in GB. |
| `disk_type` | String | No       | EBS disk type: `nvme` or `ssd`. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique cluster ID. |
| `status`     | String | Cluster status (e.g. `Active`, `Pending`). |
| `ip`         | String | Control plane IP address. |
| `dns`        | String | Control plane DNS endpoint. |
| `created_at` | String | Creation timestamp. |
