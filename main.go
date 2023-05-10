package main

import (
	"github.com/AmadoJunior/gRPCTest/client"
	"github.com/AmadoJunior/gRPCTest/server"
)

func main() {
	//Server
	server := server.NewServer("tcp", ":9000")
	go server.InitServer()

	//Client
	client := client.NewClient(":9000")
	client.InitClient()
	client.SendMessage("Hello From Client!")
	client.CloseClient()
}
