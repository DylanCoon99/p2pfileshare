package main

// This is an entry point for the server

import (
	"github.com/DylanCoon99/p2pfileshare/src/server"
)





func main() {

	const port = "8080"
	server.StartServer(port)

}