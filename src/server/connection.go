package server

import (
	//"io"
	//"fmt"
	"net"
	"log"
	//"time"
	//"bytes"
	"bufio"
	"encoding/json"
)




func ParseRequest(buf []byte) *Request {
	// return a Request struct containing info in the request


	req := new(Request)

	err := json.Unmarshal(buf, req)

	log.Printf("HERE IS THE REQUEST: %v", req)

	if err != nil {
		log.Printf("Error decoding request: %v", err)
	}


	return req
}



func (serverCfg ServerState) HandleRequest(req *Request, conn net.Conn) {
	// takes a request as input and fulfills that request

	switch t := req.Type; t {
	case REGISTER:
		serverCfg.Register(req, conn)
	default:
		// bad request; return 400
		serverCfg.BadRequest(req, conn)
	}


}




// here is the functionality for the server handling connections

func (serverCfg *ServerState) HandleConnection(conn net.Conn) {
	defer conn.Close()

	log.Println("Server: Handling the connection now!")
	// Read data from the connection
	// Determine the type of request
	reader := bufio.NewReader(conn)
	
	requestLine, err := reader.ReadString('\n')

	log.Println(requestLine)

	if err != nil {
		log.Println("Error reading request:", err)
		return
	}

	req := ParseRequest([]byte(requestLine))

	serverCfg.HandleRequest(req, conn)
	

	return
}



func ConnectToPeer(peer *Peer) (net.Conn, error) {

	// attempt to connect to the server via tcp

	conn, err := net.Dial("tcp", peer.IP)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Server: Successfully connected to peer!")


	return conn, nil

}