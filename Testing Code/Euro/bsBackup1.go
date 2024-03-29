package rs

// Bus Struct
type Bus struct {
	availSeats int
	passOn     int
	currStop   string
	nextStop   string
}

/*
// Busc is the Threading Function
func Busc(name string, path []*BusStop) {
	//need to declare global count = 0
	m := make(map[string]int)
	pos := countPos
	countPos++
	var len int = len(path)
	var count int = 0
	//create bus struct instance
	busStruct := Bus{
		availSeats: 30,
		passOn:     0,
		currStop:   *&path[pos].Name,
		nextStop:   *&path[pos+1].Name,
	}
	for i := 0; i < 10; i++ {
		m[path[i].Name] = 0
	}
	//code for bus traveling (busstop to another busstop)
	for {
		if pos < len && name != "test" {
			// time.Sleep(time.Second * 1)
			busStruct.currStop = *&path[pos].Name
			busStruct.nextStop = *&path[(pos+1)%len].Name

			for i := 0; i < busStruct.availSeats; i++ {
				if path[i%10].Q.Size != 0 {
					m[path[i%10].Q.Pop().Destination]++
					busStruct.passOn++
					busStruct.availSeats--
				}
			}
			busStruct.passOn -= m[busStruct.currStop]
			busStruct.availSeats += m[busStruct.currStop]
			m[busStruct.currStop] = 0

			fmt.Println(count, name, busStruct.currStop, busStruct.nextStop, busStruct.availSeats, busStruct.passOn)
			pos++
			count++
		} else {
			pos = 0
		}
	}
}
*/
