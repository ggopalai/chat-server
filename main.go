package main

import (
	"log"
	"net"
)

func main() {
	serv := newServer()

	// default room
	serv.roomMap["general"] = &room{
		name:    "general",
		members: make(map[net.Addr]*client),
	}

	// listens to the channel for commands from the connected clients 
	go serv.run()

	// binds the server to the machine's port 8888
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Can't start server - %s", err.Error())
	}
	defer listener.Close()
	log.Printf("Started server on :8888")

	// keep listening for newer client connections 
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
