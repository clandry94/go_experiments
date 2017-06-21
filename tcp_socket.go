package main

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func handleConnection(conn net.Conn) {
	fmt.Println("Handling conn")

	timeout := 10 * time.Second
	buffReader := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(timeout))

		// Read tokens delimited by newline
		bytes, err := buffReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", bytes)
	}

	fmt.Println("Closing..")
	conn.Close()
}

func main() {

	fmt.Println("Setting up")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("listening..")
		conn, err := ln.Accept()
		fmt.Println("Connection received!")
		if err != nil {
			fmt.Println(err)
		}
		handleConnection(conn)
	}
}
