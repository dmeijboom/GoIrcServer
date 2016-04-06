package main;

// Database contains a centeralized database containing all:
// - Clients
// - Channels
// - Servers
type Database struct {
    clients map[string]*Client
}

// Initialize allocates memory for all maps
func (db *Database) Initialize() {
    db.clients = make(map[string]*Client)
}

// GetClients is a wrapper to `clients` but it is mainly abstracted because
// I might implement a real database for IRC
func (db *Database) GetClients() map[string]*Client {
    return db.clients
}