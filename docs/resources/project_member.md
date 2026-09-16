---
page_title: "Project Member - Utho"
subcategory: "Account / Projects"
description: |-
  Add and manage members in a Utho Project.
---

# utho_project_member

Adds a member to a Utho Project with a specific role. Members are IAM sub-users who can access resources within the project based on their role.

## Example Usage

### Add a member

```hcl
resource "utho_project_member" "devops" {
  project_id = utho_project.main.id
  user_id    = 785755
  role_id    = 2
}
```

### Add multiple members with different roles

```hcl
resource "utho_project_member" "owner" {
  project_id = utho_project.main.id
  user_id    = 785755
  role_id    = 1
}

resource "utho_project_member" "developer" {
  project_id = utho_project.main.id
  user_id    = 785756
  role_id    = 2
}

resource "utho_project_member" "viewer" {
  project_id = utho_project.main.id
  user_id    = 785757
  role_id    = 3
}
```

## Argument Reference

| Argument     | Type   | Required | Description |
|--------------|--------|----------|-------------|
| `project_id` | String | Yes      | Project ID. Changing this forces a new resource. |
| `user_id`    | Number | Yes      | IAM user ID to add. Get this from `utho_iam_user.id`. Changing this forces a new resource. |
| `role_id`    | Number | Yes      | Role ID for the member. See [Roles](#roles). Updatable. |

## Roles

| Role ID | Name    | Description |
|---------|---------|-------------|
| `1`     | Owner   | Full access — can manage members and all resources. |
| `2`     | Member  | Can manage resources but not project settings or members. |
| `3`     | Viewer  | Read-only access to project resources. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{project_id}:{user_id}`. |

## Notes

- Changing `role_id` removes the member and re-adds them with the new role.
- Destroying this resource removes the member from the project.
- The user must exist as an IAM sub-user before being added as a member.