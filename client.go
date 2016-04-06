package main;

import (
    "net"
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