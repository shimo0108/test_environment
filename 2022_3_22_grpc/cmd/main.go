package main

import (
	"fmt"
	"log"
	"net"

	pb "test_environment/2022_3_22_grpc/proto"
	"test_environment/2022_3_22_grpc/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := 50051
	listenPort, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, service.NewUserService())

	reflection.Register(server)
	server.Serve(listenPort)
}
