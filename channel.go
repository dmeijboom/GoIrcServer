package main;

type Channel struct {
    Name string
    Topic string
    clients map[string]*Client
}

// Join adds a client to the channel
func (channel *Channel) Join(client *Client) {
    channel.clients[client.Addr.String()] = client
}

// Leave (or `part` in IRC) kicks a client from the channel
func (channel *Channel) Leave(client *Client) {
    delete(channel.clients, client.Addr.String())
}

// GetNicknames Generates a list of nicknames from clients in this channel
func (channel *Channel) GetNicknames() []string {
    nicknames := []string{}
    
    for _, client := range channel.clients {
        nicknames = append(nicknames, client.Nickname)
    }
    
    // Also add the `ChanServ` service
    nicknames = append(nicknames, "@ChanServ")
    
    return nicknames
}

// GetClients returns all clients within this channel
func (channel *Channel) GetClients() map[string]*Client {
    return channel.clients
}