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
	//for putting busc in main
	countPos int = 0
	count    int = 0
)

// Bus Struct
type Bus struct {
	availSeats int
	passOn     int
	currStop   string
	nextStop   string
}

//busc threading function---------------------------------------------------------------
func Busc(name string, path []*rs.BusStop) {
	//need to declare global count = 0
	m := make(map[string]int)
	pos := countPos
	countPos++
	var len int = len(path)
	var count int = 0
	//create bus struct instance
	busStruct := Bus{
		availSeats: 30,
		passOn:     0,
		currStop:   *&path[pos].Name,
		nextStop:   *&path[pos+1].Name,
	}
	for i := 0; i < 10; i++ {
		m[path[i].Name] = 0
	}
	//code for bus traveling (busstop to another busstop)
	for {
		if pos < len && name != "test" {
			// time.Sleep(time.Second * 1)
			busStruct.currStop = *&path[pos].Name
			busStruct.nextStop = *&path[(pos+1)%len].Name

			for i := 0; i < busStruct.availSeats; i++ {
				if path[i%10].Q.Size != 0 {
					m[path[i%10].Q.Pop().Destination]++
					busStruct.passOn++
					busStruct.availSeats--
				}
			}
			busStruct.passOn -= m[busStruct.currStop]
			busStruct.availSeats += m[busStruct.currStop]
			m[busStruct.currStop] = 0

			fmt.Println(count, name, busStruct.currStop, busStruct.nextStop, busStruct.availSeats, busStruct.passOn)
			fmt.Println(globalHour, globalMin)
			pos++
			count++
		} else {
			pos = 0
		}
	}
}

//End busc--------------------------------------------------------------------------------------------------------

func main() {

	graph := rs.Graph{}

	aBuilding := rs.BusStop{Name: "aBuilding", TimeTaken: 5}
	bBuilding := rs.BusStop{Name: "bBuilding", TimeTaken: 6}
	cBuilding := rs.BusStop{Name: "cBuilding", TimeTaken: 3}
	dBuilding := rs.BusStop{Name: "dBuilding", TimeTaken: 4}
	eBuilding := rs.BusStop{Name: "eBuilding", TimeTaken: 6}
	fBuilding := rs.BusStop{Name: "fBuilding", TimeTaken: 3}
	gBuilding := rs.BusStop{Name: "gBuilding", TimeTaken: 1}
	hBuilding := rs.BusStop{Name: "hBuilding", TimeTaken: 10}

	tempList := graph.StopList

	tempList = append(tempList, &aBuilding)
	tempList = append(tempList, &bBuilding)
	tempList = append(tempList, &cBuilding)
	tempList = append(tempList, &dBuilding)
	tempList = append(tempList, &eBuilding)
	tempList = append(tempList, &fBuilding)
	tempList = append(tempList, &gBuilding)
	tempList = append(tempList, &hBuilding)
	// stopList = append(stopList, &hBuilding)
	// stopList = append(stopList, &iBuilding)
	// stopList = append(stopList, &jBuilding)

	graph.AddEdge(&aBuilding, &bBuilding, 1)
	//  graph.AddEdge(&aBuilding, &hBuilding, 1)
	graph.AddEdge(&bBuilding, &aBuilding, 1)
	graph.AddEdge(&bBuilding, &cBuilding, 1)
	graph.AddEdge(&cBuilding, &bBuilding, 1)
	graph.AddEdge(&cBuilding, &dBuilding, 1)
	graph.AddEdge(&dBuilding, &cBuilding, 1)
	graph.AddEdge(&dBuilding, &eBuilding, 1)
	graph.AddEdge(&eBuilding, &dBuilding, 1)
	graph.AddEdge(&eBuilding, &fBuilding, 1)
	graph.AddEdge(&fBuilding, &eBuilding, 1)
	graph.AddEdge(&fBuilding, &gBuilding, 1)
	graph.AddEdge(&gBuilding, &fBuilding, 1)
	graph.AddEdge(&gBuilding, &hBuilding, 1)
	graph.AddEdge(&hBuilding, &gBuilding, 1)
	graph.AddEdge(&hBuilding, &aBuilding, 1)
	graph.AddEdge(&aBuilding, &hBuilding, 1)
	graph.GenerateTraffic(rs.CarGroup(), &aBuilding, &bBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &bBuilding, &aBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &bBuilding, &cBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &cBuilding, &bBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &cBuilding, &dBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &dBuilding, &cBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &dBuilding, &eBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &eBuilding, &dBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &eBuilding, &fBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &fBuilding, &eBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &fBuilding, &gBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &gBuilding, &fBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &gBuilding, &hBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &hBuilding, &gBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &hBuilding, &aBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &aBuilding, &hBuilding)

	//  graph.GenerateTraffic(, &aBuilding, &bBuilding)
	//  graph.GenerateTraffic(&bBuilding, &aBuilding)
	//  graph.GenerateTraffic(&bBuilding, &cBuilding)
	//  graph.GenerateTraffic(&cBuilding, &bBuilding)
	//  graph.GenerateTraffic(&cBuilding, &dBuilding)
	//  graph.GenerateTraffic(&dBuilding, &cBuilding)
	//  graph.GenerateTraffic(&dBuilding, &eBuilding)
	//  graph.GenerateTraffic(&eBuilding, &dBuilding)
	//  graph.GenerateTraffic(&eBuilding, &fBuilding)
	//  graph.GenerateTraffic(&fBuilding, &eBuilding)
	//  graph.GenerateTraffic(&fBuilding, &gBuilding)
	//  graph.GenerateTraffic(&gBuilding, &fBuilding)
	//  graph.GenerateTraffic(&gBuilding, &hBuilding)
	//  graph.GenerateTraffic(&hBuilding, &gBuilding)
	//  graph.GenerateTraffic(&hBuilding, &aBuilding)
	//  graph.GenerateTraffic(&aBuilding, &hBuilding)

	graph.GetSpeed(&aBuilding, &bBuilding) // 0
	//  graph.AddEdge(&fBuidling, &gBuilding, 1)
	//  graph.AddEdge(&d, &e, 2)
	//  graph.AddEdge(&d, &g, 30)
	//  graph.AddEdge(&d, &f, 10)
	//  graph.AddEdge(&f, &g, 1)

	//  fmt.Println(graph.String())

	for _, nodeStart := range tempList {
		costTable := graph.Dijkstra(nodeStart)
		// Make the costTable nice to read :)
		fmt.Printf("Start node is %s\n", nodeStart.Name)
		for node, cost := range costTable {
			fmt.Printf("Distance from %s to %s = %d km\n", nodeStart.Name, node.Name, cost)
		}
		fmt.Println("----------------------")
	}

	fmt.Println(aBuilding.Name)
	fmt.Println(graph.GetSpeed(&aBuilding, &bBuilding))
	fmt.Println(graph.Edges[1].Parent)
	fmt.Println(graph.Edges[1].Child)
	fmt.Println(graph.Edges[1].Cost)

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

	// rs.TimeTick(globalHour, globalMin)

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
		go Busc("bus"+fmt.Sprint((i+1)), stopList)
	}
	//Added for timetick
	for {
		time.Sleep(time.Nanosecond * 100)
		rs.TimeTick(&globalHour, &globalMin)
		// fmt.Println(globalHour, globalMin)
	}
	// rs.Busc("test", stopList)
	// fmt.Println("Ending main package...")
}
