package main

import (
	"fmt"
	"github.com/clandry94/go_ds/queue"
	"math/rand"
	"net"
)

type Request struct {
	Cookie Cookie
}

type Cookie struct {
	Id int
}

func printQueue(workQueue queue.Queue) {
	fmt.Printf("==== Queue Size: %v ====\n", workQueue.Size)
	//p := workQueue.Front
	//for p != nil {
	//	id := p.Value.Id
	//	fmt.Println(id)
	//	p = p.Next
	//}
	fmt.Println()
}

func worker(work *queue.Queue) {
	for {
		if work.Size > 0 {
			job := work.Front.Value
			job.Handle()
			fmt.Println(job)
			work.Pop()
		}
	}
}

func main() {
	work := queue.New()
	fmt.Println("Setting up")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
	}

	go worker(&work)

	for {
		fmt.Println("listening..")
		newConnection, err := ln.Accept()
		conn := queue.Conn{Id: rand.Int(), Connection: newConnection}
		fmt.Printf("Connection received! ID: %v\n", conn.Id)
		if err != nil {
			fmt.Println(err)
		}
		node := queue.Node{Value: conn}
		work.Push(node)
		printQueue(work)
	}
}
