package main

import (
	"log"
	"net"

	pb "github.com/bryutus/til-golang/grpc/echo/proto"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[echo] ")
}

func main() {
	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterEchoServiceServer(server, &echoService{})
	log.Printf("start server on port %s\n", port)
	if err := server.Serve(listener); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
