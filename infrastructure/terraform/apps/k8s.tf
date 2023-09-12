resource "kubernetes_manifest" "traefik_pod_monitor" {
  manifest = {
    apiVersion = "monitoring.coreos.com/v1"
    kind       = "PodMonitor"

    metadata = {
      name      = "traefik"
      namespace = "kube-system"
    }

    spec = {
      podMetricsEndpoints = [
        {
          port = "metrics"
          path = "/metrics"
        }
      ]
      selector = {
        matchLabels = {
          "app.kubernetes.io/name" = "traefik"
        }
      }
    }
  }
}

resource "kubernetes_namespace" "ns" {
  metadata {
    name = module.this.namespace
  }
}

resource "kubernetes_secret" "aws_secret" {
  metadata {
    name      = "aws-creds"
    namespace = kubernetes_namespace.ns.metadata.0.name
  }
  data = {
    CLOUD_AWS_CREDENTIALS_ACCESS_KEY_ID     = var.aws_access_key_id,
    CLOUD_AWS_CREDENTIALS_SECRET_ACCESS_KEY = var.aws_secret_access_key,
    CLOUD_AWS_CREDENTIALS_SESSION_TOKEN     = var.aws_session_token,
  }
}
