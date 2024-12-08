#!/bin/bash

go build ./cmd/protoc-gen-clio-go
mv ./protoc-gen-clio-go ~/.local/bin/protoc-gen-clio-go
buf generate
