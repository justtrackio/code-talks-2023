output "grafana_access_key_id" {
  value = module.iam_user_grafana.iam_access_key_id
}

output "grafana_secret_access_key" {
  value     = module.iam_user_grafana.iam_access_key_secret
  sensitive = true
}
