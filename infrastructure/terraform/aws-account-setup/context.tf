module "this" {
  source = "justtrackio/label/null"

  namespace   = "codetalks"
  environment = "prod"
  tenant      = "demo"
  stage       = "terminal-trips"

  label_order = ["namespace", "environment", "tenant", "stage", "name", "attributes"]

  aws_account_id = var.aws_account_id
  aws_region     = var.aws_region
}
