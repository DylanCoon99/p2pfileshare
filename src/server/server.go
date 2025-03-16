package server

import (
	"net"
	"log"
	"time"
)


const MaxNumPeers = 10


type ServerState struct {
	MaxPeers int
	CurNumPeers int
	Peers *[]Peer
}


type RequestType int

const (
	REGISTER RequestType = iota
	PING
)



type Request struct {
	// type of request sent to the server
	Type RequestType
	Body []byte
	Peer *Peer
	Conn net.Conn
}



type Peer struct {
	// going to have an IP address with port
	// whether the node is active or not

	IP string                    `json:"ip"`  // "ip:port" in string format
	Active bool                  `json:"active"`
	LastServerContact time.Time  `json:"lastservercontact"`
}



func StartServer(port string) {



	// this will be a server that exists purely for peer discovery.

	// a node will register with the server

	// the node can query the server for a list of peers

	// the server sends the list of peers

	// the server can periodically ping nodes to ensure they are still alive

	// need to implement REGISTER and GET_PEERS requests over TCP




	// instantiate a serverstate

	peers := make([]Peer, MaxNumPeers)

	serverCfg := ServerState{
		MaxPeers: MaxNumPeers,
		CurNumPeers: 0,
		Peers: &peers,
	}


	// instantiate tcp server here
	ln, err := net.Listen("tcp", ":" + port)  // returns (Listener, error)


	if err != nil {
		log.Fatal(err) // prints the error and exits the program
	}
	log.Println("Server: I am listening on port 8080.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err) // prints the error to standard error
		}
		// handle the connnection
		go serverCfg.HandleConnection(conn)
	}


}