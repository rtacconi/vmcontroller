package controller

import (
	"context"
	"log"

	"vmcontroller/pkg/api"
)

type Controller struct {
	api.UnimplementedVMControllerServer
	api.UnimplementedAgentServiceServer
}

// ExecuteCommand receives VM commands from the client.
func (c *Controller) ExecuteCommand(ctx context.Context, req *api.VMCommandRequest) (*api.VMCommandResponse, error) {
	log.Printf("Received command: %s for VM ID: %s", req.Command.String(), req.VmId)
	// Here you would add logic to choose an agent, dispatch the command via serf or other means, etc.
	// For now, we simply acknowledge the command.
	return &api.VMCommandResponse{
		Success: true,
		Message: "Command executed successfully",
	}, nil
}

// ReportResources handles resource reports from agents.
func (c *Controller) ReportResources(ctx context.Context, req *api.NodeResources) (*api.VMCommandResponse, error) {
	log.Printf("Resource report from node %s: CPU=%d, RAM=%d MB, Storage=%d GB", req.NodeId, req.CpuCores, req.RamMb, req.StorageGb)
	// You might store or act on these metrics here.
	return &api.VMCommandResponse{
		Success: true,
		Message: "Resource data received",
	}, nil
}

