package main

import (
	"flag"
	"net"
	"net/http"
)

var endpoint string

func init() {
	flag.StringVar(&endpoint, "e", ":8080", "")
}

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./web/dist")))

	lis, err := net.Listen("tcp", endpoint)
	Must(err)

	Must(http.Serve(lis, mux))
}

