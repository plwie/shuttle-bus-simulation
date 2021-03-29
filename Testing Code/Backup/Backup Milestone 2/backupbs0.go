package rs

// Bus Struct
type Bus struct {
	AvailSeats int
	PassOn     int
	CurrStop   string
	NextStop   string
	M          map[string]int
}

func GetPass(path []*BusStop, bus *Bus, countP *int) {
	var target *BusStop
	for _, v := range path {
		if bus.CurrStop == v.Name {
			target = v
		}
	}
	for bus.AvailSeats != 0 {
		if target.Q.Size != 0 {
			bus.M[target.Q.Pop().Destination]++
			bus.PassOn++
			bus.AvailSeats--
			(*countP)++
		} else {
			break
		}
	}

}

func DropPass(bus *Bus) {
	bus.PassOn -= bus.M[bus.CurrStop]
	bus.AvailSeats += bus.M[bus.CurrStop]
	bus.M[bus.CurrStop] = 0
}
