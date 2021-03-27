package rs

// Bus Struct
type Bus struct {
	AvailSeats int
	PassOn     int
	CurrStop   string
	NextStop   string
	M          map[string]int
	DistToNext float64
	AtStop     bool
	Pos        int
	FirstTime  bool
	Status     string
}

func GetPassngr(path []*BusStop, bus *Bus, count *int) {
	var target *BusStop
	for _, v := range path {
		if bus.CurrStop == v.Name {
			target = v
		}
	}
	for bus.AvailSeats != 0 {
		// fmt.Println(target.Q.Size)
		if target.Q.Size != 0 {
			bus.M[target.Q.Pop().Destination]++
			bus.PassOn++
			*count++
			bus.AvailSeats--
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
