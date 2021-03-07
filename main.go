package main

import (
	"fmt"
	rs "rs/lib"
)

var (
	stopList   []rs.BusStop
	inputNoBus int
)

func main() {
	stopList = append(stopList, rs.BusStop{Name: "aBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "bBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "cBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "dBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "eBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "fBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "gBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "hBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "iBuilding"})
	stopList = append(stopList, rs.BusStop{Name: "jBuilding"})
	fmt.Printf("Initiated bus stop list: %v\n", stopList)

	fmt.Println("This is the main package: ")
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	for i := 0; i < inputNoBus; i++ {
		go rs.Busc("bus"+fmt.Sprint((i+1)), stopList)
	}

	psgr := rs.NewPassenger(stopList)

	fmt.Printf("PSGR: %T\n", psgr)
	// for _, u := range psgr {
	// 	for _, ele := range stopList {
	// 		if psgr.ele.src == stopList.ele.name {
	// 			stopList.q.Add(psgr(i))
	// 		} else if psgr.ele.src != stopList.ele.name {
	// 			continue
	// 		}
	// 	}
	// }

	rs.Busc("test", stopList)
	fmt.Println("Ending main package...")
}
