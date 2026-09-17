---
page_title: "Container Registry Immutable Rule - Utho"
subcategory: "Container Registry"
description: |-
  Create immutable tag rules for a Utho Container Registry.
---

# utho_container_registry_immutable_rule

Creates an immutable tag rule for a Utho Container Registry. Immutable tag rules prevent specific image tags from being overwritten or deleted, protecting production releases from accidental modification.

## Example Usage

### Protect all version tags

```hcl
resource "utho_container_registry_immutable_rule" "versions" {
  project_name = utho_container_registry.main.project_name
  tag_pattern  = "v*"
  repo_pattern = "**"
}
```

### Protect production tags in specific repo

```hcl
resource "utho_container_registry_immutable_rule" "prod" {
  project_name = utho_container_registry.main.project_name
  tag_pattern  = "prod-*"
  repo_pattern = "my-app"
}
```

### Protect release candidates

```hcl
resource "utho_container_registry_immutable_rule" "rc" {
  project_name = utho_container_registry.main.project_name
  tag_pattern  = "*-rc*"
  repo_pattern = "**"
}
```

## Argument Reference

| Argument       | Type   | Required | Description |
|----------------|--------|----------|-------------|
| `project_name` | String | Yes      | Registry project name. Changing this forces a new resource. |
| `tag_pattern`  | String | Yes      | Tag pattern to protect. Supports wildcards (e.g. `v*`, `prod-*`). Changing this forces a new resource. |
| `repo_pattern` | String | Yes      | Repository pattern. Use `**` for all repos. Changing this forces a new resource. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Identifier in the format `{project_name}:{tag_pattern}`. |
| `rule_id` | Number | Rule ID assigned by the registry. Required for deletion. |

## Notes

- Immutable tags cannot be overwritten or deleted once pushed.
- Use `v*` to protect all semantic version tags.
- Use `**` as `repo_pattern` to apply the rule to all repositories.
- Changing any field destroys and recreates the rule.