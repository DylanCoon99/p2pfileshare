package peer

import (
	//"io"
	"net"
	"log"
	"fmt"
	"time"
	//"bytes"
	"encoding/json"
	"bufio"
	"strings"
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
	data := []byte("REGISTER\n")
	_, err := conn.Write(data)
	if err != nil {
		log.Println("Error: ", err)
		return err
	}
	// Read the response from the server

	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Error reading response: ", err)
		return err
	}

	fmt.Printf("Received %s\n", response)

	return nil

}



func GetPeers(conn net.Conn) (*[]Peer, error) {
	// Write Get Peers to the connection


	// Write GET_PEERS to the connection
	req := []byte("GET_PEERS\n")
	_, err := conn.Write(req)
	if err != nil {
		log.Println("Error: ", err)
		return nil, err
	}


	// Read the response from the server
	reader := bufio.NewReader(conn)
	response, err := reader.ReadString('\n')
	response = strings.TrimSpace(response)
	if err != nil {
		log.Println("Error reading response: ", err)
		return nil, err
	}

	//log.Printf("Here is the list of peers: %v", response)

    var peers []Peer

    err = json.Unmarshal([]byte(response), &peers)


    if err != nil {
    	log.Printf("Error decoding peers: %v", err)
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



func (peer Peer) Listen(port string) {
	// This is where this peer will listen for other peers attempting to connect


	ln, err := net.Listen("tcp", ":" + port)  // returns (Listener, error)


	if err != nil {
		log.Fatal(err) // prints the error and exits the program
	}
	log.Println("Peer: I am listening for other peers.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err) // prints the error to standard error
		}
		// handle the connnection
		go peer.HandleConnection(conn)
	}

}


