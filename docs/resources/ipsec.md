---
page_title: "IPSec VPN Tunnel - Utho"
subcategory: "Networking / VPN"
description: |-
  Create and manage Utho IPSec site-to-site VPN tunnels.
---

# utho_ipsec

Creates and manages an IPSec site-to-site VPN tunnel. Use this when you need to connect a remote network (your office, data center, or another cloud) to your Utho VPC over an encrypted tunnel.

After the tunnel is created, define connections with `utho_ipsec_connection`. Each connection specifies the cryptographic parameters and IP ranges for one remote site.

## Example Usage

### Connect your office to Utho VPC

```hcl
# The tunnel gateway lives in your VPC
resource "utho_ipsec" "office" {
  name         = "office-vpn"
  dcslug       = "inmumbaizone2"
  vpc          = utho_subnet.private.id
  billingcycle = "monthly"
  cpumodel     = "amd"
}

# One connection per remote site
resource "utho_ipsec_connection" "hq" {
  ipsec_id          = utho_ipsec.office.id
  name              = "office-hq"
  remote_ip         = "203.0.113.1"       # your office public IP
  remote_local_ip   = "192.168.1.0/24"   # your office internal network
  local_ip          = "10.0.2.0/24"      # your Utho VPC subnet
  psk               = var.vpn_psk
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

output "tunnel_psk" {
  value     = utho_ipsec.office.psk
  sensitive = true
}
```

### Multi-site VPN hub

Connect multiple branch offices to the same Utho VPC.

```hcl
resource "utho_ipsec" "hub" {
  name         = "vpn-hub"
  dcslug       = "inmumbaizone2"
  vpc          = utho_subnet.private.id
  billingcycle = "monthly"
}

locals {
  branches = {
    delhi     = { ip = "203.0.113.10", network = "192.168.1.0/24" }
    mumbai    = { ip = "203.0.113.20", network = "192.168.2.0/24" }
    bangalore = { ip = "203.0.113.30", network = "192.168.3.0/24" }
  }
}

resource "utho_ipsec_connection" "branch" {
  for_each = local.branches

  ipsec_id          = utho_ipsec.hub.id
  name              = each.key
  remote_ip         = each.value.ip
  remote_local_ip   = each.value.network
  local_ip          = "10.0.2.0/24"
  psk               = var.vpn_psk
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
| `dcslug`       | String | Yes      | Data center. Changing this forces a new resource. |
| `vpc`          | String | Yes      | VPC subnet ID where the tunnel gateway will live. Changing this forces a new resource. |
| `billingcycle` | String | Yes      | `monthly`, `3month`, `6month`, or `12month`. Changing this forces a new resource. |
| `cpumodel`     | String | No       | `amd` or `intel`. Changing this forces a new resource. |

## Attribute Reference

| Attribute    | Type   | Description |
|--------------|--------|-------------|
| `id`         | String | Unique tunnel ID. Use this in `utho_ipsec_connection`. |
| `psk`        | String | Auto-generated pre-shared key for the tunnel. **Sensitive.** Configure this on your remote VPN device. |
| `status`     | String | Tunnel status. |
| `created_at` | String | Creation timestamp. |

## Notes

- The `psk` attribute is the master tunnel PSK. Individual connections can have their own PSKs defined in `utho_ipsec_connection`.
- Configure your remote VPN device (router, firewall, or cloud VPN gateway) with the Utho tunnel IP and PSK.
- Use Dead Peer Detection (`dpd_action = "restart"`) to automatically recover from connection drops.
