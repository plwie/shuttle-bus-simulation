package main

import (
	"fmt"
)

// Node object contain a string type data
type Node struct {
	data *Passenger
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
		q.head = &node
		q.tail = &node
	} else {
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

// Printer print the value inside the struct
type Printer interface{ printS() }

// printQ does not return anything
func (q Queue) printS() {
	for i := q.head; i != nil; i = i.next {
		fmt.Printf("%v ", i.data)
	}
	fmt.Printf("\nCurrent Queue: %v\nHead: %v\nTail: %v\nSize: %v\n", q, q.head, q.tail, q.size)
}

func (bStop BusStop) printS() {
	fmt.Println("------------------------------------")
	fmt.Printf("Bus Stop Data: %v\n", bStop)
	bStop.q.printS()
	fmt.Println("------------------------------------")
}

func (p Passenger) printS() {
	fmt.Println("------------------------------------")
	fmt.Printf("Passenger Data: %v\n", p)
	fmt.Println("------------------------------------")
}

// BusStop create a bus stop object
type BusStop struct {
	name             string
	waitingPassenger int
	q                Queue
}

// Passenger create a passenger object
type Passenger struct {
	src BusStop
	dst BusStop
}

func addPSG(psg *Passenger, bStop *BusStop) {
	tg := Node{data: psg}
	bStop.q.Add(tg)
	bStop.waitingPassenger++
}

func main() {
	fmt.Println("Hello")

	// Bus stop
	stop1 := BusStop{name: "MIT Stop"}
	stop2 := BusStop{name: "CMKL Stop"}
	stop1.printS()
	stop2.printS()

	// Passenger
	psg1 := Passenger{src: stop1}
	psg2 := Passenger{src: stop2}
	psg1.printS()
	psg2.printS()

	// Enqueing
	addPSG(&psg1, &stop1)
	stop1.printS()
	stop2.printS()
}
