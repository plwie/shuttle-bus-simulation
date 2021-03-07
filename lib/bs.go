package rs

import (
	"fmt"
	"time"
)

// Bus Struct
type Bus struct {
	availSeats int8
	passOn     int8
	passOff    int8
	currStop   string
	nextStop   string
}

// Busc is the Threading Function
func Busc(name string, path []BusStop, cnt int) {
	//need to declare global count = 0
	pos := cnt
	cnt++
	var len int = len(path)
	var count int = 0
	//create bus struct instance
	busStruct := Bus{
		availSeats: 0,
		passOn:     0,
		passOff:    0,
		currStop:   path[pos].Name,
		nextStop:   path[pos+1].Name,
	}
	//code for bus traveling (busstop to another busstop)
	for {
		if pos < len && name != "test" {
			time.Sleep(time.Second * 1)
			busStruct.currStop = path[pos].Name
			busStruct.nextStop = path[(pos+1)%len].Name
			fmt.Println(count, name, busStruct.currStop, busStruct.nextStop)
			pos++
			count++
		} else {
			pos = 0
		}
	}
}
