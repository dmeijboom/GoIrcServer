package main;

// import (
//     "testing"
//     "os"
// )

// // Mock socket connection
// var server *Server

// func handleMessage(raw string) {
//     message, _ := ParseMessage(raw + "\n")
    
//     server.HandleMessage(message)
// }

// func reset() {
//     server = nil
//     server = &Server{}
// }

// func TestMain(m *testing.M) {
//     os.Exit(m.Run())
// }

// TestRegistration tests wether it's possible to connect and register
// to the IRC server
// func TestRegistration() {
//     reset()
//     handleMessage("NICK dmeijboom")
//     handleMessage("USER dmeijboom 0 *")
// }