package main

import (
	"context"

	pb "github.com/bryutus/til-golang/grpc/echo/proto"
)

type echoService struct{}

func (s *echoService) Echo(ctx context.Context,
	req *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: req.GetMessage()}, nil
}
