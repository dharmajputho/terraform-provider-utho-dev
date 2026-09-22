---
page_title: "SSL Certificates - Utho"
subcategory: "Networking / Load Balancing"
description: |-
  List all SSL certificates in your Utho account.
---

# utho_ssl_certificates

Fetches all SSL certificates in your account. Use certificate IDs when creating HTTPS frontends on load balancers.

## Example Usage

### List all certificates

```hcl
data "utho_ssl_certificates" "all" {}

output "certificates" {
  value = [
    for c in data.utho_ssl_certificates.all.certificates :
    { id = c.id, name = c.name, expires = c.expire_at, days_left = c.remaining_days }
  ]
}
```

### Find a certificate by name for HTTPS frontend

```hcl
data "utho_ssl_certificates" "all" {}

locals {
  cert = one([
    for c in data.utho_ssl_certificates.all.certificates :
    c if c.name == "my-ssl-cert" && c.state == "verified"
  ])
}

resource "utho_loadbalancer_frontend" "https" {
  loadbalancer_id = utho_loadbalancer.main.id
  name            = "https-frontend"
  algorithm       = "roundrobin"
  proto           = "https"
  port            = "443"
  cookie          = "0"
  redirecthttps   = "0"
  certificate_id  = local.cert.id   # ← from data source
}
```

### List only verified certificates

```hcl
data "utho_ssl_certificates" "all" {}

output "verified_certs" {
  value = [
    for c in data.utho_ssl_certificates.all.certificates :
    { id = c.id, name = c.name, days_left = c.remaining_days }
    if c.state == "verified"
  ]
}
```

### Check for expiring certificates

```hcl
data "utho_ssl_certificates" "all" {}

output "expiring_soon" {
  description = "Certificates expiring within 30 days."
  value = [
    for c in data.utho_ssl_certificates.all.certificates :
    { id = c.id, name = c.name, expires = c.expire_at, days_left = c.remaining_days }
    if c.remaining_days < 30
  ]
}
```

## Attribute Reference

### Top-level

| Attribute      | Type | Description |
|----------------|------|-------------|
| `certificates` | List | List of all SSL certificates in your account. |

### certificates

| Attribute            | Type   | Description |
|----------------------|--------|-------------|
| `id`                 | String | Certificate ID. Use this as `certificate_id` in `utho_loadbalancer_frontend`. |
| `name`               | String | Certificate name. |
| `type`               | String | Certificate type: `custom` (uploaded) or `lets_encrypt` (auto-issued). |
| `state`              | String | Certificate state: `verified`, `pending`, or `failed`. Only use `verified` certificates. |
| `status`             | String | Certificate status. |
| `primary_domain`     | String | Primary domain the certificate covers. |
| `issuer`             | String | Certificate issuer details. |
| `issued_at`          | String | Issue timestamp. |
| `expire_at`          | String | Expiry timestamp. Set up alerts before this date. |
| `remaining_days`     | Number | Days remaining before the certificate expires. |
| `key_algorithm`      | String | Key algorithm: `RSA` or `EC`. |
| `signature_algorithm`| String | Signature algorithm (e.g. `RSA-SHA256`). |
| `created_at`         | String | Upload/creation timestamp. |

## Notes

- Only `verified` certificates can be attached to load balancer frontends.
- To upload a new SSL certificate use `utho_ssl_certificate`.
- Monitor `remaining_days` to rotate certificates before expiry.
- Use `certificate_id = "0"` in `utho_loadbalancer_frontend` when no SSL is needed (HTTP only).