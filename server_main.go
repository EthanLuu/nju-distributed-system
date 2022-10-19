package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"rpc-example/mock"
	"rpc-example/service"
)

func main() {
	// Read inline input
	var port string
	flag.StringVar(&port, "p", "2333", "Port number for server, default 2333")
	flag.Parse()

	// Register the auth and time service
	rpc.RegisterName("AuthService", &service.AuthService{AllowedTokens: mock.MockTokens})
	rpc.RegisterName("TimeService", &service.TimeService{})
	socket, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal("Register server error, please check your port", err)
	}
	fmt.Println("Time Server has started...")

	// Wait for new connections
	defer socket.Close()
	for {
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		go jsonrpc.ServeConn(conn)
	}
}
