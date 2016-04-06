package main;

import (
    "net"
    "bufio"
    "irc/assert"
    "fmt"
)

// Server defines the IRC server
type Server struct {
    Db *Database
}

// Connect opens a socket connection
func (server *Server) Connect() {
    serv, _ := net.Listen("tcp", "localhost:9000")
    
    for {
        conn, _ := serv.Accept()
        
        go server.HandleConnection(&conn)
    }
}

// HandleConnection starts an infinite reading loop and starts the
// message handeling chain
func (server *Server) HandleConnection(conn *net.Conn) {
    reader := bufio.NewReader(*conn)
    
    for {
        line, _, _ := reader.ReadLine()
        parsed, _ := ParseMessage(string(line))
        
        server.HandleMessage(parsed)
    }
}

// HandleMessage takes action after recieving a parsed message
func (server *Server) HandleMessage(message *Message) {
    server.DebugInput(message.Raw())
}

// DebugInput ..
func (server *Server) DebugInput(raw string) {
    fmt.Printf("%vIn%v:  %v\n", assert.BOLD_ON, assert.BOLD_OFF, raw)
}

// DebugOutput ..
func (server *Server) DebugOutput(raw string) {
    fmt.Printf("%vOut%v: %v\n", assert.BOLD_ON, assert.BOLD_OFF, raw)
}

func main() {
    server := Server{}
    server.Connect()
}