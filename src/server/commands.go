package server

import (
	//"encoding/json"
	//"log"
)

func (serverCfg ServerState) Register(req *Request) {
	// add this peer to the peer list and set it's status to active

	// write a response to the connection

	conn := req.Conn
	conn.Write([]byte("This is the server speaking. You sent a register request"))
	peers := serverCfg.Peers

	//i := serverCfg.CurNumPeers

	//*peers[i] = req.Peer
	//serverCfg.CurNumPeers = i + 1

	*peers = append(*peers, *req.Peer)

	/*
	data, err := json.Marshal(req.Peer)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	conn.Write([]byte(data))
	*/
}


func (serverCfg ServerState) GetPeers(req *Request) {
	// return a list of active peers


}


func (serverCfg ServerState) Unregister(req *Request) {
	// Set this peers status to unactive if the node is present in the peer list


}


func (serverCfg ServerState) Disconnect(req *Request) {
	// End this connection with the peer


}


func (serverCfg ServerState) BadRequest(req *Request) {

}

