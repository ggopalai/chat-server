package main

type commandId int

const (
	CMD_JOIN commandId = iota
	CMD_NICK
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

// command struct contains the type of command, 
// the pointer to the client issuing the command, 
// and the respective arguments. 
type command struct {
	id     commandId
	client *client
	args   []string
}
