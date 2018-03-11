package main

import (
  "context"
  "net/http"

  pb "github.com/yoshikazu-ooba/twirp-sample/rpc/helloworld"
)

type HelloWorldServer struct {}

func (s *HelloWorldServer) Hello(context context.Context, request *pb.HelloRequest) (response *pb.HelloResponse, err error) {
  return &pb.HelloResponse{Text: "Hello " + request.GetSubject()}, nil
}

func main() {
  twirpHandler := pb.NewHelloWorldServer(&HelloWorldServer{}, nil)
  mux := http.NewServeMux()
  mux.Handle(pb.HelloWorldPathPrefix, twirpHandler)
  http.ListenAndServe(":8080", mux)
}
