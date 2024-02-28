package fileService

import (
	"context"

	fileService "github.com/sefaphlvn/suubar/busproto"
)

func (s *AgentServer) UpdateBootstrap(ctx context.Context, req *fileService.Bootstrap) (*fileService.FileResponse, error) {
	// Bootstrap güncelleme işlemleri
	return &fileService.FileResponse{Message: "Bootstrap updated"}, nil
}
