package main

import (
	"fmt"
)

// Node object contain a string type data
type Node struct {
	data string
	next *Node
}

// Queue implementation using Node
type Queue struct {
	head *Node
	tail *Node
	size int
}

// Adder add node to queue
type Adder interface{ Add(node Node) }

// Add does not return anything
func (q *Queue) Add(node Node) {
	if q.head == nil {
		fmt.Printf("Add Case 1: %v\n", node)
		q.head = &node
		q.tail = &node
	} else {
		fmt.Printf("Add Case 2: %v\n", node)
		q.tail.next = &node
		q.tail = &node
	}
	q.size++
}

// Popper remove node from queue
type Popper interface{ Pop() *Node }

// Pop return pointer of the removed node
func (q *Queue) Pop() *Node {
	if q.head == nil {
		fmt.Println("Error: queue is empty")
		return nil
	}
	temp := q.head
	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.size--
	return temp
}

// Printer print the value inside the queue
type Printer interface{ printQ() }

// printQ does not return anything
func (q Queue) printQ() {
	fmt.Println("------------------------------------")
	for i := q.head; i != nil; i = i.next {
		fmt.Printf("%v ", i.data)
	}
	fmt.Printf("\nCurrent Queue: %v\nHead: %v\nTail: %v\nSize: %v\n", q, q.head, q.tail, q.size)
	fmt.Println("------------------------------------")
}

// BusStop create a bus stop object
type BusStop struct {
	name             string
	waitingPassenger int
	q                *Queue
}

func main() {
	fmt.Println("Hello")
}
