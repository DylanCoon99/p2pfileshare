package server

import (
	"encoding/json"
	//"bytes"
	"log"
	"time"
	"net"
)

func (serverCfg ServerState) Register(req *Request, conn net.Conn) {
	
	// CORRECT THIS TO NOT ADD A PEER TWICE

	// add this peer to the peer list and set it's status to active

	// write a response to the connection

	//peers := serverCfg.Peers
	defer conn.Close()

	//*peers = append(*peers, *req.Peer)
	*serverCfg.Peers = append(*serverCfg.Peers, *req.Peer)
	serverCfg.CurNumPeers += 1

	log.Printf("Here is the updated list of peers: %v", serverCfg.Peers)


	conn.Write([]byte("You are successfully registered.\n"))


	time.Sleep(5 * time.Second)

	err := serverCfg.ShareAllPeers()

	if err != nil {
		log.Printf("Error sharing updated peer list: %v", err)
	}

}


func (serverCfg ServerState) GetPeers(req *Request) {
	// write a list of active peers to a response

	//log.Println(serverCfg.Peers)


	//conn := req.Conn

	peers := serverCfg.Peers



	// goal: encode array of Peer struct to json encoding of bytes

	jsonPeers, err := json.Marshal(peers)
	if err != nil {
		log.Printf("Error marshalling JSON Peers: %v", err)
	}


	//conn.Write([]byte(string(jsonPeers) + "\n"))

	log.Println(jsonPeers)
}



func (serverCfg ServerState) BadRequest(req *Request, conn net.Conn) {

}

