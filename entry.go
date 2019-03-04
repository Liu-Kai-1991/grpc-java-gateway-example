package main

import (
  "flag"
  "net/http"

  "github.com/golang/glog"
  "golang.org/x/net/context"
  "github.com/grpc-ecosystem/grpc-gateway/runtime"
  "google.golang.org/grpc"

  helloWorldGateway "./src/main/proto/org/kai/example"
)

var (
  gRpcServer = flag.String("echo_endpoint", "localhost:9091", "endpoint of YourService")
)

func run() error {
  ctx := context.Background()
  ctx, cancel := context.WithCancel(ctx)
  defer cancel()

  mux := runtime.NewServeMux()
  opts := []grpc.DialOption{grpc.WithInsecure()}

  helloWorldErr := helloWorldGateway.RegisterHelloWorldServiceHandlerFromEndpoint(ctx, mux, *gRpcServer, opts)
  if helloWorldErr != nil {
    println("helloWorldError")
    println(helloWorldErr)
    return helloWorldErr
  }

  return http.ListenAndServe(":8080", mux)
}

func main() {
  flag.Parse()
  defer glog.Flush()

  if err := run(); err != nil {
    glog.Fatal(err)
  }

}