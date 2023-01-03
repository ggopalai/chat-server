package main

import "net"

type client struct {
	conn        net.Conn
	name        string
	room        *room
	commandChan chan<- command
}
