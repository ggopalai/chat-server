package main

import (
	"fmt"
	"log"
	"net"
	"reflect"
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

func (s *server) run() {
	for cmd := range s.commandChan {
		log.Println("Inside server run function", cmd.id)
		switch cmd.id {
		case CMD_NICK:
			s.name(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.currentRooms(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quitRoom(cmd.client, cmd.args)
		}
	}

}

func (s *server) name(c *client, args []string) {
	name := args[1]
	c.name = name
	c.msg(fmt.Sprintf("Let's call you %s", name))
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]

	r, ok := s.rooms[roomName]
	if !ok {
		r := &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.rooms[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.name))
	c.msg(fmt.Sprintf("Welcome to %s ", r.name))
}

func (s *server) currentRooms(c *client, args []string) {
	rooms := reflect.ValueOf(s.rooms).MapKeys()
	c.msg(fmt.Sprintf("Available rooms - %s", rooms))
}

func (s *server) msg(c *client, args []string) {

}

func (s *server) quitRoom(c *client, args []string) {

}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s left the room", c.name))
	}
}
