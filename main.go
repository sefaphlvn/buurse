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
	cmd := exec.Command("sudo /usr/bin/vtysh", "-c", "show ip bgp summary")

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out    // Standart çıktıyı yakala
	cmd.Stderr = &stderr // Hata çıktısını yakala

	err := cmd.Run()
	if err != nil {
		fmt.Println("Komut Hatası:", err)
		fmt.Println("STDERR:", stderr.String()) // Hata çıktısını yazdır
		return
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
