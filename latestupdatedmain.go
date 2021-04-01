package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	rs "rs/lib"
	"sync"
	"time"
)

var (
	stopList   []*rs.BusStop
	inputNoBus int
	//for putting busc in main
	countPos       int
	count          int
	graph          = rs.Graph{}
	totalTime      float64
	passTotal      int
	waitingTime    float64
	totalPassenger int
	// bwg         sync.WaitGroup
	countBWG  *int
	worldTime int
	mutx      sync.Mutex
	BusArr    []*rs.Bus
	doOnce    sync.Once
)

//busc threading function---------------------------------------------------------------
func Busc(name int, path []*rs.BusStop, BusArr *rs.Bus, bwg *sync.WaitGroup) {
	var pos int
	// countPos++
	var lenPath int = len(path)
	var spd float64
	var dist float64
	var calcDist float64
	// var calcTime float64
	var distTrav float64
	var countPass int = 0
	var calculatedT int = 0
	// var pWaitTime *float64 = &waitingTime
	// var pPassTotal *int = &passTotal

	// // Assign key value
	// for i := 0; i < lenPath; i++ {
	// 	busStruct.M[path[i].Name] = 0
	// }

	//pass struct values

	//initial position
	//only enter this condition when first run simulation
	if BusArr.FirstTime == false {
		BusArr.Pos = name
		// fmt.Println("FIRST POS:", BusArr.Pos)
		BusArr.M = make(map[string]int)
		for i := 0; i < lenPath; i++ {
			BusArr.M[path[i].Name] = 0
		}
		BusArr.AvailSeats = 30
		BusArr.FirstTime = true
	}
	if BusArr.Pos >= 10 {
		BusArr.Pos = 0
	}

	pos = BusArr.Pos
	if BusArr.DistToNext == 0 {
		BusArr.DistToNext = float64(graph.Edges[pos].Cost)
	}
	dist = float64(graph.Edges[pos].Cost)
	spd = float64(graph.GetSpeed(path[pos], path[(pos+1)%lenPath]))
	distTrav = (dist / 60) * spd
	calcDist = BusArr.DistToNext

	BusArr.CurrStop = path[(pos)%lenPath].Name
	BusArr.NextStop = path[(pos+1)%lenPath].Name

	if (calcDist - distTrav) > 1 {
		//move 1 step
		calcDist -= distTrav
		BusArr.DistToNext = calcDist
		// fmt.Println(calcDist)
	} else {
		BusArr.DistToNext = 0
		mutx.Lock()
		rs.DropPass(BusArr)
		rs.GetPassngr(path, BusArr, &countPass, &calculatedT)
		// calculatedT = rs.GetPassngr(path, BusArr, &countPass, &calculatedT)
		if passTotal != totalPassenger {
			totalTime += (float64(calculatedT) * 60)
		}
		mutx.Unlock()
		BusArr.Pos++
		BusArr.CurrStop = path[BusArr.Pos%lenPath].Name
		BusArr.NextStop = path[(BusArr.Pos+1)%lenPath].Name
	}
	mutx.Lock()
	// fmt.Println(countPass)
	// fmt.Println(pPassTotal)
	passTotal += countPass
	// fmt.Println("PT:", passTotal)
	mutx.Unlock()
	bwg.Done()

}

//End busc--------------------------------------------------------------------------------------------------------

func main() {
	buildingInputJson := `{
		"busStopList": [
			{
				"source": "aBuilding",
				"destination": "bBuilding",
				"distance": 1
			},
			{
				"source": "bBuilding",
				"destination": "cBuilding",
				"distance": 2
			},
			{
				"source": "cBuilding",
				"destination": "dBuilding",
				"distance": 1
			},
			{
				"source": "dBuilding",
				"destination": "eBuilding",
				"distance": 2
			},
			{
				"source": "eBuilding",
				"destination": "fBuilding",
				"distance": 3
			},
			{
				"source": "fBuilding",
				"destination": "gBuilding",
				"distance": 1
			},
			{
				"source": "gBuilding",
				"destination": "hBuilding",
				"distance": 2
			},
			{
				"source": "hBuilding",
				"destination": "iBuilding",
				"distance": 1
			},
			{
				"source": "iBuilding",
				"destination": "jBuilding",
				"distance": 3
			},
			{
				"source": "jBuilding",
				"destination": "aBuilding",
				"distance": 2
			}
		]
	}`
	// Initialize "building", "stopList", "add weight to edge" and "generate traffic"
	graph.GenerateBuildingBusStop(&stopList, buildingInputJson)
	
	fmt.Printf("Initiated bus stop list: %v\n", stopList)
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	var proceedDecision string
	if inputNoBus > len(stopList) {
		fmt.Println("-------------------------------------------------------------------------------------------")
		fmt.Println("Warning: The number of Bus exceeded the number of bus stop")
		fmt.Println("the bus might overlapped with each other")
		fmt.Println("-------------------------------------------------------------------------------------------")
		fmt.Println("Would you like to proceed? y/N")
		fmt.Scanln(&proceedDecision)
		if proceedDecision == "N" {
			fmt.Println("Exiting the simulation...")
			fmt.Println("-------------------------------------------------------------------------------------------")
			os.Exit(1)
		}
	}
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
	fmt.Println("-The longer the queue, the longer the later person has to wait")
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Results: ")

	start := time.Now()
	psgr := rs.NewPassenger()
	rand.Seed(time.Now().UnixNano())
	//Passenger Generated -------------------------
	random1 := rs.Random(150, 200)

	totalPassenger = inputPsg
	// Init -------------------------------------------------
	if inputPsg != 0 {
		rs.GnrPsg(stopList, inputPsg, psgr)
		//rs.GnrTrf(CarGroup())
		totalPsg += inputPsg
	} else {
		rs.GnrPsg(stopList, random1, psgr)
		//rs.GnrTrf(CarGroup())
		totalPsg += random1
	}
	//create bus struct array

	for i := 0; i <= inputNoBus; i++ {
		BusArr = append(BusArr, &rs.Bus{})
	}

	for worldTime <= 600 {
		var bwg sync.WaitGroup
		bwg.Add(1)
		go rs.Event(&graph, stopList, psgr, worldTime, &bwg)
		bwg.Wait()
		worldTime++
		for i := 0; i < inputNoBus; i++ {
			bwg.Add(1)
			go Busc(i, stopList, BusArr[i], &bwg)
		}
		bwg.Wait()
		rs.IncreasePassengerWaitingTime(stopList)
	}

	waitingTime = ((totalTime) / float64(passTotal)) / 60
	secc := math.Round((((math.Mod(waitingTime, 1)) * 60) * 1000) / 1000)
	minn := (math.Floor(waitingTime / 1))
	fmt.Println("Average Passengers Waiting Time:", minn, "minutes", secc, "secs")
	fmt.Println("Total Passengers Delivered: ", passTotal)
	duration := time.Since(start)
	fmt.Println("Simulation run time: ", duration)
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Simulation has ended...")
}