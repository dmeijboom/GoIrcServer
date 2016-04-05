package main;

import (
    "net"
    "bufio"
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
    fmt.Printf("In:  %v\n", raw)
}

func main() {
    server := Server{}
    server.Connect()
}