package bus

import (
	"fmt"
	"time"
)

//Bus Struct
type bus struct {
	availSeats int8
	passOn     int8
	passOff    int8
	currStop   string
	nextStop   string
}

//Threading Function
func busc(name string, path []BusStop) {
	//need to declare global count = 0
	pos := count
	count++
	var len int = len(path)
	var count int = 0
	//create bus struct instance
	busStruct := bus{
		availSeats: 0,
		passOn:     0,
		passOff:    0,
		currStop:   path[pos].name,
		nextStop:   path[pos+1].name,
	}
	//code for bus traveling (busstop to another busstop)
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
