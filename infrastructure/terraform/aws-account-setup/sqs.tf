module "label_sqs" {
  source = "justtrackio/label/null"

  context = module.this.context
  name    = "trips"
}

module "sqs_trips" {
  source  = "justtrackio/sqs/aws"
  version = "1.4.1"

  context = module.label_sqs.context
}
