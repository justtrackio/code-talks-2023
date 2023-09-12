resource "kubernetes_service" "service" {
  metadata {
    name      = module.this.id
    namespace = module.this.namespace
  }
  spec {
    selector = local.deployment_label_selector
    port {
      port         = 8070
      app_protocol = "http"
      name         = "metadata"
      protocol     = "TCP"
    }
    port {
      port         = 8090
      app_protocol = "http"
      name         = "health"
      protocol     = "TCP"
    }
    dynamic "port" {
      for_each = var.ports
      content {
        port         = port.value
        name         = port.key
        app_protocol = "http"
        protocol     = "TCP"
      }
    }
  }
}
