package main

import (
	"fmt"
	"time"
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
type Printer interface{ printD() }

// printQ does not return anything
func (q Queue) printD() {
	fmt.Printf("Current Queue: ")
	for i := q.head; i != nil; i = i.next {
		fmt.Printf("%v ", i.data)
	}
	fmt.Printf("\nCurrent Queue Info: %v\nHead: %v\nTail: %v\nSize: %v\n", q, q.head, q.tail, q.size)
}

func (bStop BusStop) printD() {
	fmt.Println("------------------------------------")
	fmt.Printf("Bus Stop Name: %v\n", bStop.name)
	fmt.Printf("Waiting People: %v\n", bStop.waitingPassenger)
	bStop.q.printD()
	fmt.Println("------------------------------------")
}

func (p Passenger) printD() {
	fmt.Println("------------------------------------")
	fmt.Printf("Passenger Data: %v\n", p)
	fmt.Println("------------------------------------")
}

type bus struct {
	availSeats int8
	passOn     int8
	passOff    int8
	currStop   string
	nextStop   string
}

// BusStop create a bus stop object
type BusStop struct {
	name             string
	waitingPassenger int
	q                Queue
	passengerDown    int8
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

var hm = BusStop{name: "HM"}
var e12 = BusStop{name: "E12"}
var medical = BusStop{name: "Medical"}
var business = BusStop{name: "Business"}
var it = BusStop{name: "IT"}
var aBuilding = BusStop{name: "aBuilding"}
var bBuilding = BusStop{name: "bBuilding"}
var cBuilding = BusStop{name: "cBuilding"}
var dBuilding = BusStop{name: "dBuilding"}
var eBuilding = BusStop{name: "eBuilding"}

var count int = 0

func main() {
	path := []BusStop{hm, e12, medical, business, it, aBuilding, bBuilding, cBuilding, dBuilding, eBuilding}
	var input int
	fmt.Println("How many bus?")
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		go busc("bus"+fmt.Sprint((i+1)), path)
	}
	busc("test", path)
}

func busc(name string, path []BusStop) {
	pos := count
	count++
	var len int = len(path)
	var count int = 0
	busStruct := bus{
		availSeats: 0,
		passOn:     0,
		passOff:    0,
		currStop:   path[pos].name,
		nextStop:   path[pos+1].name,
	}
	for {
		if pos < len && name != "test" {
			time.Sleep(time.Second * 1)
			busStruct.currStop = path[pos].name
			busStruct.nextStop = path[(pos+1)%len].name
			fmt.Println(count, name, busStruct.currStop, busStruct.nextStop)
			pos++
			count++
		} else {
			pos = 0
		}
	}
}
