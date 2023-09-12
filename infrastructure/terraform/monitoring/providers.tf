locals {
  elasticssearch_url     = "http://elasticsearch-127.0.0.1.nip.io:8081"
  k8s_elasticssearch_url = "http://elasticsearch-es-http.elastic:9200"
  gosoline = {
    metadata = {
      domain    = "nip.io"
      use_https = false
      port      = 8081
    }
    name_patterns = {
      hostname                         = "{scheme}://{group}-{app}-metadata.127.0.0.1.{metadata_domain}:{port}"
      cloudwatch_namespace             = "{project}/{env}/{family}/{group}-{app}"
      ecs_cluster                      = "{env}"
      ecs_service                      = "{group}-{app}"
      grafana_cloudwatch_datasource    = "cloudwatch"
      grafana_elasticsearch_datasource = "elasticsearch-{project}-{group}-{app}"
      kubernetes_namespace             = "{project}"
      kubernetes_pod                   = "{group}-{app}"
      traefik_service_name             = "{project}-{group}-{app}-traffic@kubernetes"
    }
  }
}

provider "aws" {
  region  = "eu-central-1"
  profile = var.aws_profile
}

provider "gosoline" {
  metadata      = local.gosoline.metadata
  name_patterns = local.gosoline.name_patterns
  orchestrator  = "kubernetes"
}

provider "grafana" {
  url  = "http://grafana-127.0.0.1.nip.io:8081"
  auth = "admin:prom-operator"
}

provider "elasticsearch" {
  url = local.elasticssearch_url
}

provider "elasticstack" {
  elasticsearch {
    endpoints = [local.elasticssearch_url]
    insecure  = true
  }
}
