package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Vishal-2029/grpcapp/pb/github.com/Vishal-2029/grpcapp/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Gopher"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}

	fmt.Println("Server Response:", res.Message)
}
