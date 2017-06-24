package main

import (
	"bufio"
	"fmt"
	"github.com/clandry94/go_ds/queue"
	"math/rand"
	"net"
	"time"
)

type Request struct {
	Cookie Cookie
}

type Cookie struct {
	Id int
}

type Conn struct {
	Connection net.Conn
	Id         int
}

func (c *Conn) Handle() {

	timeout := 5 * time.Second
	buffReader := bufio.NewReader(c.Connection)

	c.Connection.SetReadDeadline(time.Now().Add(timeout))
	for {

		// Read tokens delimited by newline
		bytes, err := buffReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s", bytes)
	}

	fmt.Printf("%d CLOSING\n", c.Id)
	c.Connection.Close()
}

func printQueue(workQueue queue.Queue) {
	fmt.Printf("==== Queue Size: %v ====\n", workQueue.Size)
	p := workQueue.Front
	for p != nil {
		id := p.Value.(Conn).Id
		fmt.Println(id)
		p = p.Next
	}
	fmt.Println()
}

func main() {
	work := queue.New()
	fmt.Println("Setting up")
	ln, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println(err)
	}

	for {
		fmt.Println("listening..")
		newConnection, err := ln.Accept()
		conn := Conn{Id: rand.Int(), Connection: newConnection}
		fmt.Printf("Connection received! ID: %v\n", conn.Id)
		if err != nil {
			fmt.Println(err)
		}
		node := queue.Node{Value: conn}
		work.Push(node)
		printQueue(work)
	}
}
