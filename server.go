package main

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
