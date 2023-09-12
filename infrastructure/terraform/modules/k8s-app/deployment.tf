locals {
  deployment_label_selector = {
    "app.kubernetes.io/name" = module.this.id
  }
}

resource "kubernetes_deployment" "app" {
  metadata {
    name      = module.this.id
    namespace = module.this.namespace
    labels    = local.deployment_label_selector
  }

  spec {
    replicas = var.replicas

    selector {
      match_labels = local.deployment_label_selector
    }

    template {
      metadata {
        labels = local.deployment_label_selector
      }

      spec {
        container {
          name              = "app"
          image             = var.image
          image_pull_policy = "Always"

          dynamic "volume_mount" {
            for_each = var.volume_container_path != null && var.volume_host_path != null ? [true] : []
            content {
              mount_path = var.volume_container_path
              name       = "host-path-volume"
            }
          }

          resources {
            requests = {
              cpu    = "10m"
              memory = "40Mi"
            }
            limits = {
              cpu    = "100m"
              memory = "50Mi"
            }
          }

          env_from {
            secret_ref {
              name = var.aws_creds_secret
            }
          }
          env {
            name  = "API_HEALTH_PORT"
            value = "8090"
          }
          env {
            name  = "ENV"
            value = module.this.environment
          }
          env {
            name  = "APPCTX_METADATA_SERVER_PORT"
            value = "8070"
          }
          env {
            name  = "CLOUD_AWS_DEFAULTS_ENDPOINT"
            value = ""
          }
          env {
            name  = "DX_AUTO_CREATE"
            value = "false"
          }
          env {
            name  = "DX_USE_RANDOM_PORTS"
            value = "false"
          }
          env {
            name  = "LOG_HANDLERS_MAIN_FORMATTER"
            value = "json"
          }
          env {
            name  = "METRIC_ENABLED"
            value = "true"
          }
          env {
            name  = "METRIC_WRITER"
            value = "cw"
          }
          dynamic "env" {
            for_each = var.env_vars

            content {
              name  = env.key
              value = env.value
            }
          }
          port {
            container_port = 8070
            name           = "metadata"
          }
          port {
            container_port = 8090
            name           = "health"
          }
          dynamic "port" {
            for_each = var.ports
            content {
              container_port = port.value
              name           = port.key
            }
          }
        }

        dynamic "volume" {
          for_each = var.volume_container_path != null && var.volume_host_path != null ? [true] : []

          content {
            name = "host-path-volume"
            host_path {
              path = var.volume_host_path
            }
          }
        }
      }
    }
  }
}
