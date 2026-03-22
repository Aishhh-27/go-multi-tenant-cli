provider "kubernetes" {
  config_path = "~/.kube/config"
}

variable "tenant_name" {}

resource "kubernetes_namespace" "tenant_ns" {
  metadata {
    name = var.tenant_name
  }

  lifecycle {
    ignore_changes = all
  }
}
