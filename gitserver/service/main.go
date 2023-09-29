package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	gitserver "github.com/jdvgh/brgo-cd/gitserver"
)

const (
	grpcPort    int    = 50051
	grpcAddress string = ":50051"
)

type Server struct {
    gitserver.UnimplementedCloneRepoServiceServer

	port              int
}

func main() {
	// gRPC
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	gitServer := Server{}

    gitserver.RegisterCloneRepoServiceServer(grpcServer, &gitServer)

	fmt.Printf("Starting gRPC-server. Listening on port: %d\n", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
