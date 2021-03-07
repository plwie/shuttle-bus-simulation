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

	// ------------------------------------------ของปลื้ม-----------------------------------------------
	rand.Seed(time.Now().Unix())
	// rando2 := rand.Intn(len(stopList))
	// fmt.Println(stopList[rando2].Name)
	// fmt.Println(len(stopList))

	psgr := rs.NewPassenger(stopList)

	fmt.Println(psgr)
	fmt.Println(*psgr[0])
	fmt.Println(psgr[0])
	fmt.Println(*&psgr[3].Source)
	// fmt.Println(*&psgr[3].Destination)
	// for i := 1; i < len(psgr)-1; i++ {
	// 	fmt.Println(*&psgr[i].Source)
	// 	fmt.Println(*&psgr[3].Destination)
	// }
	stopList[3].Q.Add(*psgr[2])
	stopList[3].Q.Add(*psgr[0])
	stopList[3].Q.Add(*psgr[3])
	fmt.Println(stopList[2].WaitingPassenger)
	// for i := 1; i < len(psgr)-1; i++{
	// 	for i := 1; i < len(psgr)-1; i++
	// }

	// fmt.Println(stopList)
	// for _, u := range psgr {
	// 	for _, ele := range stopList {
	// 		if psgr.u.src == stopList.ele.name {
	// 			stopList.q.Add(psgr[i])
	// 		} else if psgr.ele.src != stopList.ele.name {
	// 			continue
	// 		}
	// 	}
	// }

	// ------------------------------------------ของปลื้ม-----------------------------------------------

	// rs.Busc("test", stopList, count)
	// fmt.Println("Ending main package...")
}
