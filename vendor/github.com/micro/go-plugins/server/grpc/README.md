# GRPC Server

The grpc server is a [micro.Server](https://godoc.org/github.com/micro/go-micro/server#Server) compatible server.

## Overview

The server makes use of the [google.golang.org/grpc](google.golang.org/grpc) framework for the underlying server 
but continues to use micro handler signatures and protoc-gen-micro generated code.

## Usage

Specify the server to your micro service

```go
import (
        "github.com/micro/go-micro"
        "github.com/micro/go-plugins/server/grpc"
)

func main() {
        service := micro.NewService(
                micro.Name("greeter"),
                micro.Server(grpc.NewServer()),
        )
}
```
