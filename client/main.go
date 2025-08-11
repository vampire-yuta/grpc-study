package main

import (
	"context"
	"log"

	pb "main/go_protocol_buffer"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)

	response, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		panic(err)
	}

	log.Printf("Response from server: %s", response.GetMessage())
}
