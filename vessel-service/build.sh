#!/bin/bash


protoc -I .  proto/vessel/vessel.proto \
   --go_out=plugins=micro:${GOPATH}/src/github.com/renegmed/go-micros-shippy/vessel-service 
  
docker build -t vessel-service . 

docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns vessel-service
