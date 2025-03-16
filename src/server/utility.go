package server



import (
	"log"
	"encoding/json"
)


func (serverCfg ServerState) ShareAllPeers() error {

	// iterates over all peers and sends all active peers to each peer

	peers := serverCfg.Peers

	var reqType RequestType
	reqType = 2

	buf, err := json.Marshal(peers)

	if err != nil {
		log.Printf("Error encoding peer list: %v", err)
		return err
	}

	for _, peer := range *peers {

		if !peer.Active {
			continue
		}

		conn, err := ConnectToPeer(&peer)

		// build a req
		req := new(Request)
		req.Type = reqType
		req.Body = buf
		req.Peer = nil
		req.Conn = conn

		encodedReq, err := json.Marshal(req)

		if err != nil {
			log.Printf("Error encoding peer list request: %v", err)
			return err
		}

		conn.Write(encodedReq)

		conn.Close()

	}

	return nil

}