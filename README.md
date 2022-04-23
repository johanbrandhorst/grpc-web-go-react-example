# Small Go/React/TypeScript gRPC-Web example

This is a tiny example set up for using gRPC-Web with Go and React, powered
by https://github.com/timostamm/protobuf-ts for the client and grpc-go and
https://github.com/improbable-eng/grpc-web for the server.

Bundling is handled by esbuild.github.io.

## Requirements

1. Go
2. Docker (to run npm 🤢)

## Install tools

```shell
$ make install
```

## Running

```shell
$ make serve
go run main.go
Serving on http://0.0.0.0:8080
```

## Regenerate protobuf

```shell
$ make generate
buf generate
```

## Bundle JS

```shell
$ make bundle
esbuild ./frontend/app.tsx --bundle --minify --sourcemap --outfile=./dist/js/bundle.js

  dist/js/bundle.js      189.0kb
  dist/js/bundle.js.map  559.6kb

⚡ Done in 61ms
```
