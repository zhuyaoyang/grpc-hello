package main

import (
	"fmt"
	message "grpc-hello/src/helloworld/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("run success")
	srv := grpc.NewServer()
	message.RegisterMessageSenderServer(srv, &messageServiceImpl{})
	//RegisterMessageSenderService(srv, &message.MessageSenderService{})
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = srv.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
