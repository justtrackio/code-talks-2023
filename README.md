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

```shell
cd infrastructure/terraform/aws-account-setup
terraform init
terraform apply -auto-approve

# Now let's update the datasource confimap for grafana
echo "writing grafana user credentials into configmap"
echo "GRAFANA_AWS_ACCESS_KEY_ID=$(terraform output -json | jq -r '.grafana_access_key_id.value')"
echo "GRAFANA_AWS_SECRET_ACCESS_KEY=$(terraform output -json | jq -r '.grafana_secret_access_key.value')"
```

#### Setup the aws account
```shell
ROOT_DIR=$PWD
cd infrastructure/terraform/aws-account-setup
terraform init
terraform apply -auto-approve

# Now let's update the datasource confimap for grafana
echo "writing grafana user credentials into configmap"
GRAFANA_AWS_ACCESS_KEY_ID=$(terraform output -json | jq -r '.grafana_access_key_id.value')
GRAFANA_AWS_SECRET_ACCESS_KEY=$(terraform output -json | jq -r '.grafana_secret_access_key.value')

echo "starting aws sso login"
aws sso login --profile ${AWS_ACCOUNT_ID}:AWSAdministratorAccess
AWS_SSO_CACHE_PATH="$HOME/.aws/sso/cache"
LATEST_SSO_SIGNIN_FILE=$(ls -t ${AWS_SSO_CACHE_PATH} | head -n1)
SSO_ACCESS_TOKEN=$(jq -r '.accessToken' ${AWS_SSO_CACHE_PATH}/${LATEST_SSO_SIGNIN_FILE})
CREDS=$(aws sso get-role-credentials --role-name AWSAdministratorAccess --account-id ${AWS_ACCOUNT_ID} --access-token ${SSO_ACCESS_TOKEN})

TMP_AWS_ACCESS_KEY_ID=$(echo -n ${CREDS} | jq -r '.roleCredentials.accessKeyId')
TMP_AWS_SECRET_ACCESS_KEY=$(echo -n ${CREDS} | jq -r '.roleCredentials.secretAccessKey')
TMP_AWS_SESSION_TOKEN=$(echo -n ${CREDS} | jq -r '.roleCredentials.sessionToken')
AWS_CREDS_FILENAME="${ROOT_DIR}/.envrc-aws"
echo "writing creds into ${AWS_CREDS_FILENAME}"
sed -i '' -e "s#AWS_ACCESS_KEY_ID=.*#AWS_ACCESS_KEY_ID=${TMP_AWS_ACCESS_KEY_ID}#" "${AWS_CREDS_FILENAME}"
sed -i '' -e "s#AWS_SECRET_ACCESS_KEY=.*#AWS_SECRET_ACCESS_KEY=${TMP_AWS_SECRET_ACCESS_KEY}#" "${AWS_CREDS_FILENAME}"
sed -i '' -e "s#AWS_SESSION_TOKEN=.*#AWS_SESSION_TOKEN=${TMP_AWS_SESSION_TOKEN}#" "${AWS_CREDS_FILENAME}"
sed -i '' -e "s#GRAFANA_AWS_ACCESS_KEY_ID=.*#GRAFANA_AWS_ACCESS_KEY_ID=${GRAFANA_AWS_ACCESS_KEY_ID}#" "${AWS_CREDS_FILENAME}"
sed -i '' -e "s#GRAFANA_AWS_SECRET_ACCESS_KEY=.*#GRAFANA_AWS_SECRET_ACCESS_KEY=${GRAFANA_AWS_SECRET_ACCESS_KEY}#" "${AWS_CREDS_FILENAME}"

cd -
```

#### Setup and bootstrap the cluster
```shell
cd infrastructure/k8s
k3d cluster create --api-port 6550 -p "8081:80@loadbalancer" --volume "$(realpath ${PWD}/../../volumes):/var/lib/rancher/k3s/storage@all" --wait
kubectl kustomize base/monitoring --context k3d-k3s-default --enable-helm | envsubst | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
kubectl kustomize base/monitoring --context k3d-k3s-default --enable-helm | envsubst | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
kubectl kustomize dev --context k3d-k3s-default --enable-helm | kubectl apply -f - --context k3d-k3s-default --server-side --force-conflicts
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