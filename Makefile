export PATH := $(GOPATH)/bin:$(PATH)
BIN=./bin
FLAGS=-mod=vendor


ifeq ($(OS),Windows_NT)
	EXT=.exe
	PROTOFLAG=
else
	EXT=
	PROTOFLAG=-I/usr/local/include
endif

all: spider backend

backend: gateway server spider

spider: hacg taohua

hacg:
	go build $(FLAGS) -o $(BIN)/hacg$(EXT) ./spider/hacg

taohua:
	go build $(FLAGS) -o $(BIN)/taohua$(EXT) ./spider/taohua

server:
	go build $(FLAGS) -o $(BIN)/server$(EXT) ./cmd/server

gateway:
	protoc $(PROTOPATH) -I. -I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		./backend/backend.proto
	protoc $(PROTOPATH) -I. -I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--grpc-gateway_out=logtostderr=true:. \
      	./backend/backend.proto
	go build $(FLAGS) -o $(BIN)/gateway$(EXT) ./cmd/gateway


.PHONY: gateway
