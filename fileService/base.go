package fileService

import (
	pb "github.com/sefaphlvn/suubar/busproto"

	"google.golang.org/grpc"
)

type AgentServer struct {
	pb.UnimplementedFileServiceServer
}

// Bu fonksiyon, gRPC sunucusuna bu servisi kaydeder.
func RegisterWithServer(grpcServer *grpc.Server) {
	pb.RegisterFileServiceServer(grpcServer, &AgentServer{})
}
