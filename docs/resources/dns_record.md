---
page_title: "DNS Record - Utho"
subcategory: "Networking / DNS"
description: |-
  Create and manage DNS records in a Utho DNS zone.
---

# utho_dns_record

Creates and manages a DNS record inside a Utho DNS zone. Supports A, AAAA, CNAME, MX, TXT, SRV, and NS record types.

## Example Usage

### A record pointing to a load balancer

```hcl
resource "utho_dns_record" "root" {
  domain   = utho_dns_zone.main.domain
  type     = "A"
  hostname = "@"
  value    = utho_loadbalancer.main.ip
  ttl      = "3600"
}
```

### CNAME record

```hcl
resource "utho_dns_record" "www" {
  domain   = utho_dns_zone.main.domain
  type     = "CNAME"
  hostname = "www"
  value    = "example.com"
  ttl      = "3600"
}
```

### MX record

```hcl
resource "utho_dns_record" "mail" {
  domain   = utho_dns_zone.main.domain
  type     = "MX"
  hostname = "@"
  value    = "mail.example.com"
  ttl      = "3600"
}
```

### TXT record for domain verification

```hcl
resource "utho_dns_record" "verify" {
  domain   = utho_dns_zone.main.domain
  type     = "TXT"
  hostname = "@"
  value    = "v=spf1 include:_spf.google.com ~all"
  ttl      = "3600"
}
```

### AAAA record (IPv6)

```hcl
resource "utho_dns_record" "ipv6" {
  domain   = utho_dns_zone.main.domain
  type     = "AAAA"
  hostname = "@"
  value    = "2001:db8::1"
  ttl      = "3600"
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `domain`   | String | Yes      | Domain name the record belongs to. Changing this forces a new resource. |
| `type`     | String | Yes      | Record type: `A`, `AAAA`, `CNAME`, `MX`, `TXT`, `SRV`, `NS`. |
| `hostname` | String | Yes      | Hostname for the record. Use `@` for the root domain. |
| `value`    | String | Yes      | Record value — IP address, target hostname, or text content. |
| `ttl`      | String | Yes      | Time to live in seconds (e.g. `3600`). |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique DNS record ID assigned by Utho. |

## Notes

- Changing `domain` destroys and recreates the record.
- Changing `type`, `hostname`, `value`, or `ttl` updates the record in place.
- Use `@` as the hostname to create a record for the root domain.
