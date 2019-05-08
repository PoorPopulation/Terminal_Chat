package main

import (
	"fmt"
	"net"
	"strings"
)

var conns = make([]net.Conn, 10)

func connHandler(c net.Conn) {
	if c == nil {
		return
	}

	fmt.Println("client remoteAddress: ", c.RemoteAddr().String())
	conns = append(conns, c)
	buf := make([]byte, 4096)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inStr := strings.TrimSpace(string(buf[0:cnt]))
		inputs := strings.Split(inStr, " ")
		switch inputs[0] {
		case "ping":
			c.Write([]byte("pong\n"))
		case "quit":
			c.Close()
			break
		default:
			publish(inputs[0], c)
		}
	}
	fmt.Printf("Connection from %v closed. \n", c.RemoteAddr())
}

func publish(s string, client net.Conn)  {
	for _, conn := range conns {
		if conn != nil && conn != client {
			conn.Write([]byte(s + "\n"))
		}
	}
}


func main() {
	server, err := net.Listen("tcp", ":2333")
	if err != nil {
		fmt.Println("start error", err)
	}

	fmt.Println("server started")

	for {
		conn, err := server.Accept()
		if err != nil{
			fmt.Println(" ", err)
			break
		}
		go connHandler(conn)
	}
}
