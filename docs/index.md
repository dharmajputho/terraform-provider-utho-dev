---
page_title: "Provider: Utho"
description: |-
  Use the Utho provider to manage cloud infrastructure on Utho Cloud — VMs, Kubernetes, databases, networking, storage, and more — all from Terraform.
---

# Utho Provider

The **Utho Terraform provider** lets you define, provision, and manage your entire cloud infrastructure on [Utho Cloud](https://utho.com) using infrastructure-as-code. Instead of clicking through a console, you write `.tf` files that describe what you want — Terraform creates it, tracks it, and tears it down when you're done.

## Why Use Terraform with Utho?

- **Reproducible infrastructure** — check your `.tf` files into git and spin up identical environments for dev, staging, and production
- **No manual steps** — one `terraform apply` creates everything in the correct dependency order
- **Safe teardown** — `terraform destroy` removes everything cleanly without leaving orphaned resources
- **Drift detection** — `terraform plan` shows you when your real infrastructure diverges from your code

## Quick Start

### 1. Configure the provider

```hcl
terraform {
  required_providers {
    utho = {
      source  = "dharmajputho/utho-dev"
      version = "~> 0.1"
    }
  }
}

provider "utho" {
  api_key = var.utho_api_key
}

variable "utho_api_key" {
  description = "Utho API key"
  type        = string
  sensitive   = true
}
```

### 2. Get your API key

Log in to the [Utho Dashboard](https://console.utho.com), go to **Settings → API Tokens**, and generate a new token.

### 3. Set the key securely

```bash
# Option A — environment variable (recommended)
export TF_VAR_utho_api_key="live_your_key_here"

# Option B — tfvars file (add to .gitignore)
echo 'utho_api_key = "live_your_key_here"' > terraform.tfvars
```

### 4. Deploy

```bash
terraform init    # download the provider
terraform plan    # preview what will be created
terraform apply   # create resources
```

## Authentication

| Method | Configuration |
|--------|--------------|
| Provider block | `api_key = "live_..."` |
| Environment variable | `export UTHO_API_KEY="live_..."` |
| Terraform variable | `export TF_VAR_utho_api_key="live_..."` |

~> **Security:** Never commit API keys to version control. Use environment variables or a secrets manager like HashiCorp Vault.

## Provider Arguments

| Argument  | Type   | Required | Description |
|-----------|--------|----------|-------------|
| `api_key` | String | Yes      | Utho API key. Also reads from `UTHO_API_KEY` environment variable. |

## Data Centers

| Slug | Location |
|------|----------|
| `innoida` | Delhi (Noida), India |
| `inmumbaizone2` | Mumbai, India |
| `inbangalore` | Bangalore, India |

## Resources

### Compute / Cloud Instances

| Resource | Description |
|----------|-------------|
| [utho_cloud](resources/cloud) | Create and manage cloud instances (VMs) |
| [utho_cloud_power](resources/cloud_power) | Control power state — start, stop, reboot |
| [utho_cloud_resize](resources/cloud_resize) | Resize CPU and RAM |
| [utho_cloud_snapshot](resources/cloud_snapshot) | Create and restore snapshots |
| [utho_cloud_iso](resources/cloud_iso) | Mount an ISO for custom OS installs |
| [utho_cloud_firewall](resources/cloud_firewall) | Attach or detach security groups |
| [utho_cloud_storage](resources/cloud_storage) | Add general-purpose storage disks |
| [utho_cloud_ebs](resources/cloud_ebs) | Attach and manage EBS block volumes |
| [utho_cloud_public_ip](resources/cloud_public_ip) | Assign or release additional public IPs |
| [utho_cloud_vpc](resources/cloud_vpc) | Attach or detach VPC subnets |
| [utho_ssh_key](resources/ssh_key) | Import and manage SSH public keys |

### Compute / Kubernetes

| Resource | Description |
|----------|-------------|
| [utho_kubernetes](resources/kubernetes) | Create and manage managed Kubernetes clusters |
| [utho_kubernetes_node_pool](resources/kubernetes_node_pool) | Add and scale node pools |

### Compute / Auto Scaling

| Resource | Description |
|----------|-------------|
| [utho_autoscaling](resources/autoscaling) | Create auto scaling groups |
| [utho_autoscaling_policy](resources/autoscaling_policy) | Define metric-based scaling policies |
| [utho_autoscaling_schedule](resources/autoscaling_schedule) | Schedule scaling at specific times |

### Networking / VPC

| Resource | Description |
|----------|-------------|
| [utho_vpc](resources/vpc) | Create isolated private networks |
| [utho_subnet](resources/subnet) | Create subnets inside a VPC |
| [utho_nat_gateway](resources/nat_gateway) | Enable outbound internet from private subnets |
| [utho_route_table](resources/route_table) | Create and manage route tables |
| [utho_route](resources/route) | Add routes to a route table |
| [utho_elastic_ip](resources/elastic_ip) | Allocate static public IP addresses |
| [utho_vpc_peering](resources/vpc_peering) | Connect two VPCs together |

### Networking / Security

| Resource | Description |
|----------|-------------|
| [utho_firewall](resources/firewall) | Create security groups |
| [utho_firewall_rule](resources/firewall_rule) | Add inbound and outbound rules |
| [utho_firewall_server](resources/firewall_server) | Attach instances to a security group |

### Networking / Load Balancing

| Resource | Description |
|----------|-------------|
| [utho_loadbalancer](resources/loadbalancer) | Create network or application load balancers |
| [utho_loadbalancer_frontend](resources/loadbalancer_frontend) | Define listener ports and protocols |
| [utho_loadbalancer_backend](resources/loadbalancer_backend) | Add backend servers |
| [utho_loadbalancer_acl](resources/loadbalancer_acl) | Add traffic routing rules |
| [utho_loadbalancer_settings](resources/loadbalancer_settings) | Configure timeouts and advanced settings |

### Networking / DNS

| Resource | Description |
|----------|-------------|
| [utho_dns_zone](resources/dns_zone) | Create public DNS zones |
| [utho_dns_record](resources/dns_record) | Add DNS records (A, CNAME, MX, TXT, etc.) |

### Networking / VPN

| Resource | Description |
|----------|-------------|
| [utho_ipsec](resources/ipsec) | Create IPSec site-to-site VPN tunnels |
| [utho_ipsec_connection](resources/ipsec_connection) | Configure tunnel connections |

### Security

| Resource | Description |
|----------|-------------|
| [utho_ssl_certificate](resources/ssl_certificate) | Upload and manage SSL/TLS certificates |

### Storage / Object Storage

| Resource | Description |
|----------|-------------|
| [utho_object_storage](resources/object_storage) | Create S3-compatible object storage buckets |
| [utho_object_storage_permission](resources/object_storage_permission) | Grant bucket access to keys |
| [utho_object_storage_key](resources/object_storage_key) | Create S3 access keys |

### Database

| Resource | Description |
|----------|-------------|
| [utho_database](resources/database) | Create managed PostgreSQL clusters |
| [utho_database_db](resources/database_db) | Create databases inside a cluster |
| [utho_database_user](resources/database_user) | Create database users |
| [utho_database_pool](resources/database_pool) | Create PgBouncer connection pools |

### Container Registry

| Resource | Description |
|----------|-------------|
| [utho_container_registry](resources/container_registry) | Create private container registries |
| [utho_container_registry_robot](resources/container_registry_robot) | Create CI/CD robot accounts |
| [utho_container_registry_webhook](resources/container_registry_webhook) | Configure event webhooks |
| [utho_container_registry_immutable_rule](resources/container_registry_immutable_rule) | Protect tags from overwrite |

### Monitoring

| Resource | Description |
|----------|-------------|
| [utho_alert_contact](resources/alert_contact) | Create notification contacts |
| [utho_alert](resources/alert) | Create metric-based alert rules |

### Account / IAM

| Resource | Description |
|----------|-------------|
| [utho_api_token](resources/api_token) | Create scoped API tokens |
| [utho_iam_user](resources/iam_user) | Invite sub-users with granular permissions |
| [utho_project](resources/project) | Organize resources into projects |
| [utho_project_member](resources/project_member) | Add team members to projects |

## Data Sources

| Data Source | Description |
|-------------|-------------|
| [utho_clouds](data-sources/clouds) | List all cloud instances in your account |
| [utho_kubernetes_config](data-sources/kubernetes_config) | Fetch kubeconfig for a Kubernetes cluster |

## Full Production Stack Example

This example provisions a complete production environment — VPC, security groups, 3 load-balanced web servers, and a managed PostgreSQL database — wired together with Terraform references so everything is created in the right order automatically.

```hcl
terraform {
  required_providers {
    utho = { source = "dharmajputho/utho-dev", version = "~> 0.1" }
  }
}

variable "utho_api_key" { sensitive = true }
variable "root_password" { sensitive = true }
provider "utho" { api_key = var.utho_api_key }

# SSH Key
resource "utho_ssh_key" "deploy" {
  name   = "deploy-key"
  sshkey = file("~/.ssh/id_ed25519.pub")
}

# VPC
resource "utho_vpc" "main" {
  name   = "production"
  network = "10.0.0.0"
  size   = "16"
  dcslug = "inmumbaizone2"
  planid = "1008"
}

resource "utho_subnet" "public" {
  name            = "public"
  vpc_id          = utho_vpc.main.id
  network         = "10.0.1.0"
  size            = 24
  type            = "public"
  assign_publicip = 1
}

# Security Group
resource "utho_firewall" "web" { name = "web-sg" }

resource "utho_firewall_rule" "http" {
  firewall_id  = utho_firewall.web.id
  type         = "incoming"
  service      = "HTTP"
  protocol     = "tcp"
  port         = "80"
  port_range   = "80"
  addresses    = "0.0.0.0/0"
  source_range = "0.0.0.0/0"
}

# Web Servers
resource "utho_cloud" "app" {
  count           = 3
  hostname        = "app-${count.index + 1}.mhc"
  dcslug          = "inmumbaizone2"
  planid          = "10360"
  billingcycle    = "hourly"
  auth            = "option2"
  sshkeys         = utho_ssh_key.deploy.id
  image           = "ubuntu-22.04-x86_64"
  enable_publicip = "true"
  vpc             = utho_subnet.public.id
}

resource "utho_firewall_server" "app" {
  count       = 3
  firewall_id = utho_firewall.web.id
  cloud_id    = utho_cloud.app[count.index].id
}

# Load Balancer
resource "utho_loadbalancer" "main" {
  name            = "main-lb"
  type            = "application"
  dcslug          = "inmumbaizone2"
  enable_publicip = "true"
  vpc             = utho_subnet.public.id
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

resource "utho_loadbalancer_backend" "app" {
  count           = 3
  loadbalancer_id = utho_loadbalancer.main.id
  frontend_id     = utho_loadbalancer_frontend.http.id
  backend_port    = "80"
  weight          = "1"
  type            = "cloud"
  cloudid         = utho_cloud.app[count.index].id
}

# Database
resource "utho_database" "pg" {
  cluster_name  = "production-db"
  dcslug        = "inmumbaizone2"
  engine        = "pg"
  version       = "17"
  size          = "10157"
  network_type  = "private"
  billing       = "monthly"
  pitr_enabled  = "1"
  replica_count = "1"
}

# DNS
resource "utho_dns_zone" "main" { domain = "myapp.com" }

resource "utho_dns_record" "root" {
  domain   = utho_dns_zone.main.domain
  type     = "A"
  hostname = "@"
  value    = utho_loadbalancer.main.ip
  ttl      = "300"
}

# Outputs
output "lb_ip"   { value = utho_loadbalancer.main.ip }
output "db_host" { value = "pg-${utho_database.pg.id}.db.onutho.com" }
output "app_ips" { value = utho_cloud.app[*].ip }
```

Run `terraform apply` — Terraform resolves all the dependencies automatically and creates resources in the right order.
