package server

import (
	"log"
	"net"

	"github.com/AmadoJunior/gRPCTest/chat"
	"google.golang.org/grpc"
)

type Server struct {
	listener     net.Listener
	listenerType string
	listenerPort string
	listenerErr  error
	chatServer   *chat.Server
	grpcServer   *grpc.Server
}

func NewServer(listenerType string, listenerPort string) *Server {
	return &Server{listenerType: listenerType, listenerPort: listenerPort}
}

func (s *Server) InitServer() {
	s.listener, s.listenerErr = net.Listen(s.listenerType, s.listenerPort)
	if s.listenerErr != nil {
		log.Fatalf("Failed to Listen: %v", s.listenerErr)
	}

	s.chatServer = &chat.Server{}

	s.grpcServer = grpc.NewServer()

	chat.RegisterChatServiceServer(s.grpcServer, s.chatServer)

	if err := s.grpcServer.Serve(s.listener); err != nil {
		log.Fatalf("Failed to Server gRPC: %v", err)
	}
}
