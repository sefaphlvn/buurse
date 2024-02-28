package main

import (
	"log"
	"net"

	"github.com/sefaphlvn/suubar/fileService"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50041")
	if err != nil {
		log.Fatalf("Port not listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// Servisleri gRPC sunucusuna kaydet
	fileService.RegisterWithServer(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server not started: %v", err)
	}
}
