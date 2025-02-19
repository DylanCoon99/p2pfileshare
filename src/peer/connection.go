package peer


import (
	//"io"
	"net"
	"log"
	"time"
	//"bytes"
	//"bufio"
)



// types of requests that a peer needs to handle
// (1) Another peer can request a chunk
// (2) Server may ping to see if peer is still active
// (3) Another peer can send its DHT data


func ParseRequest(buf []byte, conn net.Conn) *Request {
	// return a Request struct containing info in the request


	req := new(Request)
	peer := new(Peer)

	localAddr := conn.LocalAddr().(*net.TCPAddr)

	peer.IP = localAddr.IP
	peer.Active = true
	peer.LastServerContact = time.Now()

	req.Body = buf
	req.Peer = peer
	req.Conn = conn


	//log.Printf("This is the buf binary: %b", buf)


	switch str := string(buf); str {
	case "CONNECT\n":
		req.Type = CONNECT
	case "DHT_INFO\n":
		req.Type = DHT_INFO
	case "FILE_CHUNK\n":
		req.Type = FILE_CHUNK
	case "PING\n":
		req.Type = PING
	case "DISCONNECT\n":
		req.Type = DISCONNECT
	default:
		// bad request
		log.Printf("This is a bad request dummy: %s", str)
	}


	return req
}



func (peer Peer) HandleRequest(req *Request) {
	// takes a request as input and fulfills that request


	switch t := req.Type; t {
	case CONNECT:
		peer.Connect(req)
	case DHT_INFO:
		peer.DHTInfo(req)
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



	return
}


