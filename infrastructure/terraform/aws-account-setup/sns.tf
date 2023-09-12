module "label_sns" {
  source = "justtrackio/label/null"

  context     = module.this.context
  name        = "alarms"
  label_order = ["environment", "name"]
}

resource "aws_sns_topic" "alarms" {
  name = module.label_sns.id
}
