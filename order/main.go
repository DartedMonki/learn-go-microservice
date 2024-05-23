package main

import (
	"log"
	"net"

	common "github.com/dartedmonki/learn-go-microservice/common"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

var (
	grpcAddress = common.EnvString("GRPC_ADDR", "localhost:3000")
)

func main() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer listener.Close()

	store := NewStore()
	service := NewService(store)
	NewGrpcHandler(grpcServer, service)

	log.Printf("Start gRPC server at port %s", grpcAddress)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
