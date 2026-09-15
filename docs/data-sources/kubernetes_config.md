---
page_title: "Kubernetes Config Data Source - Utho"
subcategory: "Compute / Kubernetes"
description: |-
  Fetch the kubeconfig for a Utho Kubernetes cluster.
---

# utho_kubernetes_config

Fetches the kubeconfig YAML for a Utho Kubernetes cluster. Use the raw config to authenticate `kubectl` or configure Kubernetes providers.

## Example Usage

### Fetch kubeconfig

```hcl
data "utho_kubernetes_config" "main" {
  cluster_id = utho_kubernetes.main.id
}

output "kubeconfig" {
  value     = data.utho_kubernetes_config.main.raw_config
  sensitive = true
}
```

### Save kubeconfig to a file

```hcl
data "utho_kubernetes_config" "main" {
  cluster_id = utho_kubernetes.main.id
}

resource "local_file" "kubeconfig" {
  content  = data.utho_kubernetes_config.main.raw_config
  filename = "${path.module}/kubeconfig.yaml"
}
```

### Use with the Kubernetes provider

```hcl
data "utho_kubernetes_config" "main" {
  cluster_id = utho_kubernetes.main.id
}

provider "kubernetes" {
  config_path = "${path.module}/kubeconfig.yaml"
}
```

## Argument Reference

| Argument     | Type   | Required | Description |
|--------------|--------|----------|-------------|
| `cluster_id` | String | Yes      | Kubernetes cluster ID. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Same as `cluster_id`. |
| `raw_config` | String | Full kubeconfig YAML content. **Sensitive.** |

~> **Security note:** The kubeconfig contains cluster credentials. Mark any outputs using this value as `sensitive = true` and never commit it to version control.
