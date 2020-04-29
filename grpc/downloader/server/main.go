package main

import (
	"log"
	"net"

	pb "github.com/bryutus/til-golang/grpc/downloader/proto"
	"google.golang.org/grpc"
)

const port = ":50051"

func init() {
	log.SetFlags(0)
	log.SetPrefix("[file] ")
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	server := grpc.NewServer()
	pb.RegisterFileServiceServer(server, &fileService{})
	log.Printf("start server on port %s\n", port)
	if err := server.Serve(listener); err != nil {
		log.Printf("failed to serve: %v\n", err)
	}
}
