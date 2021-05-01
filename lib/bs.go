package rs

import (
	"sync"
)

// Bus Struct
type Bus struct {
	AvailSeats int
	PassOn     int
	CurrStop   string
	NextStop   string
	M          map[string]int
	DistToNext float64
	//Pos "does not refer to (x,y) coordinates of the bus"
	//Since the route is in 1D this Pos variable is to keep the bus in its track
	Pos       int
	FirstTime bool
	State     string
}

var (
	muut sync.Mutex
)

func GetPassngr(path []*BusStop, bus *Bus, calculatedTime *int) {
	var target *BusStop
	for _, v := range path {
		if bus.CurrStop == v.Name {
			target = v
		}
	}
	for bus.AvailSeats > 0 {
		// fmt.Println(target.Q.Size)
		if target.Q.Size != 0 {
			*calculatedTime += target.Q.Head.WaitTime
			bus.M[target.Q.Pop().Destination]++
			bus.PassOn++
			bus.AvailSeats--
		} else {
			break
		}
	}
}

func DropPass(bus *Bus, count *int) {
	*count += bus.M[bus.CurrStop]
	bus.PassOn -= bus.M[bus.CurrStop]
	bus.AvailSeats += bus.M[bus.CurrStop]
	bus.M[bus.CurrStop] = 0
}
