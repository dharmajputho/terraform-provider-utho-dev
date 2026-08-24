---
page_title: "Provider: Utho"
description: |-
  Use the Utho provider to manage Utho Cloud infrastructure.
---

# Utho Provider

The Utho provider allows you to manage infrastructure on [Utho Cloud](https://utho.com).

## Example Usage

```hcl
terraform {
  required_providers {
    utho = {
      source  = "dharmajputho/utho-dev"
      version = "~> 0.1"
    }
  }
}

provider "utho" {
  api_key = "your-api-key"
}
```

## Authentication

The provider requires an API key from your Utho account.

You can set it in the provider block or via environment variable:

```bash
export UTHO_API_KEY="your-api-key"
```

## Argument Reference

- `api_key` (Required) - Utho API key.