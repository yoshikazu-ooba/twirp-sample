package main

import (
  "context"
  "fmt"
  "log"
  "net/http"
  "os"

  pb "github.com/yoshikazu-ooba/twirp-sample/rpc/helloworld"
)

func main() {
  client := pb.NewHelloWorldProtobufClient("http://localhost:8080", &http.Client{})

  subject := ""
  if len(os.Args) > 1 {
    subject = os.Args[1]
  }

  response, err := client.Hello(context.Background(), &pb.HelloRequest{Subject: subject})
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(response.GetText())
}
