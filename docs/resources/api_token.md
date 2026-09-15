---
page_title: "API Token - Utho"
subcategory: "Account / IAM"
description: |-
  Create and manage Utho API tokens.
---

# utho_api_token

Creates and manages a Utho API token. API tokens are used to authenticate with the Utho API for CI/CD pipelines, automation scripts, and third-party integrations.

~> **Important:** The token value is only available at creation time. Save it immediately — it cannot be retrieved later.

## Example Usage

### Read/write token

```hcl
resource "utho_api_token" "ci" {
  name  = "github-actions"
  write = "on"
}

output "api_token" {
  value     = utho_api_token.ci.token
  sensitive = true
}
```

### Read-only token

```hcl
resource "utho_api_token" "readonly" {
  name  = "monitoring-token"
  write = "off"
}
```

### Use token in CI/CD

```hcl
resource "utho_api_token" "deploy" {
  name  = "deploy-pipeline"
  write = "on"
}

# Pass to your CI system as a secret
output "deploy_token" {
  value     = utho_api_token.deploy.token
  sensitive = true
}
```

### Multiple tokens for different teams

```hcl
resource "utho_api_token" "infra_team" {
  name  = "infra-team"
  write = "on"
}

resource "utho_api_token" "ops_team" {
  name  = "ops-team"
  write = "off"
}
```

## Argument Reference

| Argument | Type   | Required | Description |
|----------|--------|----------|-------------|
| `name`   | String | Yes      | Token name. Changing this forces a new resource. |
| `write`  | String | Yes      | Write access level: `on` (read and write) or `off` (read only). Changing this forces a new resource. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique token ID assigned by Utho. |
| `token`      | String | Generated API token value. **Sensitive — shown only at creation time.** |
| `created_at` | String | Timestamp when the token was created. |

## Notes

- Tokens cannot be updated in place — destroy and recreate to change name or write access.
- The `token` value begins with `live_` and is only visible immediately after creation.
- Store the token value in a secrets manager (e.g. GitHub Actions secrets, HashiCorp Vault) — never in version control.
- Use `write = "off"` for read-only access (monitoring, auditing) and `write = "on"` for deployment pipelines.