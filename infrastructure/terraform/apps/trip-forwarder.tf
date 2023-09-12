module "k8s_app_trip_forwarder" {
  source = "../modules/k8s-app"

  depends_on = [kubernetes_namespace.ns]

  aws_creds_secret = kubernetes_secret.aws_secret.metadata.0.name
  context          = module.this.context
  name             = "forwarder"
  label_order      = ["stage", "name"]
  image            = "ghcr.io/justtrackio/code-talks-2023-trip-forwarder:1.0.0"

  volume_host_path      = "/var/lib/rancher/k3s/storage/forwarder"
  volume_container_path = "/input"
  env_vars = {
    FORWARDER_PATH        = "/input/trips.json"
    FORWARDER_GATEWAY_URL = "http://terminal-trips-gateway.codetalks:8080/trip"
  }
}
