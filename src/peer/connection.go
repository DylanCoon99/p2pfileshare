package peer


import (
	//"io"
	"fmt"
	"net"
	"log"
	"time"
	//"bytes"
	//"bufio"
	"encoding/json"
)



// types of requests that a peer needs to handle
// (1) Another peer can request a chunk
// (2) Server may ping to see if peer is still active
// (3) Another peer can send its DHT data


func ParseRequest(buf []byte, conn net.Conn) *Request {
	// return a Request struct containing info in the request


	req := new(Request)
	peer := new(Peer)

	// decode the bytes into a req
	err := json.Unmarshal(buf, req)


	if err != nil {
		log.Printf("Error parsing request: %v", err)
		return nil
	}


	localAddr := conn.LocalAddr().(*net.TCPAddr)

	peer.IP                = fmt.Sprintf("%s:%d", localAddr.IP.String(), localAddr.Port)
	peer.Active            = true
	peer.LastServerContact = time.Now()

	req.Peer = peer


	return req

}



func (cfg *PeerCfg) HandleRequest(req *Request) {
	// takes a request as input and fulfills that request


	switch t := req.Type; t {
	case CONNECT:
		cfg.Connect(req)
	case METADATA_SHARE:
		cfg.MetadataShare(req)
	case PEER_LIST:
		log.Println("CONFIRMING THIS LINE IS WORKING")
		cfg.PeerList(req)
	case FILE_CHUNK:
		cfg.FileChunk(req)
	case PING:
		cfg.Ping(req)
	case DISCONNECT:
		cfg.Disconnect(req)
	default:
		// bad request; return 400
		cfg.BadRequest(req)
	}


}



// here is the functionality for the peer handling connections

func (cfg *PeerCfg) HandleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("Peer: Handling the connection now!")
	// Read data from the connection
	// Determine the type of request
	req := new(Request)
	buf := new([]byte) // this returns actual buf not ptr

	_, err := conn.Read(*buf)

	if err != nil {
		log.Printf("Failed to read request from connection: %v", err)
	}



	req = ParseRequest(*buf, conn)

	log.Println("HERE IS A REQUEST: %v", req)

	cfg.HandleRequest(req)

	return
}


