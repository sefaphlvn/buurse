package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os/exec"

	"github.com/sefaphlvn/suubar/fileService"
	"google.golang.org/grpc"
)

func main() {
	cmd := exec.Command("vtysh", "-c", "show ip route json")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println("Hata:", err)
	}
	fmt.Println("Çıktı:", out.String())

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
