package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"server/hello"

	"google.golang.org/grpc"
)

type HelloServer struct{}

func (h *HelloServer) GetHelloWorld(ctx context.Context, req *hello.HelloWorldRequest) (res *hello.HelloWorldResponse, e error) {
	text := fmt.Sprintf("Hello %s", req.Message)
	res = &hello.HelloWorldResponse{Message: text}
	fmt.Printf("Request: %v\n", req)
	return res, nil
}

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	hello.RegisterHelloWorldServer(s, &HelloServer{})

	s.Serve(listenPort)
}
