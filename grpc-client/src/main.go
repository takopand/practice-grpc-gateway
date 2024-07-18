package main

import (
	"client/hello"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	address := "server:8080"
	conn, err := grpc.NewClient(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		log.Fatal("Connection failed.")
		return
	}
	defer conn.Close()

	client := hello.NewHelloWorldClient(conn)

	ctx := context.Background()
	req := hello.HelloWorldRequest{
		Message: "Takopanda",
	}

	res, err := client.GetHelloWorld(ctx, &req)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("res: %v\n", res)

}
