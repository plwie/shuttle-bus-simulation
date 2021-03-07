<<<<<<< Updated upstream:New/busStop.txt
=======
package rs

>>>>>>> Stashed changes:lib/bst.go
import (
	"bufio"
	"fmt"
	"os"
)

// Node object contain a pointer type data
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
	// Init
	reader := bufio.NewReader(os.Stdin)
	var stopNum int
	var stopList []BusStop

	// Creating bus stops
	fmt.Println("Enter the number of bus stops: ")
	fmt.Scanln(&stopNum)
	for i := 1; i <= stopNum; i++ {
		fmt.Printf("Enter the name of bus stop (%v/%v): \n", i, stopNum)
		stopName, _ := reader.ReadString('\n')
		stopList = append(stopList, BusStop{name: stopName})
	}
	fmt.Printf("Bus stop list: %v", stopList)

	// Bus stop
	stop1 := stopList[0]
	stop2 := stopList[1]
	stop1.printD()
	stop2.printD()

	// Passenger
	psg1 := Passenger{src: stop1}
	psg2 := Passenger{src: stop2}
	psg1.printD()
	psg2.printD()

	// Enqueing
	addPSG(&psg1, &stop1)
	stop1.printD()
	stop2.printD()
}
