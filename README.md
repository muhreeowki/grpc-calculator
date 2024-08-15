# Calculator Service

This repo contains the source code for a Calculator API gRPC server I built to learn the basics of gRPC.

## Requirements

To run this code, you'll need the following dependencies:

- Go 1.22.5
- Make

To make modifications to this code, make sure you have protobuf tools installed, as well as grpcui for client.

### Linux

Install protobuf tools:
```
$ sudo apt install -y protobuf-compiler
```

Install grpcui:
```
$ go install github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

alternatively you can install the protoc plugins using golang

```
$ go install google.golang.org/grpc/cmd/protoc-gen-go@v1.1
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```


## Building

To build this code, run `go build .` or `make build`.

If you make any modifications to the protobuf, run `make generate`

## Running Locally

To run the server, cd into the server directory and run `go run main.go`

To interact with the server with grpcui run `grpcui --plaintext 127.0.0.1:8080`
