package main;

import (
    "net"
    "fmt"
    "strings"
)

type Client struct {
    Nickname string
    Username string
    Realname string
    Mode string
    Addr net.Addr
    Conn net.Conn
    IsRegistered bool
}

// GeneratePrefix generates a prefix ..
func (client *Client) GeneratePrefix() string {
    host := client.Addr.String()
    
    // Strip a port number from the address
    index := strings.IndexByte(host, ':')
    
    if index > -1 {
        host = host[0:index]
    }
    
    return fmt.Sprintf("%v!~%v@%v", client.Nickname, client.Username, host)
}