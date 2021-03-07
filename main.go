package main

import (
	"fmt"
	rs "rs/lib"
)

var (
	count    int = 0
	aBuilding = rs.BusStop{name: "aBuilding"}
	bBuilding = rs.BusStop{name: "bBuilding"}
	cBuilding = rs.BusStop{name: "cBuilding"}
	dBuilding = rs.BusStop{name: "dBuilding"}
	eBuilding = rs.BusStop{name: "eBuilding"}
	fBuilding = rs.BusStop{name: "fBuilding"}
	gBuilding = rs.BusStop{name: "gBuilding"}
	hBuilding = rs.BusStop{name: "hBuilding"}
	iBuilding = rs.BusStop{name: "iBuilding"}
	jBuilding = rs.BusStop{name: "jBuilding"}
	stopList []rs.BusStop{aBuilding, bBuilding, cBuilding, dBuilding, eBuilding, fBuilding, gBuilding, hBuilding, iBuilding, jBuilding}
	inputNoBus int
)

func main() {
	fmt.Println("This is the main package: ")
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	for i := 0; i < inputNoBus; i++ {
		go rs.busc("bus"+fmt.Sprint((i+1)), path,count)
	}
	psgr := NewPassenger(stopList)

	for _, ele := range psgr {
		for _, ele := range stopList {
			if psgr.ele.src == stopList.ele.name {
				stopList.q.Add(psgr.ele)
			} else if psgr.ele.src != stopList.ele.name {
				continue
			}
		}
	}

	rs.busc("test", path,count)
	fmt.Println("Ending main package...")
}
