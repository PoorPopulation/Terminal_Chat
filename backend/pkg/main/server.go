package main

import (
	"fmt"
	"net"
	"strings"
)

var conns = make([]net.Conn, 10)

func connHandler(conn net.Conn) {
	if conn == nil {
		return
	}
	connRemoteAddress := conn.RemoteAddr().String()
	fmt.Println("client remoteAddress: ", connRemoteAddress)
	conns = append(conns, conn)
	buf := make([]byte, 4096)
	for {
		cnt, err := conn.Read(buf)
		if err != nil || cnt == 0 {
			conn.Close()
			break
		}
		msg := strings.TrimSpace(string(buf[0:cnt]))
		switch msg {
		case "ping":
			sendMsg(conn, "pong")
		case "quit":
			conn.Close()
			break
		default:
			publish(msg, conn)
		}
	}
	fmt.Printf("Connection from %v closed. \n", conn.RemoteAddr())
}

func sendMsg(conn net.Conn, msg string)  {
	if conn == nil || len(msg) == 0 {
		return
	}

	go conn.Write([]byte(msg + "\n"))
}

func publish(s string, client net.Conn)  {
	for _, conn := range conns {
		if conn != nil && conn != client {
			sendMsg(conn, conn.RemoteAddr().String() + ":" + s)
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
