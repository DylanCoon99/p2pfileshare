package peer

import (
	"log"
	"encoding/json"
)

// These are cmds that are called by the peer request handler

func (cfg *PeerCfg) Connect(req *Request) {


}



func (cfg *PeerCfg) MetadataShare(req *Request) {

	// this handles a metadata share request (when another node is sharing metadata for it's files)


	// decode the bytes into a list of metadata
	metadata := new([]Metadata)
	err := json.Unmarshal(req.Body, metadata)

	if err != nil {
		log.Printf("Error decoding metadata: %v", err)
		return
	}

	// we need to be able to read current metadata for this peer and then add the new metadata without duplicates
	
	err = cfg.ConstructMetadata(*metadata)

	if err != nil {
		log.Printf("Error contructing metadata")
	}



}


func (cfg *PeerCfg) FileChunk(req *Request) {


}


func (cfg *PeerCfg) Ping(req *Request) {


}


func (cfg *PeerCfg) Disconnect(req *Request) {


}


func (cfg *PeerCfg) BadRequest(req *Request) {


}