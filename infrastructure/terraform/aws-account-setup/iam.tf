module "iam_grafana_read_only_policy" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-read-only-policy"
  version = "v5.29"

  name = "grafana-cloudwatch-read-only"
  path = "/"

  allowed_services = ["cloudwatch", "logs"]
}

module "iam_user_grafana" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-user"
  version = "v5.29"

  name                          = "grafana"
  force_destroy                 = true
  password_reset_required       = false
  create_iam_user_login_profile = false
}

module "iam_group_grafana" {
  source  = "terraform-aws-modules/iam/aws//modules/iam-group-with-policies"
  version = "v5.29"

  name                     = "grafana"
  enable_mfa_enforcement   = false
  group_users              = [module.iam_user_grafana.iam_user_name]
  custom_group_policy_arns = [module.iam_grafana_read_only_policy.arn]
}
