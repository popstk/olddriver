export PATH := $(GOPATH)/bin:$(PATH)
BIN=./bin

ifeq ($(OS),Windows_NT)
	EXT=.exe
else
	EXT=
endif

spider: hacg taohua

hacg:
	go build -o $(BIN)/hacg$(EXT) ./spider/hacg

taohua:
	go build -o $(BIN)/taohua$(EXT) ./spider/taohua

