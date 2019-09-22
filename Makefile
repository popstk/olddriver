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

all: web backend spider

backend: proto server search

web:
	cd web; cnpm install;cnpm run-script build

spider: hacg taohua

hacg:
	go build $(FLAGS) -o $(BIN)/hacg$(EXT) ./spider/hacg

taohua:
	go build $(FLAGS) -o $(BIN)/taohua$(EXT) ./spider/taohua

search:
	go build $(FLAGS) -o $(BIN)/search$(EXT) ./cmd/search

server:
	go build $(FLAGS) -o $(BIN)/server$(EXT) ./cmd/server

proto:
	protoc $(PROTOPATH) -I. -I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:. \
		./backend/backend.proto
	protoc $(PROTOPATH) -I. -I$(GOPATH)/src \
		-I$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--grpc-gateway_out=logtostderr=true:. \
      	./backend/backend.proto

.PHONY: gateway web
