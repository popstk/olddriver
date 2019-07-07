package main

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"

	gw "github.com/popstk/olddriver/backend"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint",  ":19527", "gRPC server endpoint")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := gw.RegisterSpiderHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err!= nil {
		return err
	}

	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()

	if err:= run(); err != nil {
		log.Println(err)
	}
}
