package peer

import (
	"net"
	"log"
	"fmt"
	"time"
	"encoding/binary"
	"bytes"
)



type serverCfg struct {
	conn net.Conn
	port string
}


type Peer struct {
	// going to have an IP address with port
	// whether the node is active or not

	IP net.IP                    `json:"ip"`
	Active bool                  `json:"active"`
	LastServerContact time.Time  `json:"lastservercontact"`
}



func ConnectToServer(port string) net.Conn {
	//defer conn.Close()

	// attempt to connect to the server via tcp


	conn, err := net.Dial("tcp", "localhost:" + port)
	if err != nil {
		log.Println(err)
		return nil
	}
	log.Println("Peer: Successfully connected to the server!")


	return conn


}

func Register(conn net.Conn) error {

	// Write register to the connection
	data := []byte("REGISTER")
	_, err := conn.Write(data)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	// Read the response from the server

	buffer := make([]byte, 0)

	n, err := conn.Read(buffer)
  	if err != nil {
    	fmt.Println("Error:", err)
    	return err
    }

    // Process and use the data (here, we'll just print it)
    fmt.Printf("Received: %s\n", buffer[:n])

	return nil

}



func GetPeers(conn net.Conn) (*[]Peer, error) {
	// Write Get Peers to the connection


	// Write GET_PEERS to the connection
	req := []byte("GET_PEERS")
	_, err := conn.Write(req)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}




	// Read the response from the server
	buffer := make([]byte, 0)
	_, err = conn.Read(buffer)
  	if err != nil {
    	fmt.Println("Error:", err)
    	return nil, err
    }

    var peers []Peer
    reader := bytes.NewReader(buffer)

    for i := 0; i < len(buffer)/4; i++ {
    	var peer Peer
    	err := binary.Read(reader, binary.BigEndian, &peer)
    	if err != nil {
    		log.Println("Error decoding: ", err)
    		return nil, err
    	}
    	peers = append(peers, peer)

    }




	return &peers, nil


}


func Ping(conn net.Conn) {
	// peers have to periodically ping the server to let them know they are still alive

}


func DisconnectFromServer(conn net.Conn) {
	conn.Close()
	return
}
