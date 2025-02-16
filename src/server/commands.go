package server

import (
	"encoding/binary"
	"bytes"
	"log"
)

func (serverCfg ServerState) Register(req *Request) {
	
	// CORRECT THIS TO NOT ADD A PEER TWICE


	// add this peer to the peer list and set it's status to active

	// write a response to the connection

	conn := req.Conn
	conn.Write([]byte("This is the server speaking. You sent a register request"))
	//peers := serverCfg.Peers


	//*peers = append(*peers, *req.Peer)
	*serverCfg.Peers = append(*serverCfg.Peers, *req.Peer)

	// print the list of peers for testing purposes
	//log.Println(serverCfg.Peers)

}


func (serverCfg ServerState) GetPeers(req *Request) {
	// write a list of active peers to a response

	//log.Println(serverCfg.Peers)


	conn := req.Conn
	buf := new(bytes.Buffer)

	peers := serverCfg.Peers


	for _, peer := range *peers {

		if peer.Active == true {
			err := binary.Write(buf, binary.BigEndian, peer)
			if err != nil {
				log.Println("Error encoding: ", err) // some values are not fixed size in server.Peer
				return
			}
		}
	}


	log.Printf("This is the GetPeers command: %v", buf.Bytes())

	conn.Write(buf.Bytes())


}


func (serverCfg ServerState) Unregister(req *Request) {
	// Set this peers status to unactive if the node is present in the peer list



}


func (serverCfg ServerState) Disconnect(req *Request) {
	// End this connection with the peer


}


func (serverCfg ServerState) BadRequest(req *Request) {

}

