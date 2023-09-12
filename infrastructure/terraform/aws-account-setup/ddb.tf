module "ddb_trips" {
  source  = "justtrackio/dynamodb-table/aws"
  version = "1.0.4"

  context = module.this.context
  name    = "trips"

  enable_autoscaler = false
  hash_key          = "uuid"

  # set to 1 (and recreate table) to show retry in logs. resource needs to be recreated, due to lifecycle ignore on read/write capacity
  autoscale_min_write_capacity = 5
  autoscale_min_read_capacity  = 5
}

module "ddb_vendors" {
  source  = "justtrackio/dynamodb-table/aws"
  version = "1.0.4"

  context = module.this.context
  name    = "vendors"

  enable_autoscaler = false
  hash_key          = "Id"
  hash_key_type     = "N"

  # set to 1 (and recreate table) to show retry in logs. resource needs to be recreated, due to lifecycle ignore on read/write capacity
  autoscale_min_write_capacity = 5
  autoscale_min_read_capacity  = 5
}