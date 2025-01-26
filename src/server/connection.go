package server

import (
	"net"
	"log"
)

// here is the functionality for the server handling connections


func HandleConnection(conn net.Conn) {
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

	}

	// Write data to the connection

	return
}