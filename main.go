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
)

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

	rs.TimeTick(globalHour, globalMin)

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
		go rs.Busc("bus"+fmt.Sprint((i+1)), stopList)
	}
	rs.Busc("test", stopList)

	fmt.Println("Ending main package...")
}
