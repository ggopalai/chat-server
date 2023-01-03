package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("TCP Chat Server - WIP")
	serv := newServer()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Can't start server - %s", err.Error())
	}

	defer listener.Close()
	log.Printf("Started server on :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection.")
			continue
		}

		newClient := serv.newClient(conn)
		go newClient.readInput()
	}
}
