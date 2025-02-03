package server

import (
	"net"
	"log"
	"time"
)




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


	str := string(buf)

	switch str {
	case "REGISTER":
		req.Type = REGISTER
	case "UNREGISTER":
		req.Type = UNREGISTER
	case "GET_PEERS":
		req.Type = GET_PEERS
	case "DISCONNECT":
		req.Type = DISCONNECT
	default:
		// bad request
	}


	return req
}



func (serverCfg ServerState) HandleRequest(req *Request) {
	// takes a request as input and fulfills that request

	switch t := req.Type; t {
	case REGISTER:
		serverCfg.Register(req)
	case UNREGISTER:
		serverCfg.Unregister(req)
	case GET_PEERS:
		serverCfg.GetPeers(req)
	case DISCONNECT:
		serverCfg.Disconnect(req)
	default:
		// bad request; return 400
		serverCfg.BadRequest(req)
	}


}




// here is the functionality for the server handling connections

func (serverCfg ServerState) HandleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("Server: I am handling the connection now!")
	// Read data from the connection
	// Determine the type of request
	buffer := make([]byte, 1024)
	//t := ""

	for {
		// Read the data from the connection
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error: ", err)
			return
		}

		log.Printf("Received: %s\n", buffer[:n])
		req := ParseRequest(buffer, conn)
		serverCfg.HandleRequest(req)

	}

	// Write data to the connection

	return
}