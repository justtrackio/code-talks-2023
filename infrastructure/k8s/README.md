# Setup
Setup and bootstrap the cluster
```shell
# start the cluster
k3d cluster create --api-port 6550 -p "8081:80@loadbalancer" --volume "$(realpath ${PWD}/../../volumes):/var/lib/rancher/k3s/storage@all" --wait

# install all the apps (this will first fail due to CRDs not being installed yet)
kubectl kustomize base/monitoring --context k3d-k3s-default --enable-helm | envsubst | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
kubectl kustomize base/monitoring --context k3d-k3s-default --enable-helm | envsubst | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
kubectl kustomize dev --context k3d-k3s-default --enable-helm | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
```

## Services

### Elasticsearch

#### Credentials
none

#### URL
http://elasticsearch-127.0.0.1.nip.io:8081/

### Kibana

#### Credentials
none

#### URL
http://kibana-127.0.0.1.nip.io:8081/

### fluentD

### fluent-bit

### Grafana

#### Credentials
Username: `admin`
Password: `prom-operator`

#### URL
http://grafana-127.0.0.1.nip.io:8081/dashboards

### Prometheus

#### Credentials
none

#### URL
http://prometheus-127.0.0.1.nip.io:8081/

# Shutdown
```shell
k3d cluster delete
```