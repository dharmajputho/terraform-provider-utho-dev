---
page_title: "Load Balancer - Utho"
subcategory: "Networking / Load Balancing"
description: |-
  Create and manage Utho Load Balancers.
---

# utho_loadbalancer

Creates and manages a Utho Load Balancer. A load balancer distributes incoming traffic across multiple backend servers, improving availability and scalability. Choose `network` type for TCP/UDP traffic or `application` type for HTTP/HTTPS with advanced routing.

After creating a load balancer, add frontends (listeners) with `utho_loadbalancer_frontend` and backend servers with `utho_loadbalancer_backend`.

## Load Balancer Types

| Type | Best For |
|------|----------|
| `network` | TCP/UDP traffic, non-HTTP protocols, lower latency |
| `application` | HTTP/HTTPS with path-based routing, SSL termination, sticky sessions |

## Example Usage

### Simple HTTP load balancer

A network LB distributing traffic to 3 web servers on port 80.

```hcl
resource "utho_loadbalancer" "main" {
  name            = "web-lb"
  type            = "network"
  dcslug          = "inmumbaizone2"
  enable_publicip = "true"
}

resource "utho_loadbalancer_frontend" "http" {
  loadbalancer_id = utho_loadbalancer.main.id
  name            = "http"
  algorithm       = "roundrobin"
  proto           = "http"
  port            = "80"
  cookie          = "0"
  redirecthttps   = "0"
  certificate_id  = "0"
}

resource "utho_loadbalancer_backend" "web" {
  count           = 3
  loadbalancer_id = utho_loadbalancer.main.id
  frontend_id     = utho_loadbalancer_frontend.http.id
  backend_port    = "80"
  weight          = "1"
  type            = "cloud"
  cloudid         = utho_cloud.web[count.index].id
}

output "lb_ip"  { value = utho_loadbalancer.main.ip }
output "lb_dns" { value = utho_loadbalancer.main.dns }
```

### HTTPS load balancer with SSL termination

SSL is terminated at the load balancer — backends receive plain HTTP. Requires an SSL certificate uploaded via `utho_ssl_certificate`.

```hcl
resource "utho_ssl_certificate" "main" {
  name            = "myapp-cert"
  type            = "Custom"
  certificate_key = file("cert.pem")
  private_key     = file("key.pem")
}

resource "utho_loadbalancer" "main" {
  name            = "web-lb"
  type            = "application"
  dcslug          = "inmumbaizone2"
  enable_publicip = "true"
}

# Redirect HTTP to HTTPS
resource "utho_loadbalancer_frontend" "http" {
  loadbalancer_id = utho_loadbalancer.main.id
  name            = "http"
  algorithm       = "roundrobin"
  proto           = "http"
  port            = "80"
  cookie          = "0"
  redirecthttps   = "1"    # redirect to HTTPS
  certificate_id  = "0"
}

# HTTPS with SSL cert
resource "utho_loadbalancer_frontend" "https" {
  loadbalancer_id = utho_loadbalancer.main.id
  name            = "https"
  algorithm       = "roundrobin"
  proto           = "https"
  port            = "443"
  cookie          = "0"
  redirecthttps   = "0"
  certificate_id  = utho_ssl_certificate.main.id
}

resource "utho_loadbalancer_backend" "web" {
  count           = 3
  loadbalancer_id = utho_loadbalancer.main.id
  frontend_id     = utho_loadbalancer_frontend.https.id
  backend_port    = "80"
  weight          = "1"
  type            = "cloud"
  cloudid         = utho_cloud.web[count.index].id
}
```

### Private load balancer inside a VPC

Internal load balancer — only accessible from within the VPC. Used for service-to-service communication.

```hcl
resource "utho_loadbalancer" "internal" {
  name            = "internal-lb"
  type            = "network"
  dcslug          = "inmumbaizone2"
  vpc             = utho_subnet.private.id
  enable_publicip = "false"
  firewall        = utho_firewall.internal.id
}
```

## Argument Reference

### Required

| Argument          | Type   | Description |
|-------------------|--------|-------------|
| `name`            | String | Load balancer name. |
| `type`            | String | `network` or `application`. Changing this forces a new resource. |
| `dcslug`          | String | Data center. Changing this forces a new resource. |
| `enable_publicip` | String | `"true"` to assign a public IP, `"false"` for internal only. |

### Optional

| Argument   | Type   | Description |
|------------|--------|-------------|
| `vpc`      | String | VPC subnet ID. Required for private load balancers. Changing this forces a new resource. |
| `firewall` | String | Security group ID to attach at creation. Changing this forces a new resource. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique load balancer ID. |
| `ip`         | String | Public IP address. Point your DNS A record here. |
| `dns`        | String | DNS hostname provided by Utho. |
| `status`     | String | Load balancer status (`Active`, `Pending`). |
| `created_at` | String | Creation timestamp. |

## Notes

- Load balancer provisioning takes 1–3 minutes. Wait for `status = "Active"` before adding backends.
- The `dns` hostname is stable even if the IP changes — point your DNS CNAME record to it for production.
- For zero-downtime updates, add new backends before removing old ones.
- Advanced timeout and performance settings are available via `utho_loadbalancer_settings`.
