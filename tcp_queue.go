package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	//"time"
)

type Node struct {
	conn net.Conn
}

type Queue struct {
	nodes []*Node
	head  int
	tail  int
	count int
}

func handleConnection(conn net.Conn, threadID int) {

	//	timeout := 10 * time.Second
	buffReader := bufio.NewReader(conn)

	for {
		fmt.Printf("%d READING\n", threadID)
		//conn.SetReadDeadline(time.Now().Add(timeout))

		// Read tokens delimited by newline
		bytes, err := buffReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s", bytes)
	}

	fmt.Printf("%d CLOSING\n", threadID)
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
		threadID := rand.Int()
		fmt.Printf("Creating thread ID: %d\n", threadID)
		go handleConnection(conn, threadID)
	}
}
