package main

import (
	"fmt"
	"math/rand"
	"time"
)

type bus struct {
	nameBus     string
	avail_seats int8
	pass_on     int8
	pass_off    int8
	curr_stop   string
	next_stop   string
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
	pos := rand.Intn(3)
	var count int = 0
	busStruct := bus{
		avail_seats: 0,
		pass_on:     0,
		pass_off:    0,
		curr_stop:   path[pos],
		next_stop:   path[pos+1],
	}
	for {
		if pos < 10-1 && name != "test" {
			time.Sleep(time.Second * 2)
			busStruct.curr_stop = path[pos]
			busStruct.next_stop = path[pos+1]
			fmt.Println(count, name, busStruct.curr_stop, busStruct.next_stop)
			pos += 1
			count += 1
		} else {
			pos = 0
		}
	}
}
