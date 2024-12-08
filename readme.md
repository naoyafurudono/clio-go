# clio-go

A protoc plugin to generate command line interface for a Connect service.

## sample

See https://github.com/naoyafurudono/greeting .

## usage

### Prerequest

This library / command depends on [Connect](https://connectrpc.com/).
Make sure that your service is built as a connect service.

You can [getting started with connect here](https://connectrpc.com/docs/go/getting-started).

### install and setting up

1. install `protoc-gen-clio-go` command
1. setup your `buf.gen.yaml`
1. install `colio-go` library

#### install `protoc-gen-clio-go` command

```sh
go install github.com/naoyafurudono/clio-go/cmd/protoc-gen-clio-go@latest
```

#### setup your `buf.gen.yaml`

```yaml
version: v2
plugins:
  - local: protoc-gen-go
    out: gen
    opt: paths=source_relative
  - local: protoc-gen-connect-go
    out: gen
    opt: paths=source_relative
+   - local: protoc-gen-clio-go
+     out: gen
+     opt: paths=source_relative
```

#### install `colio-go` library

```sh
go get github.com/naoyafurudono/clio-go
```
