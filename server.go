package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type server struct {
	roomMap       map[string]*room
	commandChan chan command
}

// returns the main server object
func newServer() *server {
	return &server{
		roomMap:       make(map[string]*room),
		commandChan: make(chan command),
	}
}

// create a new client using the injected connection object
func (s *server) newClient(conn net.Conn) *client {
	log.Printf("new client connected : %s", conn.RemoteAddr().String())
	return &client{
		conn:        conn,
		name:        "anonymous",
		commandChan: s.commandChan,
	}
}

// runs infinitely as a goroutine and listens to commands from clients
// before passing it on through the channel.
func (s *server) run() {
	for cmd := range s.commandChan {
		switch cmd.id {
		case CMD_NICK:
			s.name(cmd.client, cmd.args)
		case CMD_JOIN:
			s.join(cmd.client, cmd.args)
		case CMD_ROOMS:
			s.currentRooms(cmd.client)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
			s.quitRoom(cmd.client, cmd.args)
		}
	}

}

// set the name and write to client connection
func (s *server) name(c *client, args []string) {
	name := args[1]
	c.name = name
	c.msg(fmt.Sprintf("Successfully set name, hey %s", name))
}

// client joins the given room
// new room is created if it doesn't exist
func (s *server) join(c *client, args []string) {
	roomName := args[1]

	r, ok := s.roomMap[roomName]
	if !ok {
		log.Println("Creating new room", roomName)
		r = &room{
			name:    roomName,
			members: make(map[net.Addr]*client),
		}
		s.roomMap[roomName] = r
	}

	r.members[c.conn.RemoteAddr()] = c

	s.quitCurrentRoom(c)

	c.room = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.name))
	c.msg(fmt.Sprintf("Welcome to %s ", r.name))
}

// list out all the currently active rooms 
func (s *server) currentRooms(c *client) {
	var rooms []string
	for room := range s.roomMap {
		rooms = append(rooms, room)
	}
	res := strings.Join(rooms, " ")
	c.msg(fmt.Sprintf("Available rooms - %s", res))
}

// sends a message to the room
func (s *server) msg(c *client, args []string) {
	msg := strings.Join(args[1:], " ")
	if len(strings.TrimSpace(msg)) == 0 {
		c.msg("Cant send empty message!")
		return
	}
	currRoom := c.room
	currRoom.broadcast(c, fmt.Sprintf("%s: "+strings.Join(args[1:], " "), c.name))
}

// quits the connection completely, should not be named quitRoom?
func (s *server) quitRoom(c *client, args []string) {
	c.msg("Bye, hope to see you soon.")
	s.quitCurrentRoom(c)

	// Closes the connection completely. Have another option to quit just the room.
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.room != nil {
		delete(c.room.members, c.conn.RemoteAddr())
		c.room.broadcast(c, fmt.Sprintf("%s left the room", c.name))
	}
}
