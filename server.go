package main

import (
	"log"
	"net"
)

type server struct {
	rooms       map[string]*room
	commandChan chan command
}

func newServer() *server {
	return &server{
		rooms:       make(map[string]*room),
		commandChan: make(chan command),
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("new client connected : %s", conn.RemoteAddr().String())
	return &client{
		conn:        conn,
		name:        "anonymous",
		commandChan: s.commandChan,
	}
}
