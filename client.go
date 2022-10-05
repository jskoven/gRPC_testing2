package main

import (
	"log"

	Chat "github.com/jskoven/gRPC_testing2/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	var connection *grpc.ClientConn
	connection, err := grpc.Dial("192.168.43.169:9000", grpc.WithInsecure())

	_ = err
	defer connection.Close()

	client := Chat.NewChatServiceClient(connection)

	message := Chat.Message{
		Body: "Hello, this is my message",
	}

	response, err := client.SendFunction(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling the function: %s", err)
	}

	log.Printf("Response from server: %s", response.Body)

}
