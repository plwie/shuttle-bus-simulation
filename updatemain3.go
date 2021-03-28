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
	// fmt.Println(dist)
	BusArr.Status = "Traveling"
	BusArr.CurrStop = path[(pos)%lenPath].Name
	BusArr.NextStop = path[(pos+1)%lenPath].Name
	// fmt.Println("FIRSTCURR,FIRSTNEXT,:", BusArr.CurrStop, BusArr.NextStop)
	// fmt.Println("kut")
	if BusArr.Status == "Traveling" {
		// fmt.Println(calcDist)
		// fmt.Println(distTrav)
		// แก้ตรงifนี้ด้วย \/
		// fmt.Printf("Dist %f, Trav %f\n", dist, distTrav)
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
			if passTotal != totalPassenger {
				totalTime += (float64(calculatedT) * 60)
			}
			mutx.Unlock()
			BusArr.Pos++
			BusArr.CurrStop = path[BusArr.Pos%lenPath].Name
			BusArr.NextStop = path[(BusArr.Pos+1)%lenPath].Name
		}
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
	graph.AddEdge(&bBuilding, &aBuilding, 1)
	graph.AddEdge(&bBuilding, &cBuilding, 2)
	graph.AddEdge(&cBuilding, &bBuilding, 2)
	graph.AddEdge(&cBuilding, &dBuilding, 1)
	graph.AddEdge(&dBuilding, &cBuilding, 1)
	graph.AddEdge(&dBuilding, &eBuilding, 2)
	graph.AddEdge(&eBuilding, &dBuilding, 2)
	graph.AddEdge(&eBuilding, &fBuilding, 3)
	graph.AddEdge(&fBuilding, &eBuilding, 3)
	graph.AddEdge(&fBuilding, &gBuilding, 1)
	graph.AddEdge(&gBuilding, &fBuilding, 1)
	graph.AddEdge(&gBuilding, &hBuilding, 2)
	graph.AddEdge(&hBuilding, &gBuilding, 2)
	graph.AddEdge(&hBuilding, &iBuilding, 1)
	graph.AddEdge(&iBuilding, &hBuilding, 1)
	graph.AddEdge(&iBuilding, &jBuilding, 3)
	graph.AddEdge(&jBuilding, &iBuilding, 3)
	graph.AddEdge(&jBuilding, &aBuilding, 2)
	graph.AddEdge(&aBuilding, &jBuilding, 2)

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
	var proceedDecision string
	if inputNoBus > 10 {
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
	for worldTime <= 600 {
		worldTime++
		var bwg sync.WaitGroup
		// fmt.Println(BusArr[0])
		// fmt.Println(BusArr[1])
		// fmt.Println("WT:", worldTime)
		for i := 0; i < inputNoBus; i++ {
			bwg.Add(1)
			// go Busc(i, stopList, BusArr[i], &bwg)
			go Busc(i, stopList, BusArr[i], &bwg)
		}
		rs.IncreasePassengerWaitingTime(stopList)
		bwg.Wait()
	}

	waitingTime = ((totalTime) / float64(passTotal)) / 60
	secc := math.Round((((math.Mod(waitingTime, 1)) * 60) * 1000) / 1000)
	minn := (math.Floor(waitingTime / 1))
	fmt.Println("Average Passengers Waiting Time:", minn, "minutes", secc, "secs")
	fmt.Println("Total Passengers Delivered: ", passTotal)
	duration := time.Since(start)
	fmt.Println("Simulation run time: ", duration)
	// fmt.Println(len(stopList))
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Simulation has ended...")
}
