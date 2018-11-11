package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func handler(conn net.Conn) {
	defer conn.Close()
	conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	request := make([]byte, 128)
	for {
		readLen, err := conn.Read(request)
		if err != nil {
			fmt.Printf("Client state change:\n%s\n", err.Error())
			break
		}
		if readLen == 0 {
			break
		}
		fmt.Printf("%s\n", string(request[:readLen]))
		conn.Write([]byte("Hello\n"))
		conn.Close()
	}
}

func main() {
	port := ":8080"
	fmt.Printf("Started in port %s\n", port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handler(conn)
	}
}
