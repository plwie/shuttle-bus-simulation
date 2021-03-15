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
)

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

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
	random1 := random(1500, 2000)
	if inputPsg != 0 {
		fmt.Println("Total Passenger :", inputPsg)
		rs.GnrPsg(stopList, inputPsg, psgr)
	} else {
		fmt.Println("Total Passenger :", random1)
		rs.GnrPsg(stopList, random1, psgr)
	}

	fmt.Println("#,BusName,CurrentStop,NextStop,AvailableSeats,TotalPassengerOnBus ")
	for i := 0; i < inputNoBus; i++ {
		go rs.Busc("bus"+fmt.Sprint((i+1)), stopList)
	}
	rs.Busc("test", stopList)
	fmt.Println("Ending main package...")
}
