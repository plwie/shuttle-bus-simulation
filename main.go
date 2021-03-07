package main

import (
	"fmt"
	rs "rs/lib"
)

var (
	count    int = 0
	aBuilding = BusStop{name: "aBuilding"}
	bBuilding = BusStop{name: "bBuilding"}
	cBuilding = BusStop{name: "cBuilding"}
	dBuilding = BusStop{name: "dBuilding"}
	eBuilding = BusStop{name: "eBuilding"}
	fBuilding = BusStop{name: "fBuilding"}
	gBuilding = BusStop{name: "gBuilding"}
	hBuilding = BusStop{name: "hBuilding"}
	iBuilding = BusStop{name: "iBuilding"}
	jBuilding = BusStop{name: "jBuilding"}
	stopList []rs.BusStop{aBuilding, bBuilding, cBuilding, dBuilding, eBuilding, fBuilding, gBuilding, hBuilding, iBuilding, jBuilding}
)

func main() {
	fmt.Println("This is the main package: ")

	fmt.Println("Ending main package...")
}
