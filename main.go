package main

import (
	"fmt"
	"math/rand"
	rs "rs/lib"
	"time"
)

var (
	stopList   []*rs.BusStop
	inputNoBus int
	globalHour int
	globalMin  int
	//for putting busc in main
	countPos int = 0
	count    int = 0
)

// Bus Struct
type Bus struct {
	availSeats int
	passOn     int
	currStop   string
	nextStop   string
}

//busc threading function---------------------------------------------------------------
func Busc(name string, path []*rs.BusStop) {
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
			fmt.Println(globalHour, globalMin)
			pos++
			count++
		} else {
			pos = 0
		}
	}
}

//End busc--------------------------------------------------------------------------------------------------------

func main() {

	stopList = append(stopList, &rs.BusStop{Name: "aBuilding", TimeTaken: 5})
	stopList = append(stopList, &rs.BusStop{Name: "bBuilding", TimeTaken: 6})
	stopList = append(stopList, &rs.BusStop{Name: "cBuilding", TimeTaken: 3})
	stopList = append(stopList, &rs.BusStop{Name: "dBuilding", TimeTaken: 4})
	stopList = append(stopList, &rs.BusStop{Name: "eBuilding", TimeTaken: 6})
	stopList = append(stopList, &rs.BusStop{Name: "fBuilding", TimeTaken: 3})
	stopList = append(stopList, &rs.BusStop{Name: "gBuilding", TimeTaken: 1})
	stopList = append(stopList, &rs.BusStop{Name: "hBuilding", TimeTaken: 5})
	stopList = append(stopList, &rs.BusStop{Name: "iBuilding", TimeTaken: 7})
	stopList = append(stopList, &rs.BusStop{Name: "jBuilding", TimeTaken: 4})
	fmt.Printf("Initiated bus stop list: %v\n", stopList)
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	var inputPsg int
	fmt.Println("How many passenger?")
	fmt.Scanln(&inputPsg)

	psgr := rs.NewPassenger1(stopList)
	rand.Seed(time.Now().UnixNano())
	//Passenger Generated -------------------------
	random1 := rs.Random(150, 200)
	rand.Seed(time.Now().UnixNano())
	random2 := rs.Random(50, 100)
	rand.Seed(time.Now().UnixNano())
	random3 := rs.Random(150, 200)

	//Cars Generated ------------------------------
	// cars1 := rs.CarGroup()

	// rs.TimeTick(globalHour, globalMin)

	// Init -------------------------------------------------
	if inputPsg != 0 {
		fmt.Println("Total Passenger :", inputPsg)
		rs.GnrPsg(stopList, inputPsg, psgr)
	} else {
		fmt.Println("Total Passenger :", random1)
		rs.GnrPsg(stopList, random1, psgr)
	}

	// // Event Class End --------------------------------------------
	if (globalMin % 60) == 0 {
		rs.GnrPsg(stopList, random1, psgr)
	}

	// Event train ----------------------------------------
	if (globalMin % 120) == 0 {
		rs.GnrPsgAt(stopList, "hBuilding", random2, psgr)
		cars := rs.CarGroup()
		fmt.Println("Total cars")
		fmt.Println(len(cars))
	}

	// // Event After 4pm ---------------------------------------------
	if globalHour == 16 {
		rs.GnrPsg(stopList, random3, psgr)
	}

	fmt.Println("#,BusName,CurrentStop,NextStop,AvailableSeats,TotalPassengerOnBus ")
	for i := 0; i < inputNoBus; i++ {
		go Busc("bus"+fmt.Sprint((i+1)), stopList)
	}
	//Added for timetick
	for {
		time.Sleep(time.Nanosecond * 100)
		rs.TimeTick(&globalHour, &globalMin)
		// fmt.Println(globalHour, globalMin)
	}
	rs.Busc("test", stopList)
	fmt.Println("Ending main package...")
}
