#!/bin/bash
 
protoc -I .  proto/vessel/vessel.proto \
   --go_out=plugins=micro:${GOPATH}/src/github.com/renegmed/go-micros-shippy/vessel-service 