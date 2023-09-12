module "label_trip_consumer" {
  source = "justtrackio/label/null"

  context = module.this.context
  name    = "consumer"
}

module "gosoline_monitoring_trip_consumer" {
  source  = "justtrackio/ecs-gosoline-monitoring/aws"
  version = "2.1.3"

  depends_on = [grafana_folder.jt]

  context = module.label_trip_consumer.context

  containers                        = ["app"]
  elasticsearch_host                = local.k8s_elasticssearch_url
  elasticsearch_data_stream_enabled = true
  label_orders = {
    elasticsearch = ["namespace", "stage", "name"]
  }

  grafana_dashboard_enabled = true
}
