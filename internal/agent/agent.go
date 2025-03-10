package agent

import (
	"context"
	"log"

	"vmcontroller/pkg/api"
)

type Agent struct {
	api.UnimplementedAgentServiceServer
}

// ReportResources handles requests from the controller.
func (a *Agent) ReportResources(ctx context.Context, req *api.NodeResources) (*api.VMCommandResponse, error) {
	log.Printf("Received resource report request for node: %s", req.NodeId)
	// In a full implementation, query libvirt and system metrics here.
	return &api.VMCommandResponse{
		Success: true,
		Message: "Resources reported successfully",
	}, nil
}

