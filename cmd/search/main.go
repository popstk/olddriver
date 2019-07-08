package main

import (
	"context"
	"log"
	"net"

	"github.com/mongodb/mongo-go-driver/bson"
	pb "github.com/popstk/olddriver/backend"
	"github.com/popstk/olddriver/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = "127.0.0.1:19527"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)
}

type server struct{}

func (s *server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	log.Printf(`type="%s"  keyword="%s"`, in.GetType(), in.GetKeyword())

	collection, err := core.Collection(in.GetType())
	if err != nil {
		return nil, err
	}

	cursor, err := collection.Find(context.Background(),
		bson.D{
			{"title", bson.D{
				{"$regex", ".*" + in.GetKeyword() + ".*"},
				{"$options", "i"},
			}},
		}, nil)
	if err != nil {
		return nil, err
	}

	log.Println("cursor ok -> ", cursor)

	items := make([]*pb.Item, 0)
	for cursor.Next(context.Background()) {
		var tmp core.Item
		err = cursor.Decode(&tmp)
		if err != nil {
			log.Println(err)
			continue
		}
		item, err := core.ToProtoItem(&tmp)
		if err != nil {
			log.Println(err)
			continue
		}
		items = append(items, item)
	}

	return &pb.SearchReply{
		Data: items,
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
