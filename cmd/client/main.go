package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"vmcontroller/pkg/api"
)

func main() {
	// Parse command-line flags.
	command := flag.String("command", "CREATE", "Command to execute: CREATE, DELETE, RESTART, STOP, START")
	vmID := flag.String("vm_id", "", "VM identifier")
	flag.Parse()

	// Connect to the controller.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewVMControllerClient(conn)

	// Map string command to the corresponding enum.
	var cmdType api.CommandType
	switch *command {
	case "CREATE":
		cmdType = api.CommandType_CREATE
	case "DELETE":
		cmdType = api.CommandType_DELETE
	case "RESTART":
		cmdType = api.CommandType_RESTART
	case "STOP":
		cmdType = api.CommandType_STOP
	case "START":
		cmdType = api.CommandType_START
	default:
		cmdType = api.CommandType_UNKNOWN
	}

	// Create and send the command request.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &api.VMCommandRequest{
		Command:    cmdType,
		VmId:       *vmID,
		Parameters: map[string]string{"example": "value"},
	}

	resp, err := client.ExecuteCommand(ctx, req)
	if err != nil {
		log.Fatalf("Error executing command: %v", err)
	}
	log.Printf("Response: success=%v, message=%s", resp.Success, resp.Message)
}

