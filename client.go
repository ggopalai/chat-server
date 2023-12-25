package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn        net.Conn
	name        string
	room        *room
	commandChan chan<- command
}

// listen to the client connection, 
// parse the input and 
// put the command onto the channel
func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/name":
			c.commandChan <- command{
				id:     CMD_NICK,
				client: c,
				args:   args,
			}
		case "/join":
			c.commandChan <- command{
				id:     CMD_JOIN,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commandChan <- command{
				id:     CMD_ROOMS,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commandChan <- command{
				id:     CMD_MSG,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commandChan <- command{
				id:     CMD_QUIT,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("Unknown Command %s", cmd))
		}
	}
}

// write error to connection
func (c *client) err(err error) {
	c.conn.Write([]byte("Error: " + err.Error() + "\n"))
}

// write message to connection 
func (c *client) msg(msg string) {
	c.conn.Write([]byte(">> " + msg + "\n"))
}
