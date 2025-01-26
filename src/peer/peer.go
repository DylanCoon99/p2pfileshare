package peer

import (
	"net"
	"log"
)



type serverCfg struct {
	conn net.Conn
	port string
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

func Register(conn net.Conn) {

	// Write register to the connection
	data := []byte("Hi, I'd like to register")
	_, err := conn.Write(data)
	if err != nil {
		log.Println("Error: ", err)
	}
	

}



func DisconnectFromServer(conn net.Conn) {
	conn.Close()
	return
}
