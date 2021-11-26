package server

import (
	"fmt"
	"net"
)

type TCPConn struct {
	tcp      net.Conn
	writeMsg chan string
	readMsg  chan string
}

func NewClientTCPConn(addr string) (*TCPConn, error) {
	tcp, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	conn := new(TCPConn)
	conn.tcp = tcp
	conn.writeMsg = make(chan string, 100)
	conn.readMsg = make(chan string, 100)

	go conn.read()
	go conn.write()
	return conn, nil
}

func NewTCPConn(tcp net.Conn) *TCPConn {
	conn := new(TCPConn)
	conn.tcp = tcp
	conn.writeMsg = make(chan string, 100)
	conn.readMsg = make(chan string, 100)

	go conn.read()
	go conn.write()
	return conn
}

func (conn *TCPConn) read() {
	for {
		buf := make([]byte, 512)
		read, err := conn.tcp.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		conn.readMsg <- string(buf[:read])
	}
}

func (conn *TCPConn) write() {
	for {
		for writeMsg := range conn.writeMsg {
			_, err := conn.tcp.Write([]byte(" says: " + writeMsg))
			if err != nil {
				fmt.Println("Error write", err.Error())
				break
			}
		}
	}
}

func (conn *TCPConn) SendMsg(msg string) {
	conn.writeMsg <- msg
}
