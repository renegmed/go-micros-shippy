#!/bin/bash

protoc consignment-service/proto/consignment/consignment.proto --go_out=plugins=grpc:.
 
