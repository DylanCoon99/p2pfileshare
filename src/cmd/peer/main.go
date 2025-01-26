package main

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

	if conn == nil {
		log.Fatal("I Failed to connect to the server")
	}

	peer.Register(conn)

}