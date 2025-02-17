package server

import (
	//"io"
	"net"
	"log"
	"time"
	//"bytes"
	"bufio"
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


	//log.Printf("This is the buf binary: %b", buf)


	switch str := string(buf); str {
	case "REGISTER\n":
		req.Type = REGISTER
	case "UNREGISTER\n":
		req.Type = UNREGISTER
	case "GET_PEERS\n":
		req.Type = GET_PEERS
	case "DISCONNECT\n":
		req.Type = DISCONNECT
	default:
		// bad request
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

	log.Println("Server: Handling the connection now!")
	// Read data from the connection
	// Determine the type of request
	reader := bufio.NewReader(conn)
	
	requestLine, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading request:", err)
		return
	}

	req := ParseRequest([]byte(requestLine), conn)

	serverCfg.HandleRequest(req)
	

	return
}