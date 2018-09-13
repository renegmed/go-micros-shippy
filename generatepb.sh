#!/bin/bash

protoc -I consignment-service/  consignment-service/proto/consignment/consignment.proto \
--go_out=plugins=micro:${GOPATH}/src/github.com/renegmed/go-micros-shippy/consignment-service 