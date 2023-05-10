package client

import (
	"log"

	"github.com/AmadoJunior/gRPCTest/chat"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn          *grpc.ClientConn
	connErr       error
	connPort      string
	serviceClient chat.ChatServiceClient
	context       context.Context
}

func NewClient(connPort string) *Client {
	return &Client{connPort: connPort}
}

func (c *Client) InitClient() {
	c.conn, c.connErr = grpc.Dial(c.connPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if c.connErr != nil {
		log.Fatalf("Could Not Connect: %s", c.connErr)
	}

	c.serviceClient = chat.NewChatServiceClient(c.conn)

	c.context = context.Background()
}

func (c *Client) CloseClient() {
	c.conn.Close()
}

func (c *Client) SendMessage(msg string) {
	message := chat.Message{
		Body: msg,
	}

	response, err := c.serviceClient.SayHello(c.context, &message)

	if err != nil {
		log.Fatalf("Error Calling SayHello: %s", err)
	}

	log.Printf("Response from Server: %s", response.Body)
}
