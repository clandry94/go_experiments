package main

import (
	"bufio"
	"fmt"
	"github.com/clandry94/go_ds/queue"
	workqueue "github.com/clandry94/go_ds/work_queue"
	"math/rand"
	"net"
	"time"
)

var (
	work queue.Queue
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

	timeout := 10 * time.Second
	buffReader := bufio.NewReader(c.Connection)

	c.Connection.SetReadDeadline(time.Now().Add(timeout))
	for {

		// Read tokens delimited by newline
		bytes, err := buffReader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		node := &queue.Node{Value: bytes}
		wq := workqueue.GetInstance()
		wq.Push(node)
		fmt.Printf("%s", bytes)
	}

	fmt.Printf("%d CLOSING\n", c.Id)
	c.Connection.Close()
}

func main() {
	workqueue.New()
	fmt.Println("Setting up")
	ln, err := net.Listen("tcp", ":80")
	if err != nil {
		fmt.Println(err)
	}
	for {
		fmt.Println("listening..")
		newConnection, err := ln.Accept()
		conn := Conn{Id: rand.Int(), Connection: newConnection}
		fmt.Printf("Connection received! ID: %v", conn.Id)
		if err != nil {
			fmt.Println(err)
		}
		go conn.Handle()
	}
}
