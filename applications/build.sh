#!/usr/bin/env zsh
echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
docker build -t ghcr.io/justtrackio/code-talks-2023-trip-gateway:1.0.0 . --build-arg APP=trip-gateway
docker push ghcr.io/justtrackio/code-talks-2023-trip-gateway:1.0.0
docker build -t ghcr.io/justtrackio/code-talks-2023-trip-consumer:1.0.0 . --build-arg APP=trip-consumer
docker push ghcr.io/justtrackio/code-talks-2023-trip-consumer:1.0.0
docker build -t ghcr.io/justtrackio/code-talks-2023-trip-forwarder:1.0.0 . --build-arg APP=trip-forwarder
docker push ghcr.io/justtrackio/code-talks-2023-trip-forwarder:1.0.0
