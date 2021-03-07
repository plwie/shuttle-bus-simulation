package main

import (
	"fmt"
	"math/rand"
	rs "rs/lib"
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
	fmt.Println("This is the main package: ")
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	for i := 0; i < inputNoBus; i++ {
		go rs.Busc("bus"+fmt.Sprint((i+1)), stopList)
	}

	psgr := rs.NewPassenger(stopList)

	for x := range psgr {
		for y := range stopList {
			if *&psgr[x].Source == *&stopList[y].Name {
				stopList[y].Q.Add(*psgr[x])
				fmt.Println(stopList[y].Q.Size)
				fmt.Println(stopList[y].Name)
			} else if *&psgr[x].Source != *&stopList[y].Name {
				continue
			}
		}
	}

	rs.Busc("test", stopList)
	fmt.Println("Ending main package...")
}
