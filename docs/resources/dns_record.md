---
page_title: "DNS Record - Utho"
subcategory: "Networking / DNS"
description: |-
  Create and manage DNS records inside a Utho DNS zone.
---

# utho_dns_record

Creates and manages a DNS record inside a Utho DNS zone. Records tell the DNS system where to route traffic for a hostname.

## Example Usage

### A record — map hostname to IP

```hcl
# Root domain → load balancer
resource "utho_dns_record" "root" {
  domain   = "myapp.com"
  type     = "A"
  hostname = "@"           # @ means the root domain
  value    = utho_loadbalancer.main.ip
  ttl      = "300"
}

# Subdomain → specific server
resource "utho_dns_record" "api" {
  domain   = "myapp.com"
  type     = "A"
  hostname = "api"
  value    = utho_cloud.api.ip
  ttl      = "300"
}
```

### CNAME record — alias one name to another

```hcl
resource "utho_dns_record" "www" {
  domain   = "myapp.com"
  type     = "CNAME"
  hostname = "www"
  value    = "myapp.com"
  ttl      = "3600"
}
```

### MX record — email routing

```hcl
resource "utho_dns_record" "mail" {
  domain   = "myapp.com"
  type     = "MX"
  hostname = "@"
  value    = "mail.myapp.com"
  ttl      = "3600"
}
```

### TXT record — domain verification and SPF

```hcl
# SPF record
resource "utho_dns_record" "spf" {
  domain   = "myapp.com"
  type     = "TXT"
  hostname = "@"
  value    = "v=spf1 include:_spf.google.com ~all"
  ttl      = "3600"
}

# Domain verification for an external service
resource "utho_dns_record" "verify" {
  domain   = "myapp.com"
  type     = "TXT"
  hostname = "_verification"
  value    = "verify=abc123"
  ttl      = "3600"
}
```

### AAAA record — IPv6

```hcl
resource "utho_dns_record" "ipv6" {
  domain   = "myapp.com"
  type     = "AAAA"
  hostname = "@"
  value    = "2001:db8::1"
  ttl      = "3600"
}
```

## Argument Reference

| Argument   | Type   | Required | Description |
|------------|--------|----------|-------------|
| `domain`   | String | Yes      | Domain name the record belongs to (must have a zone). Changing this forces a new resource. |
| `type`     | String | Yes      | Record type: `A`, `AAAA`, `CNAME`, `MX`, `TXT`, `SRV`, `NS`. |
| `hostname` | String | Yes      | Subdomain or `@` for the root domain. |
| `value`    | String | Yes      | Record value — IP address, target hostname, or text. |
| `ttl`      | String | Yes      | Time-to-live in seconds. Lower values propagate changes faster; higher values reduce DNS load. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique DNS record ID assigned by Utho. |

## TTL Guidelines

| TTL | When to Use |
|-----|-------------|
| `60` | During migrations or when you expect to change the value soon |
| `300` | Default for most records |
| `3600` | Stable records that rarely change |
| `86400` | Very stable records (nameservers, email) |
