package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for adr, c := range r.members {
		if adr != sender.conn.RemoteAddr() {
			c.msg(msg)
		}
	}
}
