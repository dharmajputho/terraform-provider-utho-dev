---
page_title: "DNS Zone - Utho"
subcategory: "Networking / DNS"
description: |-
  Create and manage public DNS zones on Utho.
---

# utho_dns_zone

Creates and manages a public DNS zone on Utho. A DNS zone holds all DNS records for a domain. Once created, add records with `utho_dns_record`, then point your domain's nameservers to Utho to activate resolution.

## Example Usage

### Create a zone and add records

```hcl
resource "utho_dns_zone" "main" {
  domain = "myapp.com"
}

# Point the root domain to a load balancer
resource "utho_dns_record" "root" {
  domain   = utho_dns_zone.main.domain
  type     = "A"
  hostname = "@"
  value    = utho_loadbalancer.main.ip
  ttl      = "300"
}

# www subdomain
resource "utho_dns_record" "www" {
  domain   = utho_dns_zone.main.domain
  type     = "CNAME"
  hostname = "www"
  value    = "myapp.com"
  ttl      = "300"
}

# API subdomain points to a different server
resource "utho_dns_record" "api" {
  domain   = utho_dns_zone.main.domain
  type     = "A"
  hostname = "api"
  value    = utho_cloud.api.ip
  ttl      = "60"
}
```

### After creating the zone — update your registrar

Point your domain's nameservers to Utho at your domain registrar (GoDaddy, Namecheap, etc.). The nameservers to use will be shown in the Utho dashboard after the zone is created.

## Argument Reference

| Argument | Type   | Required | Description |
|----------|--------|----------|-------------|
| `domain` | String | Yes      | Domain name (e.g. `myapp.com`). Do not include a trailing dot. Changing this forces a new resource. |

## Attribute Reference

| Attribute      | Type   | Description |
|----------------|--------|-------------|
| `id`           | String | The domain name (used as identifier). |
| `nspoint`      | String | `YES` if nameservers are pointed to Utho, `NO` otherwise. |
| `record_count` | String | Number of DNS records in this zone. |
| `created_at`   | String | Creation timestamp. |

## Notes

- Creating a zone in Utho does not automatically make the domain resolve — you must update your registrar's nameservers.
- DNS propagation after changing nameservers can take up to 48 hours, though typically takes 1–4 hours.
- Use a low TTL (e.g. `60`) when you're about to change a record's value — this reduces the time old IPs are cached.
- Use a high TTL (e.g. `3600`) for stable records — this reduces DNS query load and improves response time.
