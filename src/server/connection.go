package server

import (
	"io"
	"net"
	"log"
	"time"
	"bytes"
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


	log.Printf("This is the buf binary: %b", buf)


	switch str := string(buf); str {
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
		// For some fucked up reason the GET_PEERS req is getting routed to here.
		log.Printf("This is a bad request dummy: %s", str)
	}


	return req
}



func (serverCfg ServerState) HandleRequest(req *Request) {
	// takes a request as input and fulfills that request

	//log.Println(serverCfg.Peers)

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
	buffer := bytes.Buffer{}
	temp := make([]byte, 1024)
	//t := ""


	for {
		// Read the data from the connection
		n, err := conn.Read(temp)
		if err != nil {
			if err == io.EOF {
				break // end of stream; ignore error
			}
			log.Println("Error: ", err)
			return
		}

		buffer.Write(temp[:n]) // append the chuck to the buffer
		

	}

	req := ParseRequest(buffer.Bytes(), conn)

	serverCfg.HandleRequest(req)
	
	buffer.Truncate(0)

	return
}