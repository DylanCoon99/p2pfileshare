package peer

import (
	//"io"
	//"os"
	"net"
	"log"
	"fmt"
	"time"
	//"bytes"
	"sync"
	"encoding/json"
	"bufio"
	//"strings"
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
	PEER_LIST
	FILE_CHUNK
	PING
	DISCONNECT
	REGISTER
)



type Request struct {
	// type of request sent to the server
	Type RequestType `json:"type"`
	Body []byte      `json:"body"`
	Peer *Peer       `json:"peer"`
	//Conn net.Conn    `json:"conn"`
}



type MetadataSet struct {
	metadata map[string]Metadata
}


type PeerCfg struct {
	MetadataPath string
	DirectoryPath string
	Metadata *MetadataSet
	IP string
	Peers *[]Peer
	// DHT map

}



func InitPeer(serverConn net.Conn) {

	var wg sync.WaitGroup

	//peer cfg
	metadata_path := "/mnt/c/Users/Dylan/My Documents/self_learning/p2p/src/peer/metadata/metadata.json"
	dir_path := "/mnt/c/Users/Dylan/My Documents/self_learning/p2p/src/cmd/peer/dir/"


	metadataSet := CreateMetadataSet()


	peerCfg := PeerCfg {
		MetadataPath: metadata_path,
		DirectoryPath: dir_path,
		Metadata: metadataSet,
		IP: "",
		Peers: nil,
		// add DHT here later
	}


	// register with the server
	peerCfg.Register(serverConn)

	// Generate Metadata

	peerCfg.GenerateMetadata()


	// Listen for incoming peer requests


	// Listen for incoming peer requests
	wg.Add(1)

	go func() {
		defer wg.Done()
		peerCfg.Listen("8000")  // Change this to reflect
	}()

	time.Sleep(5 * time.Second)

	//log.Printf("List of active peers: %v", peerCfg.Peers)

	// Chunk each file, store file chunks indices in the index folder


	// Get array of active peers
	//peers := new([]Peer)
	//peers, err := GetPeers(serverConn)


	// log.Printf("Active Peers:%v ", peers)

	/*

	// Send metadata to all active peers
	for _, peer := range *peers {
		go func() {
			// form a connection with each peer
			// send the metadata to each peer
			conn, err := ConnectToPeer(&peer)

			if err != nil {
				log.Printf("Failed to connect to peer %v: %v", peer, err)
			}

			err = peerCfg.SendMetadata(conn)

			if err != nil {
				log.Printf("Failed to send metadata to peer %v: %v", peer, err)
			}

			Disconnect(conn)

		} ()
	}
	*/

	// Construct all metadata from other peers
	


	// Generate DHT


	wg.Wait()

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




func (cfg *PeerCfg) Register(conn net.Conn) error {

	log.Println("WE ARE IN THE PEER's REGISTER FUNCTION")


	var reqType RequestType
	reqType = 0 // register req on the server

	req := new(Request)
	peer := new(Peer)

	LocalAddr := conn.LocalAddr().(*net.TCPAddr)

	peerIP := fmt.Sprintf("%s:%d", LocalAddr.IP.String(), LocalAddr.Port)


	peer.IP                = peerIP
	peer.Active            = true
	peer.LastServerContact = time.Now()

	// update peerCfg IP
	cfg.IP = peerIP


	log.Printf("HERE IS THE PEER IP: %v", peerIP)



	// build a request
	req.Type = reqType
	req.Body = nil
	req.Peer = peer

	//log.Printf("HERE IS THE CONN: %v", conn)

	// encode the request to bytes
	encodedReq, err := json.Marshal(req)
	encodedReq = append(encodedReq, '\n')

	if err != nil {
		log.Println("Error encoding request to bytes: ", err)
		return err
	}

	_, err = conn.Write(encodedReq)
	if err != nil {
		log.Println("Error writing register request to connection: ", err)
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



func (cfg *PeerCfg) SendMetadata(conn net.Conn) error {

	// this function will send metadata to another peer
	metadataPath := cfg.MetadataPath
	metadataList := new([]Metadata)

	// build the request
	req := new(Request)

	var reqType RequestType
	reqType = 1

	req.Type = reqType
	req.Peer = nil

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




func (cfg *PeerCfg) Listen(port string) {
	// This is where this peer will listen for other peers attempting to connect


	ln, err := net.Listen("tcp", ":" + port)  // returns (Listener, error)

	defer ln.Close()

	if err != nil {
		log.Println(err)
		log.Fatal(err) // prints the error and exits the program
	}
	log.Println("Peer: I am listening for other peers.")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err) // prints the error to standard error
			continue
		}
		// handle the connnection
		go cfg.HandleConnection(conn)
	}

}




