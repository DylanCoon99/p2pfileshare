package server



import (
	"log"
	"time"
	"encoding/json"
)


func (serverCfg *ServerState) ShareAllPeers() error {

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

		if err != nil {
			log.Printf("Error connecting to peer %v: %v", peer, err)
			return err
		}

		log.Printf("Connected to peer: %v\n", peer)

		// build a req
		req := new(Request)
		req.Type = reqType
		req.Body = buf
		req.Peer = nil

		encodedReq, err := json.Marshal(req)

		if err != nil {
			log.Printf("Error encoding peer list request: %v", err)
			return err
		}

		conn.Write(encodedReq)
		time.Sleep(5 * time.Second)


		conn.Close()

	}

	return nil

}