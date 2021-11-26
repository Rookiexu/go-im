package main

import (
	"../server"
)

func main() {
	s := new(server.TCPServer)
	s.Addr = "localhost:50000"
	s.Start()
}
