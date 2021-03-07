package main

import (
	"fmt"
	rs "rs/lib"
)

var (
	count      int = 0
	stopList   []rs.BusStop
	inputNoBus int
)

func main() {
	stopList = append(stopList, rs.BusStop{Name: "Stop 1"})
	stopList = append(stopList, rs.BusStop{Name: "Stop 2"})
	stopList = append(stopList, rs.BusStop{Name: "Stop 3"})
	stopList = append(stopList, rs.BusStop{Name: "Stop 4"})
	stopList = append(stopList, rs.BusStop{Name: "Stop 5"})
	stopList = append(stopList, rs.BusStop{Name: "Stop 6"})
	fmt.Printf("Initiated bus stop list: %v\n", stopList)

	fmt.Println("This is the main package: ")
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	for i := 0; i < inputNoBus; i++ {
		go rs.Busc("bus"+fmt.Sprint((i+1)), stopList, count)
	}

	psgr := rs.NewPassenger(stopList)

	for _, ele := range psgr {
		for _, ele := range stopList {
			if psgr.ele.src == stopList.ele.name {
				stopList.q.Add(psgr.ele)
			} else if psgr.ele.src != stopList.ele.name {
				continue
			}
		}
	}

	rs.Busc("test", stopList, count)
	fmt.Println("Ending main package...")
}
