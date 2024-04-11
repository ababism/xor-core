#!/bin/bash

export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

python3 ./xor-python/scripts/prepare_protos.py --xor-go
