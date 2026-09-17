---
page_title: "Container Registry - Utho"
subcategory: "Container Registry"
description: |-
  Create and manage Utho Container Registries.
---

# utho_container_registry

Creates and manages a Utho Container Registry. A container registry stores and distributes Docker images for your applications. Manage robot accounts using `utho_container_registry_robot`, webhooks using `utho_container_registry_webhook`, and immutable tag rules using `utho_container_registry_immutable_rule`.

## Example Usage

### Public registry

```hcl
resource "utho_container_registry" "main" {
  project_name  = "my-app-registry"
  dcslug        = "innoida"
  planid        = "10276"
  billingcycle  = "monthly"
  public        = "true"
}

output "push_command" {
  value = "docker push registry.utho.io/${utho_container_registry.main.project_name}/<repo>:<tag>"
}
```

### Private registry with robot account

```hcl
resource "utho_container_registry" "main" {
  project_name  = "my-app-registry"
  dcslug        = "innoida"
  planid        = "10276"
  billingcycle  = "monthly"
  public        = "false"
}

resource "utho_container_registry_robot" "ci" {
  project_name = utho_container_registry.main.project_name
  name         = "ci-deploy"
  description  = "CI/CD deployment robot"
  duration     = 365

  access = [
    { resource = "repository", action = "pull" },
    { resource = "repository", action = "push" },
  ]
}

output "robot_secret" {
  value     = utho_container_registry_robot.ci.secret
  sensitive = true
}
```

## Argument Reference

| Argument       | Type   | Required | Description |
|----------------|--------|----------|-------------|
| `project_name` | String | Yes      | Registry name. Must be unique. Changing this forces a new resource. |
| `dcslug`       | String | Yes      | Data center slug. Changing this forces a new resource. |
| `planid`       | String | Yes      | Registry plan ID. Changing this forces a new resource. |
| `billingcycle` | String | Yes      | Billing cycle: `monthly`. Changing this forces a new resource. |
| `public`       | String | Yes      | Registry visibility: `true` (public) or `false` (private). Updatable. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Registry project name (used as identifier). |
| `created_at` | String | Creation timestamp. |

## Notes

- The push endpoint format is: `registry.utho.io/{project_name}/<repo>:<tag>`
- Changing `public` updates visibility in place without destroying the registry.
- All other fields require destroy and recreate.