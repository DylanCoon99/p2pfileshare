package main

// This is an entry point for a peer

import (
	//"time"
	"log"
	"net"
	"github.com/DylanCoon99/p2pfileshare/src/peer"
)





func main() {

	const port = "8080"
	var conn net.Conn

	conn = peer.ConnectToServer(port)

	/*

	if conn == nil {
		log.Fatal("I Failed to connect to the server")
	}


	// Register the peer with the server
	peer.Register(conn)


	conn.Close()

	
	conn = peer.ConnectToServer(port)

	if conn == nil {
		log.Fatal("I Failed to connect to the server")
	}

	
	// Get all peers from the server
	peers,err := peer.GetPeers(conn)

	if err != nil {
		log.Println("Error getting peers:", err)
	}
	log.Printf("I am a peer and I received: %v\n", peers)
	
	*/

	conn.Close()
	
}