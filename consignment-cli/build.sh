#!/bin/bash

GOOS=linux GOARCH=amd64 go build 

docker build -t consignment-cli .

# mdns(multicast dns) is used to ensure both service and client are running in 'Dockerland'
# so they are both running on the same host, and using the same network layer.
docker run  -e MICRO_REGISTRY=mdns consignment-cli