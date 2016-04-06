package main;

import (
    "net"
    "bufio"
    "irc/assert"
    "strings"
    "fmt"
)

// Server defines the IRC server
type Server struct {
    Db *Database
    Hostname string
}

// Connect opens a socket connection
func (server *Server) Connect() {
    server.Db = &Database{}
    server.Db.Initialize()
    server.Hostname = "localhost"
    
    serv, _ := net.Listen("tcp", server.Hostname + ":9000")
    
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
        
        server.HandleMessage(*conn, parsed)
    }
}

// Reply writes output to the connection
func (server *Server) Reply(conn net.Conn, reply *Reply) {
    conn.Write([]byte(reply.Raw() + "\n"))
    
    server.DebugOutput(reply.Raw())
}

// HandleMessage takes action after recieving a parsed message
func (server *Server) HandleMessage(conn net.Conn, message *Message) {
    server.DebugInput(message.Raw())
    
    client, exists := server.Db.GetClient(conn.RemoteAddr().String())
    
    // Check and/or create client
    if !exists {
        client = server.Db.CreateClient(conn)
    }
    
    // Handle the actual message
    switch (message.Command) {
        case "NICK":
            client.Nickname = message.Params[0]
            break
            
        case "PING":
            server.Reply(conn, &Reply{
                Params: []string { "PONG", message.Params[0] },
                Trailing: message.Params[0],
            })
            break
            
        case "QUIT":
            server.Db.DeleteClient(conn.RemoteAddr().String())
            conn.Close()
            break
            
        case "PRIVMSG":
            channel, exists := server.Db.GetChannel(message.Params[0])
            
            if exists {
                for _, user := range channel.GetClients() {
                    if user.Username != client.Username {
                        server.Reply(user.Conn, &Reply{
                            Params: []string{ "PRIVMSG", channel.Name },
                            Trailing: message.Trailing,
                        })
                    }
                }
            } else {
                server.Reply(conn, &Reply{
                    Code: ErrorNoSuchChannel,
                    Params: []string{ client.Nickname, message.Params[0] },
                    Trailing: "No such channel",
                })
            }
            break
            
        case "WHO":
            channel, exists := server.Db.GetChannel(message.Params[0])
            
            if exists {
                for _, user := range channel.GetClients() {
                    server.Reply(conn, &Reply{
                        Code: ReplyWhoReply,
                        Params: []string{ client.Nickname, channel.Name, "~" + user.Username, user.Addr.String(), server.Hostname, user.Nickname, "H+" },
                        Trailing: fmt.Sprintf("%v %v", 0, user.Realname),
                    })
                }
            }
            
            server.Reply(conn, &Reply{
                Code: ReplyEndOfWho,
                Params: []string{ client.Nickname, channel.Name },
                Trailing: "End of WHO list",
            })
            break
            
        // After a user joins a channel a number of steps are executed:
        //
        //  1. Send JOIN confirmation
        //  2. Send channel topic
        //  3. Send a list of nicknames in the channel
        //  4. Send the voice mode for the current user (using the channel service)
        case "JOIN":
            channel, exists := server.Db.GetChannel(message.Params[0])
            
            if exists {
                channel.Join(client)
                
                server.Reply(conn, &Reply{
                    Prefix: fmt.Sprintf("%v!~%v@%v", client.Nickname, client.Username, client.Addr.String()),
                    Params: []string{ "JOIN", channel.Name },
                })
                
                if len(channel.Topic) == 0 {
                    server.Reply(conn, &Reply{
                        Code: ReplyNoTopic,
                        Params: []string{ channel.Name },
                        Trailing: "No topic is set",
                    })
                } else {
                    server.Reply(conn, &Reply{
                        Code: ReplyTopic,
                        Params: []string{ channel.Name },
                        Trailing: channel.Topic,
                    })
                }
                
                server.Reply(conn, &Reply{
                    Code: ReplyNameReply,
                    Params: []string{ client.Nickname, "@", channel.Name },
                    Trailing: strings.Join(channel.GetNicknames(), " "),
                })
                
                server.Reply(conn, &Reply{
                    Code: ReplyEndOfNames,
                    Params: []string{ client.Nickname, channel.Name },
                    Trailing: "End of NAMES list",
                })
                
                server.Reply(conn, &Reply{
                    Prefix: fmt.Sprintf("%v!%v@services.", "ChanServ", "ChanServ"),
                    Params: []string{ "MODE", channel.Name, "+v", client.Nickname },
                })
            } else {
                server.Reply(conn, &Reply{
                    Code: ErrorNoSuchChannel,
                    Params: []string{ client.Nickname, message.Params[0] },
                    Trailing: "No such channel",
                })
            }
            break
            
        case "USER":
            client.Username = message.Params[0]
            client.Mode = message.Params[1]
            client.Realname = message.Trailing
            client.IsRegistered = true
            
            server.Reply(conn, &Reply{ 
                Code: ReplyWelcome,
                Params: []string{ client.Nickname },
                Trailing: "Welcome to the IRC network",
            })
            break
            
        default:
            server.Reply(conn, &Reply{
                Code: ErrorUnknownCommand,
                Params: []string{ message.Command },
                Trailing: "Unknown command",
            })
            break
    }
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