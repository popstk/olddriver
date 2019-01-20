package main

import (
	"context"
	"log"
	"time"

	pb "github.com/popstk/olddriver/backend"
	"google.golang.org/grpc"
)

const (
	address = "localhost:19527"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSpiderClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	reply, err := c.Search(ctx, &pb.SearchRequest{
		Type:    "taohua",
		Keyword: "fc2ppv",
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, item := range reply.Data {
		log.Print(item.GetTitle())
	}
}
