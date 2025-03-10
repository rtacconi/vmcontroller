package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"vmcontroller/internal/controller"
	"vmcontroller/pkg/api"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	// The same controller handles both VM commands and agent reports.
	ctrl := &controller.Controller{}
	api.RegisterVMControllerServer(grpcServer, ctrl)
	api.RegisterAgentServiceServer(grpcServer, ctrl)

	log.Println("Controller gRPC server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

