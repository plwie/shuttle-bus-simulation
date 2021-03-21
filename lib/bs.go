package rs

// Bus Struct
type Bus struct {
	AvailSeats int
	PassOn     int
	CurrStop   string
	NextStop   string
	M          map[string]int
}

func GetPass(m map[string]int, path []*BusStop, bus *Bus) {
	len := len(path)
	for i := 0; i < bus.AvailSeats; i++ {
		if path[i%len].Q.Size != 0 {
			m[path[i%len].Q.Pop().Destination]++
			bus.PassOn++
			bus.AvailSeats--
		}
	}

}

func DropPass(m map[string]int, bus *Bus) {
	bus.PassOn -= m[bus.CurrStop]
	bus.AvailSeats += m[bus.CurrStop]
}
