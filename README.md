# Setup

* First go to `infrastructure/terraform/aws-account-setup/README.md` and follow the `Setup` section
* Then go to `k8s/README.md` and follow the `Setup` section
* Now go to `infrastructure/terraform/apps/README.md` and follow the `Setup` section
* Eventually go to `infrastructure/terraform/monitoring/README.md` and follow the `Setup` section

## Summary

#### Setup direnv files
```shell
sed -i '' -e 's#AWS_ACCOUNT_ID=.*#AWS_ACCOUNT_ID=1234567890000#' .envrc.template > .envrc
direnv reload
```

#### Setup the aws account
```shell
cd infrastructure/terraform/aws-account-setup
terraform init
terraform apply -auto-approve
cd -
```

#### Setup and bootstrap the cluster
```shell
cd infrastructure/k8s
k3d cluster create --api-port 6550 -p "8081:80@loadbalancer" --volume "$(realpath ${PWD}/../../volumes):/var/lib/rancher/k3s/storage@all" --wait
kubectl kustomize dev --context k3d-k3s-default --enable-helm | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
kubectl kustomize base/monitoring --context k3d-k3s-default --enable-helm | envsubst | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
kubectl kustomize base/monitoring --context k3d-k3s-default --enable-helm | envsubst | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
cd -
```

#### Create the k8s deployment(s) and its dependencies
```shell
cd infrastructure/terraform/apps
terraform init
terraform apply -auto-approve
cd -
```

#### Finally create the monitoring
```shell
cd infrastructure/terraform/monitoring
terraform apply -target module.gosoline_monitoring_trip_consumer.data.gosoline_application_metadata_definition.main -target module.gosoline_monitoring_trip_forwarder.data.gosoline_application_metadata_definition.main -target module.gosoline_monitoring_trip_gateway.data.gosoline_application_metadata_definition.main -auto-approve
terraform apply -auto-approve
cd -
```

http://elasticsearch-127.0.0.1.nip.io:8081/  
http://kibana-127.0.0.1.nip.io:8081/  
http://prometheus-127.0.0.1.nip.io:8081/  

http://grafana-127.0.0.1.nip.io:8081/dashboards  
Username: `admin`  
Password: `prom-operator`  

#### Shutdown cluster
```shell
k3d cluster delete
```