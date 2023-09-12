module "label_trip_gateway" {
  source = "justtrackio/label/null"

  context = module.this.context
  name    = "gateway"
}

module "gosoline_monitoring_trip_gateway" {
  source  = "justtrackio/ecs-gosoline-monitoring/aws"
  version = "2.1.3"

  depends_on = [grafana_folder.jt]

  context = module.label_trip_gateway.context

  containers                        = ["app"]
  elasticsearch_host                = local.k8s_elasticssearch_url
  elasticsearch_data_stream_enabled = true
  label_orders = {
    elasticsearch = ["namespace", "stage", "name"]
  }

  grafana_dashboard_enabled = true
}
