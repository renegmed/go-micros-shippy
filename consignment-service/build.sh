#!/bin/bash

GOOS=linux GOARCH=amd64 go build 

docker build -t consignment-service .

# mdns(multicast dns) is not typically used for service discovery in production. Instead use Consul or etcd
docker run -p 50051:50051 \
    -e MICRO_SERVER_ADDRESS=:50051 \
    -e MICRO_REGISTRY=mdns consignment-service