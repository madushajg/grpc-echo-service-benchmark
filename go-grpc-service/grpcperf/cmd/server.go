package main

import (
	"hw/pkg/hw"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":9090"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := hw.Server{}

	grpcServer := grpc.NewServer()

	hw.RegisterHelloWorldServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
