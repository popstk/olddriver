package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/popstk/olddriver/core"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"

	pb "github.com/popstk/olddriver/backend"
)

var (
	addr     string
	frontend string
	dev      bool
	searchSrv string
)

func init() {
	flag.StringVar(&addr, "addr", ":0", "")
	flag.StringVar(&frontend, "frontend", "http://127.0.0.1:8080", "Reverse proxy for frontend")
	flag.StringVar(&searchSrv, "query", ":7700", "query service grpc server address")
	flag.BoolVar(&dev, "dev", false, "dev mode: proxy <frontend>")
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	mux := http.NewServeMux()

	// static files
	handler := http.FileServer(http.Dir("./web/dist/"))
	if dev {
		fmt.Println("frontend -> ", frontend)
		uri, err := url.Parse(frontend)
		core.Must(err)
		handler = httputil.NewSingleHostReverseProxy(uri)
	}
	mux.Handle("/", handler)

	// rpc service
	opts := []grpc.DialOption{grpc.WithInsecure()}
	rpcMux := runtime.NewServeMux()
	Must(pb.RegisterSpiderHandlerFromEndpoint(context.TODO(), rpcMux, searchSrv, opts))
	mux.Handle("/v1/", rpcMux)

	lis, err := net.Listen("tcp", addr)
	Must(err)

	fmt.Println("listen -> ", lis.Addr())
	if err = http.Serve(lis, mux); err != nil {
		fmt.Println(err)
	}
}
