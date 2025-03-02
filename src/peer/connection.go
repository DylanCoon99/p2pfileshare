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



func (peer Peer) HandleRequest(req *Request) {
	// takes a request as input and fulfills that request


	switch t := req.Type; t {
	case CONNECT:
		peer.Connect(req)
	case METADATA_SHARE:
		peer.MetadataShare(req)
	case FILE_CHUNK:
		peer.FileChunk(req)
	case PING:
		peer.Ping(req)
	case DISCONNECT:
		peer.Disconnect(req)
	default:
		// bad request; return 400
		peer.BadRequest(req)
	}


}



// here is the functionality for the peer handling connections

func (peer Peer) HandleConnection(conn net.Conn) {
	defer conn.Close()

	// To be implemented


	return
}


