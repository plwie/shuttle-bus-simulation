package rs

import "fmt"

// Queue implementation using Passenger node
type Queue struct {
	Head *Passenger
	Tail *Passenger
	Size int
}

// Adder add node to queue
type Adder interface{ Add(node Passenger) }

// Add does not return anything
func (q *Queue) Add(node Passenger) {
	if q.Head == nil {
		q.Head = &node
		q.Tail = &node
	} else {
		q.Tail.Next = &node
		q.Tail = &node
	}
	q.Size++
}

// Popper remove node from queue
type Popper interface{ Pop() *Passenger }

// Pop return pointer of the removed node
func (q *Queue) Pop() *Passenger {
	if q.Head == nil {
		fmt.Println("Error: queue is empty")
		return nil
	}
	temp := q.Head
	if q.Size == 1 {
		q.Head = nil
		q.Tail = nil
	} else {
		q.Head = q.Head.Next
	}
	q.Size--
	return temp
}

// Printer print the value inside the struct
type Printer interface{ printD() }

// printQ does not return anything
func (q Queue) printD() {
	fmt.Printf("Current Queue: ")
	for i := q.Head; i != nil; i = i.Next {
		fmt.Printf("%v ", i)
	}
	fmt.Printf("\nCurrent Queue Info: %v\nHead: %v\nTail: %v\nSize: %v\n", q, q.Head, q.Tail, q.Size)
}

func (bStop BusStop) printD() {
	fmt.Println("------------------------------------")
	fmt.Printf("Bus Stop Name: %v\n", bStop.Name)
	fmt.Printf("Waiting People: %v\n", bStop.Q.Size)
	bStop.Q.printD()
	fmt.Println("------------------------------------")
}

// BusStop create a bus stop object
type BusStop struct {
	Name      string
	Q         Queue
	TimeTaken int
}
