package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Vishal-2029/grpcapp/pb/github.com/Vishal-2029/grpcapp/pb"

 
	"google.golang.org/grpc"
)

type server struct { 
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request from: %s", req.Name)
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})

	fmt.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
