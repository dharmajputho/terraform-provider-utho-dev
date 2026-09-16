---
page_title: "IPSec VPN Connection - Utho"
subcategory: "Networking / VPN"
description: |-
  Create and manage connections inside a Utho IPSec VPN tunnel.
---

# utho_ipsec_connection

Creates and manages a connection inside a Utho IPSec VPN tunnel. A connection defines the cryptographic parameters, peer IPs, and routing for traffic between your Utho VPC and a remote network.

Multiple connections can be added to the same tunnel for connecting multiple remote sites.

## Example Usage

### Standard IKEv2 connection

```hcl
resource "utho_ipsec_connection" "office" {
  ipsec_id          = utho_ipsec.main.id
  name              = "office-hq"
  remote_ip         = "203.0.113.10"
  remote_local_ip   = "192.168.1.0/24"
  local_ip          = "10.0.2.0/24"
  psk               = "my-secret-psk"
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

### High security connection with AES256-GCM

```hcl
resource "utho_ipsec_connection" "secure" {
  ipsec_id          = utho_ipsec.main.id
  name              = "secure-branch"
  remote_ip         = "203.0.113.20"
  remote_local_ip   = "172.16.0.0/24"
  local_ip          = "10.0.2.0/24"
  psk               = "ultra-secure-psk"
  phase1_encryption = "AES256-GCM-16"
  phase2_encryption = "AES256-GCM-16"
  phase1_integrity  = "SHA2-512"
  phase2_integrity  = "SHA2-512"
  phase1_dh_group   = "21"
  phase2_dh_group   = "21"
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

### Multiple remote networks

```hcl
resource "utho_ipsec_connection" "multi_net" {
  ipsec_id          = utho_ipsec.main.id
  name              = "multi-site"
  remote_ip         = "203.0.113.10"
  remote_local_ip   = "192.168.1.0/24,10.12.10.0/24"
  local_ip          = "10.0.2.0/24,10.0.3.0/24"
  psk               = "shared-secret"
  phase1_encryption = "AES128"
  phase2_encryption = "AES128"
  phase1_integrity  = "SHA1"
  phase2_integrity  = "SHA1"
  phase1_dh_group   = "14"
  phase2_dh_group   = "14"
  ike_version       = "ikev2"
  phase1_lifetime   = "28800"
  phase2_lifetime   = "3600"
  rekey_margin      = "270"
  rekey_fuzz        = "100"
  replay_window     = "1024"
  dpd_timeout       = "30"
  dpd_action        = "clear"
  startup_action    = "add"
}
```

## Argument Reference

| Argument            | Type   | Required | Description |
|---------------------|--------|----------|-------------|
| `ipsec_id`          | String | Yes      | IPSec tunnel ID. Changing this forces a new resource. |
| `name`              | String | Yes      | Connection name. Updatable. |
| `remote_ip`         | String | Yes      | Remote peer public IP address. Updatable. |
| `remote_local_ip`   | String | Yes      | Remote network CIDR(s). Comma-separated for multiple (e.g. `192.168.1.0/24,10.12.10.0/24`). Updatable. |
| `local_ip`          | String | Yes      | Local Utho network CIDR(s). Comma-separated for multiple. Updatable. |
| `psk`               | String | Yes      | Pre-shared key for authentication. **Sensitive.** Updatable. |
| `phase1_encryption` | String | Yes      | Phase 1 encryption. See [Encryption Algorithms](#encryption-algorithms). Updatable. |
| `phase2_encryption` | String | Yes      | Phase 2 encryption. Updatable. |
| `phase1_integrity`  | String | Yes      | Phase 1 integrity algorithm. See [Integrity Algorithms](#integrity-algorithms). Updatable. |
| `phase2_integrity`  | String | Yes      | Phase 2 integrity algorithm. Updatable. |
| `phase1_dh_group`   | String | Yes      | Phase 1 Diffie-Hellman group(s). See [DH Groups](#dh-groups). Updatable. |
| `phase2_dh_group`   | String | Yes      | Phase 2 DH group(s). Updatable. |
| `ike_version`       | String | Yes      | IKE version: `ikev1`, `ikev2`, or `ikev2,ikev1` for both. Updatable. |
| `phase1_lifetime`   | String | Yes      | Phase 1 SA lifetime in seconds. Default: `28800`. Updatable. |
| `phase2_lifetime`   | String | Yes      | Phase 2 SA lifetime in seconds. Default: `3600`. Updatable. |
| `rekey_margin`      | String | Yes      | Time before expiry to start rekeying in seconds. Default: `270`. Updatable. |
| `rekey_fuzz`        | String | Yes      | Rekey fuzz percentage. Default: `100`. Updatable. |
| `replay_window`     | String | Yes      | Anti-replay window size. Default: `1024`. Updatable. |
| `dpd_timeout`       | String | Yes      | Dead Peer Detection timeout in seconds. Default: `30`. Updatable. |
| `dpd_action`        | String | Yes      | DPD action: `none`, `restart`, or `clear`. Updatable. |
| `startup_action`    | String | Yes      | Startup action: `start` or `add`. Updatable. |

## Attribute Reference

| Attribute | Type   | Description |
|-----------|--------|-------------|
| `id`      | String | Unique connection ID. |
| `status`  | String | Connection status. |

## Encryption Algorithms

| Value           | Description |
|-----------------|-------------|
| `AES128`        | AES 128-bit |
| `AES256`        | AES 256-bit |
| `AES128-GCM-16` | AES 128-bit GCM |
| `AES256-GCM-16` | AES 256-bit GCM |

## Integrity Algorithms

| Value      | Description |
|------------|-------------|
| `SHA1`     | SHA-1 |
| `SHA2-256` | SHA-2 256-bit |
| `SHA2-384` | SHA-2 384-bit |
| `SHA2-512` | SHA-2 512-bit |

## DH Groups

Common values: `2`, `5`, `14`, `15`, `16`, `17`, `18`, `19`, `20`, `21`, `22`, `23`, `24`

## DPD Actions

| Value     | Description |
|-----------|-------------|
| `none`    | Do nothing when peer is detected dead. |
| `restart` | Restart the connection when peer is detected dead. |
| `clear`   | Clear the connection when peer is detected dead. |

## Startup Actions

| Value   | Description |
|---------|-------------|
| `start` | Initiate the connection immediately on startup. |
| `add`   | Add the connection but wait for the remote to initiate. |