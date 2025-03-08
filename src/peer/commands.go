package peer

import (
	"log"
	"encoding/json"
)

// These are cmds that are called by the peer request handler

func (peer Peer) Connect(req *Request) {


}



func (peer Peer) MetadataShare(req *Request) {

	// this handles a metadata share request (when another node is sharing metadata for it's files)


	// decode the bytes into a list of metadata
	metadata := new([]Metadata)
	err := json.Unmarshal(req.Body, metadata)

	if err != nil {
		log.Printf("Error decoding metadata: %v", err)
		return
	}

	// we need to be able to read current metadata for this peer and then add the new metadata without duplicates
	
}


func (peer Peer) FileChunk(req *Request) {


}


func (peer Peer) Ping(req *Request) {


}


func (peer Peer) Disconnect(req *Request) {


}


func (peer Peer) BadRequest(req *Request) {


}