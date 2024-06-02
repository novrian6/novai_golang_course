package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "simple_grpc/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()
	client := pb.NewHelloWorldClient(conn)

	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
	if err != nil {
		log.Fatalf("Error calling SayHello: %v", err)
	}
	log.Printf("Response: %s", resp.GetMessage())
}
