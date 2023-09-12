# Setup

First ensure you copy the `.envrc.template` to `.envrc` and adjust the aws account id mentioned to yours

Then create the monitoring
```shell
terraform init
# this is required because we are relying on a data resource that is not saved in the state yet
terraform apply -target module.gosoline_monitoring_trip_consumer.data.gosoline_application_metadata_definition.main -target module.gosoline_monitoring_trip_forwarder.data.gosoline_application_metadata_definition.main -target module.gosoline_monitoring_trip_gateway.data.gosoline_application_metadata_definition.main -auto-approve
terraform apply -auto-approve
```
