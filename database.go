package main;

// Database contains a centeralized database containing all:
// - Clients
// - Channels
// - Servers
type Database struct {
    Clients map[string]*Client
}

// Initialize allocates memory for all maps
func (db *Database) Initialize() {
    db.Clients = make(map[string]*Client)
}