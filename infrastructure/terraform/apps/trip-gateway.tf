module "k8s_app_trip_gateway" {
  source = "../modules/k8s-app"

  depends_on = [kubernetes_namespace.ns]

  aws_creds_secret = kubernetes_secret.aws_secret.metadata.0.name
  context          = module.this.context
  name             = "gateway"
  label_order      = ["stage", "name"]
  image            = "ghcr.io/justtrackio/code-talks-2023-trip-gateway:1.0.0"
  ports = {
    traffic = 8080
  }
}
