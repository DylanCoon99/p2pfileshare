package peer

import (
	//"io"
	//"os"
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



type MetadataSet struct {
	metadata map[string]Metadata
}


type PeerCfg struct {
	MetadataPath string
	DirectoryPath string
	Metadata *MetadataSet
	// DHT map

}



func InitPeer() {

	//peer cfg
	metadata_path := "/mnt/c/Users/Dylan/My Documents/self_learning/p2p/src/peer/metadata/metadata.json"
	dir_path := "/mnt/c/Users/Dylan/My Documents/self_learning/p2p/src/cmd/peer/dir/"


	metadataSet := CreateMetadataSet()


	peerCfg := PeerCfg {
		MetadataPath: metadata_path,
		DirectoryPath: dir_path,
		Metadata: metadataSet,
		// add DHT here later
	}


	// register with the server


	// Generate Metadata

	peerCfg.GenerateMetadata()

	// Chunk each file, store file chunks indices in the index folder


	// Get list of active peers


	// Send metadata to all active peers


	// Construct all metadata from other peers


	// Generate DHT


	// Listen for incoming peer requests


}





func ConnectToServer(port string) net.Conn {

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



func (cfg *PeerCfg) SendMetadata(conn net.Conn) (error) {

	// this function will send metadata to another peer
	metadataPath := cfg.MetadataPath
	metadataList := new([]Metadata)

	// build the request
	req := new(Request)

	var reqType RequestType
	reqType = 1

	req.Type = reqType
	req.Peer = nil
	req.Conn = conn

	// read metadata from the metadata.json
	metadataList, err := ExtractMetadata(metadataPath)

	if err != nil {
		log.Printf("Failed to extract metadata: %v", err)
	}

	byteValue, err := json.Marshal(metadataList)

	if err != nil {
		log.Printf("Failed to encode metadata list: %v", err)
	}

	req.Body = byteValue

	encodedReq, err := json.Marshal(req)

	if err != nil {
		log.Printf("Failed to share metadata request: %v", err)
	}

	conn.Write(encodedReq)

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




