package main

import (
	"fmt"
	"math"
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
	countPos  int     = 0
	count     int     = 0
	graph             = rs.Graph{}
	totalTime float64 = 0
	passTotal int     = 0
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
	// graph := rs.Graph{}
	m := make(map[string]int)
	pos := countPos
	countPos++
	var len int = len(path)
	var count int = 0
	var countPass int = 0

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
			time.Sleep(time.Millisecond * 50)
			busStruct.currStop = *&path[pos].Name
			busStruct.nextStop = *&path[(pos+1)%len].Name

			for i := 0; i < busStruct.availSeats; i++ {
				if path[i%10].Q.Size != 0 {
					m[path[i%10].Q.Pop().Destination]++
					busStruct.passOn++
					countPass++
					busStruct.availSeats--

				}
			}
			busStruct.passOn -= m[busStruct.currStop]
			busStruct.availSeats += m[busStruct.currStop]
			m[busStruct.currStop] = 0

			fmt.Println(count, name, busStruct.currStop, busStruct.nextStop, busStruct.availSeats, busStruct.passOn)
			// fmt.Println(globalHour, globalMin)
			spd := float64(graph.GetSpeed(path[pos], path[(pos+1)%len]))
			dist := float64(graph.Edges[pos].Cost)
			calcTime := float64(math.Round(((dist/spd)*3600)*100) / 100)
			totalTime += (calcTime * float64(countPass))
			fmt.Println("|distance:", dist, "|speed:", spd, "|time:", calcTime, "sec", "|totalTime:", totalTime)
			passTotal += countPass
			fmt.Println("|countpass", countPass, "|passTotal", passTotal, "totaltime: ", totalTime)
			pos++
			count++

			countPass = 0
		} else {
			pos = 0
		}
	}
}

//End busc--------------------------------------------------------------------------------------------------------

func main() {

	//WaitingTime
	//waitingTime := totalTime/float64(passTotal)

	// graph := rs.Graph{}

	aBuilding := rs.BusStop{Name: "aBuilding"}
	bBuilding := rs.BusStop{Name: "bBuilding"}
	cBuilding := rs.BusStop{Name: "cBuilding"}
	dBuilding := rs.BusStop{Name: "dBuilding"}
	eBuilding := rs.BusStop{Name: "eBuilding"}
	fBuilding := rs.BusStop{Name: "fBuilding"}
	gBuilding := rs.BusStop{Name: "gBuilding"}
	hBuilding := rs.BusStop{Name: "hBuilding"}
	iBuilding := rs.BusStop{Name: "iBuilding"}
	jBuilding := rs.BusStop{Name: "jBuilding"}

	stopList := graph.StopList

	stopList = append(stopList, &aBuilding)
	stopList = append(stopList, &bBuilding)
	stopList = append(stopList, &cBuilding)
	stopList = append(stopList, &dBuilding)
	stopList = append(stopList, &eBuilding)
	stopList = append(stopList, &fBuilding)
	stopList = append(stopList, &gBuilding)
	stopList = append(stopList, &hBuilding)
	stopList = append(stopList, &iBuilding)
	stopList = append(stopList, &jBuilding)

	graph.AddEdge(&aBuilding, &bBuilding, 1)
	graph.AddEdge(&bBuilding, &cBuilding, 2)
	graph.AddEdge(&cBuilding, &dBuilding, 3)
	graph.AddEdge(&dBuilding, &eBuilding, 4)
	graph.AddEdge(&eBuilding, &fBuilding, 5)
	graph.AddEdge(&fBuilding, &gBuilding, 6)
	graph.AddEdge(&gBuilding, &hBuilding, 7)
	graph.AddEdge(&hBuilding, &iBuilding, 8)
	graph.AddEdge(&iBuilding, &jBuilding, 9)
	graph.AddEdge(&jBuilding, &aBuilding, 10)
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
	graph.GenerateTraffic(rs.CarGroup(), &iBuilding, &hBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &iBuilding, &jBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &jBuilding, &iBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &jBuilding, &aBuilding)

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
	graph.GetSpeed(&bBuilding, &aBuilding)
	graph.GetSpeed(&bBuilding, &cBuilding)
	graph.GetSpeed(&cBuilding, &bBuilding)
	graph.GetSpeed(&cBuilding, &dBuilding)
	graph.GetSpeed(&dBuilding, &cBuilding)
	graph.GetSpeed(&dBuilding, &eBuilding)
	graph.GetSpeed(&eBuilding, &dBuilding)
	graph.GetSpeed(&eBuilding, &fBuilding)
	graph.GetSpeed(&fBuilding, &eBuilding)
	graph.GetSpeed(&fBuilding, &gBuilding)
	graph.GetSpeed(&gBuilding, &fBuilding)
	graph.GetSpeed(&gBuilding, &hBuilding)
	graph.GetSpeed(&hBuilding, &gBuilding)
	graph.GetSpeed(&hBuilding, &aBuilding)
	graph.GetSpeed(&iBuilding, &hBuilding)
	graph.GetSpeed(&iBuilding, &jBuilding)
	graph.GetSpeed(&jBuilding, &iBuilding)
	graph.GetSpeed(&jBuilding, &aBuilding)
	//  graph.AddEdge(&fBuidling, &gBuilding, 1)
	//  graph.AddEdge(&d, &e, 2)
	//  graph.AddEdge(&d, &g, 30)
	//  graph.AddEdge(&d, &f, 10)
	//  graph.AddEdge(&f, &g, 1)

	//  fmt.Println(graph.String())

	// for _, nodeStart := range stopList {
	// 	costTable := graph.Dijkstra(nodeStart)
	// 	// Make the costTable nice to read :)
	// 	fmt.Printf("Start node is %s\n", nodeStart.Name)
	// 	for node, cost := range costTable {
	// 		fmt.Printf("Distance from %s to %s = %d km\n", nodeStart.Name, node.Name, cost)
	// 	}
	// 	fmt.Println("----------------------")
	// }

	fmt.Printf("Initiated bus stop list: %v\n", stopList)
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	var inputPsg int
	fmt.Println("How many passenger?")
	fmt.Scanln(&inputPsg)

	psgr := rs.NewPassenger1(stopList)
	rand.Seed(time.Now().UnixNano())
	//Passenger Generated -------------------------
	// random1 := rs.Random(150, 200)
	// rand.Seed(time.Now().UnixNano())
	// random2 := rs.Random(50, 100)
	// rand.Seed(time.Now().UnixNano())
	// random3 := rs.Random(150, 200)

	//Cars Generated ------------------------------
	// cars1 := rs.CarGroup()

	// rs.TimeTick(globalHour, globalMin)

	// Init -------------------------------------------------
	fmt.Println("Total Passenger :", inputPsg)
	rs.GnrPsg(stopList, inputPsg, psgr)
	// // Event Class End --------------------------------------------
	// if (globalMin % 60) == 0 {
	// 	rs.GnrPsg(stopList, inputPsg, psgr)
	// }

	// Event train ----------------------------------------
	// if (globalMin % 120) == 0 {
	// 	rs.GnrPsgAt(stopList, "hBuilding", random2, psgr)
	// 	cars := rs.CarGroup()
	// 	fmt.Println("Total cars")
	// 	fmt.Println(len(cars))
	// }

	// // Event After 4pm ---------------------------------------------
	// if globalHour == 16 {
	// 	rs.GnrPsg(stopList, random3, psgr)
	// }

	fmt.Println("#,BusName,CurrentStop,NextStop,AvailableSeats,TotalPassengerOnBus ")
	for i := 0; i < inputNoBus; i++ {
		go Busc("bus"+fmt.Sprint((i+1)), stopList)
	}
	//Added for timetick
	// for {
	// 	time.Sleep(time.Nanosecond * 100)
	// 	rs.TimeTick(&globalHour, &globalMin)
	// 	// fmt.Println(globalHour, globalMin)
	// }
	Busc("test", stopList)
	fmt.Println("Ending main package...")
}
