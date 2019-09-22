package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net"

	"github.com/mongodb/mongo-go-driver/bson"
	pb "github.com/popstk/olddriver/backend"
	"github.com/popstk/olddriver/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	addr   string
	mgoURL string
)

const (
	mgoDB = "spider"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.StringVar(&addr, "addr", ":7700", "search service address")
	flag.StringVar(&mgoURL, "mongo", "mongodb://127.0.0.1:27017/", "mongodb url")
}

type Server struct {
	mgo *mongo.Client
}

func NewServer() (*Server, error) {
	mgo, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mgoURL))
	if err != nil {
		return nil, err
	}

	return &Server{mgo: mgo}, nil
}

func (s *Server) Search(ctx context.Context, in *pb.SearchRequest) (*pb.SearchReply, error) {
	log.Printf(`type="%s"  keyword="%s" page=%d pageSize=%d`,
		in.GetType(), in.GetKeyword(), in.Page, in.PageSize)

	collection := s.mgo.Database(mgoDB).Collection(in.GetType())
	op := options.Find()
	op.SetSkip(int64(in.Page * in.PageSize))
	op.SetLimit(int64(in.PageSize))

	cursor, err := collection.Find(context.Background(),
		bson.D{
			{"title", bson.D{
				{"$regex", ".*" + in.GetKeyword() + ".*"},
				{"$options", "i"},
			}},
		}, op)
	if err != nil {
		return nil, err
	}

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
		item.Tag = in.GetType()
		items = append(items, item)
	}

	return &pb.SearchReply{
		Data: items,
	}, nil
}

func main() {
	flag.Parse()

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("failed to listen: ", err)
		return
	}

	server, err := NewServer()
	if err != nil {
		fmt.Println("failed to new server: ", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterSpiderServer(s, server)
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: ", err)
	}
}
