package server

import (
	"fmt"
	"net"
)

type TCPServer struct {
	Addr     string
	listener net.Listener
}

func (server *TCPServer) init() {
	fmt.Println("Starting the server ...")
	// 创建 listener
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		fmt.Println("Error listening", err.Error())
		return //终止程序
	}
	server.listener = listener
}

func (server *TCPServer) run() {
	// 监听并接受来自客户端的连接
	for {
		conn, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting", err.Error())
			return // 终止程序
		}

		tcpConn := NewTCPConn(conn)
		go handler(tcpConn)
	}
}

func handler(conn *TCPConn) {
	for {
		for s := range conn.readMsg {
			fmt.Printf("Received data: %v\n", s)
		}
	}
}

func (server *TCPServer) Start() {
	server.init()
	server.run()
}
