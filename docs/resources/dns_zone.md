---
page_title: "DNS Zone - Utho"
subcategory: "Networking / DNS"
description: |-
  Create and manage Utho DNS zones.
---

# utho_dns_zone

Creates and manages a public DNS zone on Utho. Once created, add records using `utho_dns_record`. Point your domain's nameservers to Utho to activate DNS resolution.

## Example Usage

### Create a DNS zone

```hcl
resource "utho_dns_zone" "main" {
  domain = "example.com"
}
```

### Zone with records

```hcl
resource "utho_dns_zone" "main" {
  domain = "example.com"
}

resource "utho_dns_record" "root" {
  domain   = utho_dns_zone.main.domain
  type     = "A"
  hostname = "@"
  value    = utho_loadbalancer.main.ip
  ttl      = "3600"
}

resource "utho_dns_record" "www" {
  domain   = utho_dns_zone.main.domain
  type     = "CNAME"
  hostname = "www"
  value    = "example.com"
  ttl      = "3600"
}
```

## Argument Reference

| Argument | Type   | Required | Description |
|----------|--------|----------|-------------|
| `domain` | String | Yes      | Domain name (e.g. `example.com`). Changing this forces a new resource. |

## Attribute Reference

| Attribute      | Type   | Description |
|----------------|--------|-------------|
| `id`           | String | The domain name (used as identifier). |
| `nspoint`      | String | Whether nameservers are pointed to Utho (`YES` or `NO`). |
| `record_count` | String | Number of DNS records in this zone. |
| `created_at`   | String | Timestamp when the zone was created. |
