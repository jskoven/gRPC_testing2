package main

import (
	"log"
	"net"

	"github.com/jskoven/gRPC_testing2/chat"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("Failed to start listener on port :9000 %t", err)
	}

	serverStruct := chat.Server{}

	gRPCserver := grpc.NewServer()
	if err := gRPCserver.Serve(listener); err != nil {
		log.Fatalf("Failed to serve")
	}
	chat.RegisterChatServiceServer(gRPCserver, &serverStruct)

}
