package rs

// Bus Struct
type Bus struct {
	AvailSeats int
	PassOn     int
	CurrStop   string
	NextStop   string
}

func GetPass(m map[string]int, path []*BusStop, i int) {
	len := len(path)
	if path[i%len].Q.Size != 0 {
		m[path[i%len].Q.Pop().Destination]++
	}
}
