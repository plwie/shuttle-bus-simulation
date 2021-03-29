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
	totalPsg := 0
	fmt.Println("How many passenger?")
	fmt.Scanln(&inputPsg)

	// rs.TimeTick(globalHour, globalMin)
	psgr := rs.NewPassenger()
	carsTr := rs.CarGroupTr()
	cars := rs.CarGroup()
	rand.Seed(time.Now().UnixNano())

	//Passenger Generated -------------------------
	random1 := rs.Random(150, 200)
	rand.Seed(time.Now().UnixNano())
	random2 := rs.Random(50, 100)
	rand.Seed(time.Now().UnixNano())
	random3 := rs.Random(150, 200)

	//Cars Generated ------------------------------
	// cars1 := rs.CarGroup()
	// fmt.Println("Total cars")
	// fmt.Println(len(carsTr))

	// Init -------------------------------------------------
	if inputPsg != 0 {
		fmt.Println("Total Passenger :", inputPsg)
		rs.GnrPsg(stopList, inputPsg, psgr)
		totalPsg += inputPsg
	} else {
		fmt.Println("Total Passenger :", random1)
		rs.GnrPsg(stopList, random1, psgr)
		totalPsg += random1
	}

	// // Event Class End --------------------------------------------
	if (globalHour%1) == 0 && globalMin == 0 {
		rs.GnrPsg(stopList, random1, psgr)
		totalPsg += random1
		fmt.Println(len(cars))
	}

	// Event train ----------------------------------------
	if (globalHour%2) == 0 && globalMin == 0 {
		rs.GnrPsgAt(stopList, "hBuilding", random2, psgr)
		totalPsg += random2
		fmt.Println("Total cars")
		fmt.Println(len(carsTr))
	}

	// // Event After 4pm ---------------------------------------------
	if globalHour == 16 && globalMin == 0 {
		rs.GnrPsg(stopList, random3, psgr)
		totalPsg += random3
	}

	fmt.Println("#,BusName,CurrentStop,NextStop,AvailableSeats,TotalPassengerOnBus ")
	for i := 0; i < inputNoBus; i++ {
		go rs.Busc("bus"+fmt.Sprint((i+1)), stopList)
	}
	rs.Busc("test", stopList)

	fmt.Println("Ending main package...")
}
