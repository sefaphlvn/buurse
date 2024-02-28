package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sefaphlvn/buurse/busproto/bootstrap_grpc.pb"

	"google.golang.org/grpc"
)

func main() {
	// Sunucuyla bağlantı kur
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewBootstrapServiceClient(conn)

	// BootstrapRequest çağrısı yap
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Bootstrap(ctx, &pb.BootstrapRequest{Domain: "example.com"})
	if err != nil {
		log.Fatalf("could not bootstrap: %v", err)
	}
	log.Printf("Bootstrap Response: %s", r.GetMessage())
}
