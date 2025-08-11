package main

import (
	"context"
	"fmt"
	"log"
	pb "main/go_protocol_buffer"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHelloServiceServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{Message: "Hello " + in.GetName() + "!"}, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHelloServiceServer(grpcServer, &server{})

	fmt.Println("Server is running on localhost:8081")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
