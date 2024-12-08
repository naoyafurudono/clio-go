#!/bin/bash

go build ./cmd/protoc-gen-clio-go
mv ./protoc-gen-clio-go ~/.local/bin/protoc-gen-clio-go
protoc --clio-go_out=gen greet/v1/greet.proto 
