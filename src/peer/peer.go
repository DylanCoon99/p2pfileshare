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


const CHUNK_SIZE uint64 = 16384 //16KB Chunk size

type Peer struct {
	// going to have an IP address with port
	// whether the node is active or not

	IP string                    `json:"ip"`  // "ip:port" in string format
	Active bool                  `json:"active"`
	LastServerContact time.Time  `json:"lastservercontact"`
}



type RequestType int

const (
	CONNECT RequestType = iota
	METADATA_SHARE
	DHT_INFO
	FILE_CHUNK
	PING
	DISCONNECT
)



type Request struct {
	// type of request sent to the server
	Type RequestType
	Body []byte
	Peer *Peer
	Conn net.Conn
}




func InitPeer() {


	// register with the server


	// Generate Metadata

	file1_path := "/mnt/c/Users/Dylan/My Documents/self_learning/p2p/src/cmd/peer/dir/file1_test.txt"

	metadata_path := "/mnt/c/Users/Dylan/My Documents/self_learning/p2p/src/peer/metadata/metadata.json"

	GenerateMetadata(file1_path, metadata_path)

	// Chunk each file, store file chunks indices in the index folder


	// Get list of active peers


	// Send metadata to all active peers


	// Update DHT


	// Send DHT to all active peers


	// Listen for incoming peer requests


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



func SendMetadata(conn net.Conn) (error) {

	// this function will send metadata to another peer


	// read metadata from the metadata.json


	// encode the metadata to bytes


	// send data over the conn


	return nil
}





func Ping(conn net.Conn) {
	// peers have to periodically ping the server to let them know they are still alive

}


func Disconnect(conn net.Conn) {
	conn.Close()
	return
}


func ConnectToPeer(peer *Peer) (net.Conn, error) {

	// attempt to connect to the server via tcp

	conn, err := net.Dial("tcp", peer.IP)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("Peer: Successfully connected to peer!")


	return conn, nil

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




