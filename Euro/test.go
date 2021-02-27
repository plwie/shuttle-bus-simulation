package main

import (
	"fmt"
	"math/rand"
	"time"
)

type bus struct {
	pass_off    int8
	availSeats int8
	passOn     int8
	passOff    int8
	currStop   string
	nextStop   string
}

func main() {
	var input int
	fmt.Println("How many bus?")
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		go busc("bus" + fmt.Sprint((i + 1)))
	}
	busc("test")
}

func busc(name string) {
	path := [10]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	pos := rand.Intn(9)
	var count int = 0
	busStruct := bus{
		availSeats: 0,
		passOn:     0,
		passOff:    0,
		currStop:   path[pos],
		nextStop:   path[pos+1],
	}
	for {
		if pos < len(path) && name != "test" {
			time.Sleep(time.Second * 1)
			busStruct.currStop = path[pos]
			busStruct.nextStop = path[(pos+1)%len(path)]
			fmt.Println(count, name, busStruct.currStop, busStruct.nextStop)
			pos++
			count++
		} else {
			pos = 0
		}
	}
}
