---
page_title: "Container Registry Robot - Utho"
subcategory: "Container Registry"
description: |-
  Create and manage robot accounts for a Utho Container Registry.
---

# utho_container_registry_robot

Creates and manages a robot account for a Utho Container Registry. Robot accounts are service accounts used for automated access (CI/CD pipelines, deployment scripts) with fine-grained pull/push permissions and an expiry date.

~> **Important:** The `secret` is only shown at creation time. Save it immediately.

## Example Usage

### CI/CD robot with pull and push access

```hcl
resource "utho_container_registry_robot" "ci" {
  project_name = utho_container_registry.main.project_name
  name         = "ci-deploy"
  description  = "GitHub Actions deployment robot"
  duration     = 365

  access = [
    { resource = "repository", action = "pull" },
    { resource = "repository", action = "push" },
  ]
}

output "robot_name" {
  value = utho_container_registry_robot.ci.name
}

output "robot_secret" {
  value     = utho_container_registry_robot.ci.secret
  sensitive = true
}
```

### Read-only robot for production pulls

```hcl
resource "utho_container_registry_robot" "prod_pull" {
  project_name = utho_container_registry.main.project_name
  name         = "prod-puller"
  description  = "Production read-only access"
  duration     = 90

  access = [
    { resource = "repository", action = "pull" },
  ]
}
```

## Argument Reference

| Argument       | Type   | Required | Description |
|----------------|--------|----------|-------------|
| `project_name` | String | Yes      | Registry project name. Changing this forces a new resource. |
| `name`         | String | Yes      | Robot account name. Changing this forces a new resource. |
| `duration`     | Number | Yes      | Token validity in days. Changing this forces a new resource. |
| `access`       | List   | Yes      | Access permissions. See [Access Block](#access-block). Changing this forces a new resource. |
| `description`  | String | No       | Robot account description. Changing this forces a new resource. |

### Access Block

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `resource` | String | Yes      | Resource type: `repository`. |
| `action`   | String | Yes      | Action: `pull` or `push`. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique robot account ID. |
| `secret`     | String | Robot account secret for authentication. **Sensitive — shown only at creation time.** |
| `expires_at` | Number | Expiry timestamp (Unix). |