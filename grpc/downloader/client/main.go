package main

import (
	"context"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"

	pb "github.com/bryutus/til-golang/grpc/downloader/proto"
	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[file] ")
}

func main() {
	target := "localhost:50051"
	connect, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s\n", err)
	}
	defer connect.Close()

	c := pb.NewFileServiceClient(connect)
	name := os.Args[1]
	ctx, cancel := context.WithTimeout(
		context.Background(), 3*time.Second)
	defer cancel()

	stream, err := c.Download(ctx, &pb.FileRequest{
		Name: name})
	if err != nil {
		log.Fatalf("could not download: %s\n", err)
	}
	var blob []byte
	for {
		c, err := stream.Recv()
		if err == io.EOF {
			log.Printf("done %d bytes\n", len(blob))
			break
		}

		if err != nil {
			log.Fatalln(err)
		}
		blob = append(blob, c.GetData()...)
	}

	ioutil.WriteFile(name, blob, 0644)
}
