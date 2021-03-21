package rs

// Bus Struct
type Bus struct {
	AvailSeats int
	PassOn     int
	CurrStop   string
	NextStop   string
	M          map[string]int
}

func GetPass(path []*BusStop, bus *Bus) {
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
		}
	}

}

func DropPass(m map[string]int, bus *Bus) {
	bus.PassOn -= m[bus.CurrStop]
	bus.AvailSeats += m[bus.CurrStop]
}
