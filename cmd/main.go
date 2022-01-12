package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/doniacld/grpc-in-go/cmd/api/protos"
)

func main() {
	log.Println("Build a mood tracker using Go, gRPC and more!")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err.Error())
	}

	s := protos.Server{}

	grpcServer := grpc.NewServer()

	protos.RegisterTrackerAPIServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err.Error())
	}
}
