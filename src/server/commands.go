package server

import (
	"encoding/json"
	//"bytes"
	"log"
)

func (serverCfg ServerState) Register(req *Request) {
	
	// CORRECT THIS TO NOT ADD A PEER TWICE


	// add this peer to the peer list and set it's status to active

	// write a response to the connection

	conn := req.Conn
	//peers := serverCfg.Peers


	//*peers = append(*peers, *req.Peer)
	*serverCfg.Peers = append(*serverCfg.Peers, *req.Peer)

	//log.Printf("Here is the updated list of peers: %v", serverCfg.Peers)


	conn.Write([]byte("You are successfully registered.\n"))

}


func (serverCfg ServerState) GetPeers(req *Request) {
	// write a list of active peers to a response

	//log.Println(serverCfg.Peers)


	conn := req.Conn

	peers := serverCfg.Peers



	// goal: encode array of Peer struct to json encoding of bytes

	jsonPeers, err := json.Marshal(peers)
	if err != nil {
		log.Printf("Error marshalling JSON Peers: %v", err)
	}


	conn.Write([]byte(string(jsonPeers) + "\n"))


}


func (serverCfg ServerState) Unregister(req *Request) {
	// Set this peers status to unactive if the node is present in the peer list



}


func (serverCfg ServerState) Disconnect(req *Request) {
	// End this connection with the peer


}


func (serverCfg ServerState) BadRequest(req *Request) {

}

