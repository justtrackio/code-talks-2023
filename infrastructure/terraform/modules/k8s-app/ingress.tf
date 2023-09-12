resource "kubernetes_ingress_v1" "metadata" {
  metadata {
    name      = "${module.this.id}-metadata"
    namespace = module.this.namespace
  }
  spec {
    rule {
      host = "${module.this.id}-metadata.127.0.0.1.nip.io"
      http {
        path {
          path = "/"
          backend {
            service {
              name = kubernetes_service.service.metadata.0.name
              port {
                number = kubernetes_service.service.spec.0.port.0.port
              }
            }
          }
        }
      }
    }
  }
}

resource "kubernetes_ingress_v1" "this" {
  for_each = var.ports
  metadata {
    name      = "${module.this.id}-${each.key}"
    namespace = module.this.namespace
  }
  spec {
    rule {
      host = "${module.this.id}-${each.key}.127.0.0.1.nip.io"
      http {
        path {
          path = "/"
          backend {
            service {
              name = kubernetes_service.service.metadata.0.name
              port {
                name = each.key
              }
            }
          }
        }
      }
    }
  }
}

