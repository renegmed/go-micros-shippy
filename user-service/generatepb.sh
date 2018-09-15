#!/bin/bash
 
protoc -I .  proto/user/user.proto \
   --go_out=plugins=micro:${GOPATH}/src/github.com/renegmed/go-micros-shippy/user-service 