package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	message "grpc-hello/src/helloworld/pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:12345", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return
	}
	defer conn.Close()

	client := message.NewMessageSenderClient(conn)
	resp, err := client.Send(context.Background(), &message.MessageRequest{SaySomething: "hello world!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(resp)
}
