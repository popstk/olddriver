package main

import (
	"context"
	"log"
	"net"

	pb "github.com/popstk/olddriver/backend"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":19527"
)

type server struct{}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	return &pb.SearchReply{
		Title: "Test",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSpiderServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
