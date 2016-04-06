package main;

import (
    "net"
)

// Database contains a centeralized database containing all:
// - Clients
// - Channels
// - Servers
type Database struct {
    clients map[string]*Client
    channels map[string]*Channel
}

// Initialize allocates memory for all maps
func (db *Database) Initialize() {
    db.clients = make(map[string]*Client)
    db.channels = make(map[string]*Channel)
    
    // Example data
    db.CreateChannel("#test")
}

// GetClients is a wrapper to `clients` but it is mainly abstracted because
// I might implement a real database for IRC
func (db *Database) GetClients() map[string]*Client {
    return db.clients
}

// GetClient ..
func (db *Database) GetClient(addr string) (*Client, bool) {
    client, ok := db.clients[addr]
    return client, ok
}

// DeleteClient removes the client from the entire database
func (db *Database) DeleteClient(addr string) {
    delete(db.clients, addr)
}

//CreateClient creates a new client and adds it to the database
func (db *Database) CreateClient(addr net.Addr) *Client {
    client := Client{ Addr: addr }
    
    db.clients[addr.String()] = &client
    
    return db.clients[addr.String()]
}

// GetChannels returns all channels
func (db *Database) GetChannels() map[string]*Channel {
    return db.channels
}

// GetChannel gets an existing channel by it's name
func (db *Database) GetChannel(name string) (*Channel, bool) {
    channel, ok := db.channels[name]
    return channel, ok
}

// CreateChannel creats a new channel and adds it to the channels
func (db *Database) CreateChannel(name string) *Channel {
    channel := Channel{ Name: name }
    channel.clients = make(map[string]*Client)
    
    db.channels[name] = &channel
    
    return db.channels[name]
}