package main

import (
	"./core"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

var conns = make([]net.Conn, 10)

var crcCode = []byte("BABA")

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

		if cnt < 110 {
			conn.Close()
			break
		}
		fmt.Println(cnt)

		proto := core.PpProto{}
		length := binary.BigEndian.Uint32(buf[4:12])

		proto.WriteCode(buf[0:4]).WriteLength(length).WritePpType(buf[12:14][0]).WriteToken(string(buf[14:78])).WriteUUID(string(buf[78:110])).WriteBody(string(buf[110:cnt]))

		if !bytes.Equal(crcCode, proto.Code) {
			conn.Close()
			break
		}

		msg := proto.Body
		fmt.Println(msg)
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

func sendMsg(conn net.Conn, msg string) {
	if conn == nil || len(msg) == 0 {
		return
	}

	go conn.Write([]byte(msg + "\n"))
}

func publish(s string, client net.Conn) {
	for _, conn := range conns {
		if conn != nil && conn != client {
			sendMsg(conn, conn.RemoteAddr().String()+":"+s)
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
		if err != nil {
			fmt.Println(" ", err)
			break
		}
		go connHandler(conn)
	}
}
