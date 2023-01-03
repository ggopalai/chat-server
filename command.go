package main

type commandId int

const (
	CMD_JOIN commandId = iota
	CMD_NICK
	CMD_ROOMS
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandId
	client *client
	args   []string
}
