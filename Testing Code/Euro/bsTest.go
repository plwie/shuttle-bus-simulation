package main

import (
	"fmt"
	"time"
)

type bus struct {
	availSeats int8
	passOn     int8
	passOff    int8
	currStop   string
	nextStop   string
}

var count int = 0

func main() {
	var input int
	path := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	fmt.Println("How many bus?")
	fmt.Scanln(&input)
	for i := 0; i < input; i++ {
		go busc("bus"+fmt.Sprint((i+1)), path)
	}
	busc("test", path)
}

func busc(name string, path []string) {
	pos := count
	count++
	var len int = len(path)
	var count int = 0
	busStruct := bus{
		availSeats: 0,
		passOn:     0,
		passOff:    0,
		currStop:   path[pos],
		nextStop:   path[pos+1],
	}
	for {
		if pos < len && name != "test" {
			time.Sleep(time.Second * 1)
			busStruct.currStop = path[pos]
			busStruct.nextStop = path[(pos+1)%len]
			fmt.Println(count, name, busStruct.currStop, busStruct.nextStop)
			pos++
			count++
		} else {
			pos = 0
		}
	}
}
