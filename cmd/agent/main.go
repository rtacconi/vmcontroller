package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"vmcontroller/internal/agent"
	"vmcontroller/pkg/api"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	api.RegisterAgentServiceServer(grpcServer, &agent.Agent{})

	log.Println("Agent gRPC server listening on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

