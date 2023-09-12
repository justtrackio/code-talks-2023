# Setup

Create the k8s deployment(s) and its dependencies
```shell
terraform init
terraform apply -auto-approve
```

## try out gateway

```shell
curl -X POST -d '{"uuid":"731A313D-D266-4725-BC48-46A8D17304A2","VendorID":1,"tpep_pickup_datetime":"2023-09-01T20:00:00Z","tpep_dropoff_datetime":"2023-09-01T20:00:00Z","passenger_count":300,"trip_distance":3.1415,"RatecodeID":2,"store_and_fwd_flag":"ney","PULocationID":3,"DOLocationID":4,"payment_type":5,"fare_amount":10000000.99,"extra":0.01,"mta_tax":10000,"tip_amount":1.01,"tolls_amount":1.02,"improvement_surcharge":1.03,"total_amount":99999999,"congestion_surcharge":123456,"Airport_fee":5443321}' http://terminal-trips-gateway-traffic.127.0.0.1.nip.io:8081/trip -v
```
