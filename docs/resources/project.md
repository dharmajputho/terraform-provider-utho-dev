---
page_title: "Project - Utho"
subcategory: "Account / Projects"
description: |-
  Create and manage Utho Projects to organize resources and team members.
---

# utho_project

Creates and manages a Utho Project. Projects help you organize resources and control team access by environment (development, staging, production, testing).

## Example Usage

### Basic project

```hcl
resource "utho_project" "dev" {
  name        = "my-app-dev"
  environment = "development"
  description = "Development environment for my app"
}

output "project_id" {
  value = utho_project.dev.id
}
```

### Multiple environment projects

```hcl
resource "utho_project" "dev" {
  name        = "my-app-dev"
  environment = "development"
}

resource "utho_project" "staging" {
  name        = "my-app-staging"
  environment = "staging"
}

resource "utho_project" "prod" {
  name        = "my-app-prod"
  environment = "production"
}
```

### Project with members

```hcl
resource "utho_project" "main" {
  name        = "production"
  environment = "production"
  description = "Main production project"
}

resource "utho_project_member" "devops" {
  project_id = utho_project.main.id
  user_id    = 785755
  role_id    = 2
}
```

## Argument Reference

| Argument      | Type   | Required | Description |
|---------------|--------|----------|-------------|
| `name`        | String | Yes      | Project name. Updatable. |
| `environment` | String | Yes      | Environment type: `development`, `staging`, `production`, or `testing`. Updatable. |
| `description` | String | No       | Project description. Updatable. |

## Attribute Reference

| Attribute       | Type   | Description |
|-----------------|--------|-------------|
| `id`            | String | Unique project ID. |
| `status`        | String | Project status (e.g. `active`). |
| `is_default`    | Bool   | Whether this is the default project. |
| `member_count`  | Number | Number of members in the project. |
| `resource_count`| Number | Number of resources in the project. |
| `created_at`    | String | Creation timestamp. |