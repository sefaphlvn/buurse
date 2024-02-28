package fileService

import (
	"context"

	pb "github.com/sefaphlvn/suubar/busproto"
)

func (s *AgentServer) UpdateBootstrap(ctx context.Context, req *pb.Bootstrap) (*pb.FileResponse, error) {
	// Bootstrap güncelleme işlemleri
	return &pb.FileResponse{Message: "Bootstrap updated"}, nil
}
