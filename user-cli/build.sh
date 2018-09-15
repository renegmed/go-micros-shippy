#!/bin/bash
  
  
docker build -t user-cli . 

docker run -e MICRO_REGISTRY=mdns user-cli
