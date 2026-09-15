---
page_title: "Security Group - Utho"
subcategory: "Networking / Security"
description: |-
  Create and manage Utho Security Groups.
---

# utho_firewall

Creates and manages a Utho Security Group. A Security Group is a set of
inbound and outbound rules that control traffic to and from cloud instances.

Manage rules using `utho_firewall_rule` and attach instances using
`utho_firewall_server`.

## Example Usage

### Basic Security Group

```hcl
resource "utho_firewall" "web" {
  name = "web-security-group"
}

output "security_group_id" {
  value = utho_firewall.web.id
}
```

### Security Group with rules and server attachment

```hcl
resource "utho_firewall" "web" {
  name = "web-security-group"
}

resource "utho_firewall_rule" "allow_http" {
  firewall_id  = utho_firewall.web.id
  type         = "incoming"
  service      = "HTTP"
  protocol     = "tcp"
  port         = "80"
  port_range   = "80"
  addresses    = "0.0.0.0/0"
  source_range = "0.0.0.0/0"
}

resource "utho_firewall_rule" "allow_https" {
  firewall_id  = utho_firewall.web.id
  type         = "incoming"
  service      = "HTTPS"
  protocol     = "tcp"
  port         = "443"
  port_range   = "443"
  addresses    = "0.0.0.0/0"
  source_range = "0.0.0.0/0"
}

resource "utho_firewall_rule" "allow_ssh" {
  firewall_id  = utho_firewall.web.id
  type         = "incoming"
  service      = "SSH"
  protocol     = "tcp"
  port         = "22"
  port_range   = "22"
  addresses    = "10.0.0.0/8"
  source_range = "10.0.0.0/8"
}

resource "utho_firewall_rule" "allow_all_outbound" {
  firewall_id  = utho_firewall.web.id
  type         = "outgoing"
  service      = "ALL TCP"
  protocol     = "tcp"
  port         = "ALL"
  port_range   = "ALL"
  addresses    = "0.0.0.0/0"
  source_range = "0.0.0.0/0"
}

resource "utho_firewall_server" "attach" {
  firewall_id = utho_firewall.web.id
  cloud_id    = utho_cloud.app.id
}
```

## Argument Reference

| Argument | Type   | Required | Description |
|----------|--------|----------|-------------|
| `name`   | String | Yes      | Name of the Security Group. |

## Attribute Reference

| Attribute      | Type   | Description |
|----------------|--------|-------------|
| `id`           | String | Unique Security Group ID. |
| `created_at`   | String | Timestamp when the Security Group was created. |
| `rule_count`   | String | Number of rules currently in this Security Group. |
| `servers_count`| String | Number of servers currently attached. |