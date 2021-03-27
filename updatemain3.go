package main

import (
	"fmt"
	"math"
	"math/rand"
	rs "rs/lib"
	"sync"
	"time"
)

var (
	stopList   []*rs.BusStop
	inputNoBus int
	//for putting busc in main
	countPos    int     = 0
	count       int     = 0
	graph               = rs.Graph{}
	totalTime   float64 = 0
	passTotal   int     = 0
	waitingTime float64 = 0
	bwg         sync.WaitGroup
	countBWG    *int
	worldTime   int
	mutx        sync.Mutex
	BusArr      []*rs.Bus
)

//busc threading function---------------------------------------------------------------
func Busc(name int, path []*rs.BusStop, BusArr *rs.Bus, bwg *sync.WaitGroup) {
	// pos := countPos
	// countPos++
	var lenPath int = len(path)
	var spd float64
	var dist float64
	var calcDist float64
	var distTrav float64
	var countPass int = 0
	// var pWaitTime *float64 = &waitingTime
	var pPassTotal *int = &passTotal

	// // Assign key value
	// for i := 0; i < lenPath; i++ {
	// 	busStruct.M[path[i].Name] = 0
	// }

	//pass struct values

	//initial position
	//only enter this condition when first run simulation
	if BusArr.FirstTime == false {
		BusArr.Pos = name
		BusArr.M = make(map[string]int)
		for i := 0; i < lenPath; i++ {
			BusArr.M[path[i].Name] = 0
		}
		BusArr.FirstTime = true
	}
	if BusArr.Pos >= 10 {
		BusArr.Pos = 0
	}

	calcDist = dist
	dist = float64(graph.Edges[BusArr.Pos].Cost)
	spd = float64(graph.GetSpeed(path[BusArr.Pos], path[(BusArr.Pos+1)%lenPath]))
	distTrav = dist - ((dist*1000)/60)*spd

	if BusArr.CurrStop == "Traveling" {
		if calcDist-distTrav > 0 {
			//move 1 step
			BusArr.AtStop = false
			calcDist = dist - distTrav
		}
		if calcDist-distTrav <= 1 {
			rs.DropPass(BusArr)
			rs.GetPassngr(path, BusArr, &countPass)
			mutx.Lock()
			totalTime += (dist / spd)
			mutx.Unlock()
			BusArr.Pos++
		}
	}
	mutx.Lock()
	*pPassTotal += countPass
	mutx.Unlock()
	bwg.Done()

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
	graph.AddEdge(&cBuilding, &dBuilding, 1)
	graph.AddEdge(&dBuilding, &eBuilding, 2)
	graph.AddEdge(&eBuilding, &fBuilding, 3)
	graph.AddEdge(&fBuilding, &gBuilding, 1)
	graph.AddEdge(&gBuilding, &hBuilding, 2)
	graph.AddEdge(&hBuilding, &iBuilding, 1)
	graph.AddEdge(&iBuilding, &jBuilding, 3)
	graph.AddEdge(&jBuilding, &aBuilding, 2)

	//GnrTrf Function---------------------------------------------------------------
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
	//-----------------------------------------------------------------

	fmt.Printf("Initiated bus stop list: %v\n", stopList)
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	var inputPsg int
	totalPsg := 0
	fmt.Println("How many initial passenger?")
	fmt.Scanln(&inputPsg)
	fmt.Println("Simulation in progress.....")
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("REMINDER:")
	fmt.Println("-Not all passengers will be delivered")
	fmt.Println("-The simulation will stop after the time reaches the threshold")
	fmt.Println("-Events will add more passengers into the simulation")
	fmt.Println("-Waiting Time depends directly on the traffic of the road")
	fmt.Println("-More traffic means bus can travel slower")
	fmt.Println("-Waiting Time is calculated from the passengers that are successfully delivered")
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Results: ")

	start := time.Now()
	psgr := rs.NewPassenger()
	rand.Seed(time.Now().UnixNano())
	//Passenger Generated -------------------------
	random1 := rs.Random(150, 200)

	// rs.TimeTick(globalHour, globalMin)

	// Init -------------------------------------------------
	if inputPsg != 0 {
		// fmt.Println("Total initiate Passenger : %v/n", inputPsg)
		rs.GnrPsg(stopList, inputPsg, psgr)
		//rs.GnrTrf(CarGroup())
		totalPsg += inputPsg
	} else {
		// fmt.Println("Total initiate Passenger : %v/n", random1)
		rs.GnrPsg(stopList, random1, psgr)
		//rs.GnrTrf(CarGroup())
		totalPsg += random1
	}
	//create bus struct array

	for i := 0; i <= inputNoBus; i++ {
		BusArr = append(BusArr, &rs.Bus{})
	}

	// fmt.Println("#,BusName,CurrentStop,NextStop,AvailableSeats,TotalPassengerOnBus ")
	for worldTime <= 120 {
		worldTime++
		for i := 0; i < inputNoBus; i++ {
			bwg.Add(1)
			go Busc(i, stopList, BusArr[i], &bwg)
		}
		// go rs.ConTimeTick(&globalHour, &globalMin, stopList, psgr)

		// bussTest := rs.Bus{}
		// bwg.Add(1)
		// go Busc(inputNoBus-1, stopList, &bussTest, &bwg)
		bwg.Wait()
	}

	fmt.Println(totalTime)
	fmt.Println(waitingTime)
	fmt.Println(passTotal)
	waitingTime = ((totalTime) / float64(passTotal)) / 60
	secc := math.Round((((math.Mod(waitingTime, 1)) * 60) * 1000) / 1000)
	minn := (math.Floor(waitingTime / 1))
	fmt.Println("Waiting Time:", minn, "minutes", secc, "secs")
	fmt.Println("Total Passengers Delivered: ", passTotal)
	duration := time.Since(start)
	fmt.Println("Simulation run time: ", duration)
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Simulation has ended...")
}
