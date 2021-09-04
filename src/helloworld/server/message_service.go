package main

import (
	"context"
	message "grpc-hello/src/helloworld/pb"
	"log"
)

type messageServiceImpl struct {
	message.UnimplementedMessageSenderServer
}

func (s *messageServiceImpl) Send(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	log.Println("receive message:", req.GetSaySomething())
	resp := &message.MessageResponse{}
	resp.ResponseSomething = "roger that!"
	log.Println(ctx)
	return resp, nil
}
