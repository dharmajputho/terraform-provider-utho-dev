---
page_title: "IPSec VPN Tunnel - Utho"
subcategory: "Networking / VPN"
description: |-
  Create and manage Utho IPSec site-to-site VPN tunnels.
---

# utho_ipsec

Creates and manages a Utho IPSec site-to-site VPN tunnel. An IPSec tunnel connects your Utho VPC to an external network (office, data center, or another cloud) over an encrypted tunnel. Add connections using `utho_ipsec_connection`.

## Example Usage

### Basic IPSec tunnel

```hcl
resource "utho_ipsec" "main" {
  name         = "office-to-utho"
  dcslug       = "inmumbaizone2"
  vpc          = utho_subnet.private.id
  billingcycle = "monthly"
  cpumodel     = "amd"
}

output "tunnel_psk" {
  value     = utho_ipsec.main.psk
  sensitive = true
}
```

### Tunnel with connection

```hcl
resource "utho_ipsec" "main" {
  name         = "office-to-utho"
  dcslug       = "inmumbaizone2"
  vpc          = utho_subnet.private.id
  billingcycle = "monthly"
}

resource "utho_ipsec_connection" "office" {
  ipsec_id          = utho_ipsec.main.id
  name              = "office-connection"
  remote_ip         = "203.0.113.10"
  remote_local_ip   = "192.168.1.0/24"
  local_ip          = "10.0.2.0/24"
  psk               = "my-secret-key"
  phase1_encryption = "AES256"
  phase2_encryption = "AES256"
  phase1_integrity  = "SHA2-256"
  phase2_integrity  = "SHA2-256"
  phase1_dh_group   = "14"
  phase2_dh_group   = "14"
  ike_version       = "ikev2"
  phase1_lifetime   = "28800"
  phase2_lifetime   = "3600"
  rekey_margin      = "270"
  rekey_fuzz        = "100"
  replay_window     = "1024"
  dpd_timeout       = "30"
  dpd_action        = "restart"
  startup_action    = "start"
}
```

## Argument Reference

| Argument       | Type   | Required | Description |
|----------------|--------|----------|-------------|
| `name`         | String | Yes      | Tunnel name. Changing this forces a new resource. |
| `dcslug`       | String | Yes      | Data center slug. Changing this forces a new resource. |
| `vpc`          | String | Yes      | VPC subnet ID. Changing this forces a new resource. |
| `billingcycle` | String | Yes      | Billing cycle: `monthly`, `3month`, `6month`, `12month`. Changing this forces a new resource. |
| `cpumodel`     | String | No       | CPU model: `amd` or `intel`. Changing this forces a new resource. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique IPSec tunnel ID. |
| `psk`        | String | Auto-generated pre-shared key. **Sensitive.** |
| `status`     | String | Tunnel status. |
| `created_at` | String | Creation timestamp. |