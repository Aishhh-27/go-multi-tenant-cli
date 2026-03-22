# Go Multi-Tenant Infrastructure CLI

## Overview

This project provides a command-line tool written in Go to provision and manage isolated environments for multiple tenants. It automates the end-to-end workflow of creating infrastructure, deploying applications, and setting up corresponding GitLab projects.

The tool is designed to simulate real-world platform engineering workflows where multiple tenants require consistent, repeatable, and isolated environments.

---

## Features

* Provision Kubernetes namespaces per tenant using Terraform
* Deploy applications using Helm charts
* Create GitLab projects programmatically via API
* Execute provisioning concurrently for multiple tenants
* Generate per-tenant reports with environment details
* YAML-based configuration for batch tenant creation

---

## Architecture

The CLI orchestrates multiple components:

* **Go (Cobra CLI)** – Entry point and orchestration
* **Terraform** – Infrastructure provisioning
* **Kubernetes** – Namespace isolation
* **Helm** – Application deployment
* **GitLab API** – Project creation per tenant

Workflow:

1. Read tenant definitions from a YAML file
2. For each tenant (in parallel):

   * Create a GitLab project
   * Provision infrastructure using Terraform workspace
   * Create Kubernetes namespace
   * Deploy Helm chart
   * Generate a report file

---

## Project Structure

```
go-multi-tenant-cli/
├── cmd/
│   └── tenant.go
├── internal/
│   ├── gitlab/
│   ├── terraform/
│   ├── k8s/
│   └── utils/
├── terraform/
│   └── tenant.tf
├── charts/
│   └── gitlab/
├── main.go
├── tenants.yaml
```

---

## Prerequisites

Ensure the following tools are installed:

* Go (>= 1.20)
* Terraform
* kubectl
* Helm
* Access to a Kubernetes cluster
* GitLab account with Personal Access Token

---

## Setup

### 1. Clone the repository

```
git clone https://github.com/<your-username>/go-multi-tenant-cli.git
cd go-multi-tenant-cli
```

---

### 2. Configure GitLab token

Generate a Personal Access Token from GitLab with `api` scope.

Add it to your environment:

```
export GITLAB_TOKEN=your_token_here
```

---

### 3. Prepare tenant configuration

Create a `tenants.yaml` file:

```
tenants:
  - name: team1
  - name: team2
  - name: team3
```

---

### 4. Build the CLI

```
go mod tidy
go build -o tenant-cli
```

---

## Usage

Run the CLI with a tenant configuration file:

```
./tenant-cli create -f tenants.yaml
```

---

## Output

For each tenant, the tool performs:

* GitLab project creation
* Kubernetes namespace provisioning
* Helm deployment
* Report generation

Example report:

```
Tenant: team1
Namespace: team1
GitLab Project: https://gitlab.com/<username>/team1
```

---

## Concurrency Model

The tool uses Go routines and a WaitGroup to provision multiple tenants in parallel. Each tenant is handled independently, with failure isolation using panic recovery.

---

## Error Handling

* Terraform state locking handled with workspace isolation
* Helm failures are logged without stopping execution
* GitLab API errors are surfaced in CLI output
* Panic recovery ensures one tenant failure does not affect others

---

## Limitations

* No retry logic for transient failures
* Assumes local Kubernetes configuration
* Helm chart path is static
* GitLab project namespace defaults to user account

---

## Future Improvements

* Add retry and backoff mechanisms
* Support remote Terraform backends
* Integrate structured logging
* Add metrics and observability
* Support GitLab group/project hierarchy
* Containerize CLI for portability

---

## Why this project

This project demonstrates practical experience in:

* Infrastructure as Code (Terraform)
* Kubernetes-based multi-tenancy
* API-driven automation (GitLab)
* Concurrent systems in Go
* Platform engineering workflows

---

## License

MIT
