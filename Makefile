export PATH := $(GOPATH)/bin:$(PATH)
BIN=./bin
FLAGS=-mod=vendor

ifeq ($(OS),Windows_NT)
	EXT=.exe
else
	EXT=
endif

spider: hacg taohua

hacg:
	go build $(FLAGS) -o $(BIN)/hacg$(EXT) ./spider/hacg

taohua:
	go build $(FLAGS) -o $(BIN)/taohua$(EXT) ./spider/taohua

client:
	go build $(FLAGS) -o $(BIN)/client$(EXT) ./rpc/client

server:
	go build $(FLAGS) -o $(BIN)/server$(EXT) ./rpc/server

gateway:
	protoc -I. \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=Mgoogle/api/annotations.proto=github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api,plugins=grpc:. \
		./backend/backend.proto
	protoc -I. -I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--grpc-gateway_out=logtostderr=true:. \
      	./backend/backend.proto
	go build $(FLAGS) -o $(BIN)/gateway$(EXT) ./cmd/gateway


.PHONY: gateway
