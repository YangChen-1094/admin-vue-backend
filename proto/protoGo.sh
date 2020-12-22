#!/bin/bash

#protoc --go_out=../../Protocol/ExternalPb ./External.proto

protoc --go_out=plugins=grpc:./ ./maxSize.proto

