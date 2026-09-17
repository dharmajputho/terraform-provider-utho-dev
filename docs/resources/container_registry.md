---
page_title: "Container Registry - Utho"
subcategory: "Container Registry"
description: |-
  Create and manage Utho Container Registries for storing Docker images.
---

# utho_container_registry

Creates and manages a Utho Container Registry — a private Docker registry for storing and distributing container images. Integrated with Harbor, it supports vulnerability scanning, webhook notifications, immutable tags, and robot accounts for CI/CD automation.

## Example Usage

### Private registry for a project

```hcl
resource "utho_container_registry" "main" {
  project_name  = "myapp-registry"
  dcslug        = "innoida"
  planid        = "10276"
  billingcycle  = "monthly"
  public        = "false"
}

output "push_endpoint" {
  value = "registry.utho.io/${utho_container_registry.main.project_name}"
}
```

### Registry with a CI/CD robot account

Robot accounts are service accounts for automated pushes from CI/CD pipelines. The secret is only shown at creation — save it as a CI secret immediately.

```hcl
resource "utho_container_registry" "main" {
  project_name  = "myapp-registry"
  dcslug        = "innoida"
  planid        = "10276"
  billingcycle  = "monthly"
  public        = "false"
}

resource "utho_container_registry_robot" "ci" {
  project_name = utho_container_registry.main.project_name
  name         = "github-actions"
  description  = "GitHub Actions CI/CD robot"
  duration     = 365

  access = [
    { resource = "repository", action = "pull" },
    { resource = "repository", action = "push" },
  ]
}

# Store these as GitHub Actions secrets: REGISTRY_USERNAME, REGISTRY_PASSWORD
output "robot_username" {
  value = utho_container_registry_robot.ci.name
}
output "robot_password" {
  value     = utho_container_registry_robot.ci.secret
  sensitive = true
}
```

### Registry with push webhook

Get notified in Slack or your CI system whenever a new image is pushed.

```hcl
resource "utho_container_registry_webhook" "push_notify" {
  project_name   = utho_container_registry.main.project_name
  name           = "push-notify"
  address        = "https://hooks.slack.com/services/YOUR/WEBHOOK/URL"
  payload_format = "CloudEvents"
  skip_cert_verify = false

  event_types = [
    "PUSH_ARTIFACT",
    "DELETE_ARTIFACT",
  ]
}
```

### Protect release tags from overwrite

Immutable tag rules prevent production image tags from being accidentally overwritten.

```hcl
resource "utho_container_registry_immutable_rule" "releases" {
  project_name = utho_container_registry.main.project_name
  tag_pattern  = "v*"    # protect all tags starting with v
  repo_pattern = "**"    # apply to all repositories
}
```

## Using the Registry

After creating the registry and a robot account:

```bash
# Login
docker login registry.utho.io \
  -u "api$myapp-registry+github-actions" \
  -p "<robot_secret>"

# Tag and push
docker tag myapp:latest registry.utho.io/myapp-registry/myapp:v1.0.0
docker push registry.utho.io/myapp-registry/myapp:v1.0.0

# Pull
docker pull registry.utho.io/myapp-registry/myapp:v1.0.0
```

## Argument Reference

| Argument       | Type   | Required | Description |
|----------------|--------|----------|-------------|
| `project_name` | String | Yes      | Registry name. Must be unique. Changing this forces a new resource. |
| `dcslug`       | String | Yes      | Data center. Currently only `innoida` is supported. Changing this forces a new resource. |
| `planid`       | String | Yes      | Registry plan ID. Changing this forces a new resource. |
| `billingcycle` | String | Yes      | `monthly`. Changing this forces a new resource. |
| `public`       | String | Yes      | `"true"` for public pull access, `"false"` for private. Updatable. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Registry project name (used as identifier). |
| `created_at` | String | Creation timestamp. |

## Push URL Format

```
registry.utho.io/{project_name}/{repository}:{tag}
```

Example: `registry.utho.io/myapp-registry/backend:v1.2.3`
