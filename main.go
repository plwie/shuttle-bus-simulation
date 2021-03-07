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
	fmt.Println("This is the main package: ")
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	for i := 0; i < inputNoBus; i++ {
		go rs.Busc("bus"+fmt.Sprint((i+1)), stopList)
	}

	psgr := rs.NewPassenger1(stopList)
	rand.Seed(time.Now().UnixNano())
	random1 := random(50, 200)
	for i := 1; i < random1; i++ {
		psgr.Source = *&stopList[random(0, 5)].Name
		psgr.Destination = *&stopList[rand.Intn((5-0-1)+1)].Name
		for i := 0; i < len(stopList)-1; i++ {
			if psgr.Source == *&stopList[i].Name {
				stopList[i].Q.Add(*psgr)
				fmt.Println(stopList[i].Name)
				fmt.Println(stopList[i].Q.Size)
			}
		}
	}
	rs.Busc("test", stopList, count)
	fmt.Println("Ending main package...")
}
