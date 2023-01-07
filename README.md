# tcp-chat-server
TCP Chat Server

Uses "net" package to establish TCP connections. Each client is handled via goroutines and channels for client-server communication.

## Usage 

### Start Server
`go build`
`./tcp-chat`

### Establish client connection
`telnet localhost 8888`

### Client Commands 
`/name <name>`
`/rooms`
`/join <room-name>`
`/msg <msg>`
`/quit`



