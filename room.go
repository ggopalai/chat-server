package main

import "net"

// Struct representing a chat-room.
// Members is a hashmap of connection to the client object.
type room struct {
	name    string
	members map[net.Addr]*client
}

// Function to broadcast a message to all members part of the client's room.
func (r *room) broadcast(sender *client, msg string) {
	for adr, c := range r.members {
		if adr != sender.conn.RemoteAddr() {
			c.msg(msg)
		}
	}
}
