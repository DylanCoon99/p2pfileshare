package server

import (
	//"io"
	"fmt"
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

	peer.IP                = fmt.Sprintf("%s:%d", localAddr.IP.String(), localAddr.Port)
	peer.Active            = true
	peer.LastServerContact = time.Now()

	req.Body = buf
	req.Peer = peer
	req.Conn = conn


	//log.Printf("This is the buf binary: %b", buf)


	switch str := string(buf); str {
	case "REGISTER\n":
		req.Type = REGISTER
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