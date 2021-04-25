
package main

import (
	"fmt"
	"image/png"
	"math"
	"math/rand"
	"os"
	rs "rs/lib"
	"sync"
	"time"

	"github.com/kbinani/screenshot"

	"github.com/gizak/termui/v3/widgets"
)

var (
	stopList       []*rs.BusStop
	inputNoBus     int
	countPos       int
	count          int
	graph          = rs.Graph{}
	totalTime      float64
	passTotal      int
	waitingTime    float64
	totalPassenger int
	countBWG       *int
	worldTime      int
	mutx           sync.Mutex
	BusArr         []*rs.Bus
	doOnce         sync.Once
	renBus         []*widgets.Gauge
	infoes         []string
	baList         []string
	renAt          []*widgets.List
	// calcDist       float64
)

// Busc run a separate thread for each bus instance
func Busc(name int, path []*rs.BusStop, BusArr *rs.Bus, bwg *sync.WaitGroup) {
	var pos int
	var lenPath int = len(path)
	var spd float64
	// var dist float64
	// var calcDist float64
	var distTrav float64
	var countPass int = 0
	var calculatedT int = 0

	// First time initialize
	if BusArr.FirstTime == false {
		BusArr.Pos = name
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
	// dist = float64(graph.Edges[pos].Cost)
	spd = float64(graph.GetSpeed(path[pos], path[(pos+1)%lenPath]))
	distTrav = spd / 60
	// calcDist = BusArr.DistToNext

	BusArr.CurrStop = path[(pos)%lenPath].Name
	BusArr.NextStop = path[(pos+1)%lenPath].Name

	if (BusArr.DistToNext - distTrav) > 0.1 {
		// Move 1 step
		// calcDist -= distTrav
		BusArr.DistToNext -= distTrav

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
	mutx.Lock()
	passTotal += countPass
	mutx.Unlock()
	bwg.Done()

}

// Get distance from src and dst
func getDist(src string, dst string) int {
	for _, v := range graph.Edges {
		if src == v.Parent.Name && dst == v.Child.Name {
			return v.Cost
		}
	}
	return 0
}

//function to get screenshot
func getScreen(n int) {
	// n := screenshot.NumActiveDisplays()

	bounds := screenshot.GetDisplayBounds(0)

	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	fileName := fmt.Sprintf("%d_%dx%d.png", n, bounds.Dx(), bounds.Dy())
	file, _ := os.Create(fileName)
	defer file.Close()
	png.Encode(file, img)

	fmt.Printf("#%d : %v \"%s\"\n", n, bounds, fileName)

}

func main() {
	buildingInputJson := `{
		"busStopList": [
			{
				"source": "aBuilding",
				"destination": "bBuilding",
				"distance": 10,
				"speedlimit": 40
			},
			{
				"source": "bBuilding",
				"destination": "cBuilding",
				"distance": 20,
				"speedlimit": 50
			},
			{
				"source": "cBuilding",
				"destination": "dBuilding",
				"distance": 10,
				"speedlimit": 40
			},
			{
				"source": "dBuilding",
				"destination": "eBuilding",
				"distance": 20,
				"speedlimit": 50
			},
			{
				"source": "eBuilding",
				"destination": "fBuilding",
				"distance": 30,
				"speedlimit": 50
			},
			{
				"source": "fBuilding",
				"destination": "gBuilding",
				"distance": 10,
				"speedlimit": 30
			},
			{
				"source": "gBuilding",
				"destination": "hBuilding",
				"distance": 20,
				"speedlimit": 30
			},
			{
				"source": "hBuilding",
				"destination": "iBuilding",
				"distance": 10,
				"speedlimit": 50
			},
			{
				"source": "iBuilding",
				"destination": "jBuilding",
				"distance": 30,
				"speedlimit": 50
			},
			{
				"source": "jBuilding",
				"destination": "aBuilding",
				"distance": 20,
				"speedlimit": 50
			}
		]
	}`
	// Initialize "building", "stopList", "add weight to edge" and "generate traffic"
	graph.GenerateBuildingBusStop(&stopList, buildingInputJson)
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Printf("INITIATED BUS STOP LIST:\n")
	for _, v := range stopList {
		fmt.Printf("%v ", v.Name)
	}
	fmt.Println()

	// Get input and check for invalid bus number input
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Printf("ENTER THE NUMBER OF BUS: ")
	fmt.Scanln(&inputNoBus)
	for inputNoBus <= 0 {
		fmt.Printf("ERROR: invalid bus number input\n")
		fmt.Printf("ENTER THE NUMBER OF BUS: ")
		fmt.Scanln(&inputNoBus)
	}

	// Warning for bus number exceeding the bus stop number
	if inputNoBus > len(stopList) {
		var proceedDecision string
		fmt.Println("-------------------------------------------------------------------------------------------")
		fmt.Println("Warning: The number of Bus exceeded the number of bus stop.")
		fmt.Println("A bus might overlap with other busses.")
		fmt.Println("-------------------------------------------------------------------------------------------")
		fmt.Println("Would you like to proceed? Type N to cancel")
		fmt.Scanln(&proceedDecision)
		if proceedDecision == "N" {
			fmt.Println("Exiting the simulation...")
			fmt.Println("-------------------------------------------------------------------------------------------")
			os.Exit(1)
		}
	}
	fmt.Println("-------------------------------------------------------------------------------------------")

	// Get input and check invalid passenger input
	var inputPsg int
	fmt.Printf("ENTER THE NUMBER OF PASSENGER: ")
	fmt.Scanln(&inputPsg)
	for inputPsg < 0 {
		fmt.Printf("ERROR: invalid passenger number input\n")
		fmt.Printf("ENTER THE NUMBER OF PASSENGER: ")
		fmt.Scanln(&inputPsg)
	}

	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("REMINDER:")
	fmt.Println("-Not all passengers will be delivered")
	fmt.Println("-The simulation will stop after the time reaches the threshold")
	fmt.Println("-Events will add more passengers into the simulation")
	fmt.Println("-Waiting Time depends directly on the traffic of the road")
	fmt.Println("-More traffic means bus can travel slower")
	fmt.Println("-The longer the queue, the longer the later person has to wait")
	fmt.Println("-------------------------------------------------------------------------------------------")

	// Get and check invalid simulation step
	var inputStep int
	fmt.Printf("ENTER THE NUMBER OF HOUR: ")
	fmt.Scanln(&inputStep)
	for inputStep <= 0 {
		fmt.Printf("ERROR: invalid time input\n")
		fmt.Printf("ENTER THE NUMBER OF HOUR: ")
		fmt.Scanln(&inputStep)
	}
	inputStep *= 60

	// Generating Passenger
	start := time.Now()
	psgr := rs.NewPassenger()
	rand.Seed(time.Now().UnixNano())
	random1 := rs.Random(150, 200)
	totalPassenger = inputPsg

	if inputPsg != 0 {
		rs.GnrPsg(stopList, inputPsg, psgr)
	} else {
		rs.GnrPsg(stopList, random1, psgr)
	}

	g := rs.NewGlobDis()

	// Create bus instance and put in array
	for i := 0; i < inputNoBus; i++ {
		newBus := &rs.Bus{}
		BusArr = append(BusArr, newBus)
	}

	for worldTime < inputStep {
		var bwg sync.WaitGroup
		bwg.Add(1)
		go rs.Event(&graph, stopList, psgr, worldTime, &bwg, g)
		bwg.Wait()
		worldTime++

		for i := 0; i < inputNoBus; i++ {
			bwg.Add(1)
			go Busc(i, stopList, BusArr[i], &bwg)
		}
		bwg.Wait()
		rs.IncreasePassengerWaitingTime(stopList)
	}
	duration := time.Since(start)
	waitingTime = (totalTime / float64(passTotal)) / 60
	secc := math.Round((((math.Mod(waitingTime, 1)) * 60) * 1000) / 1000)
	minn := (math.Floor(waitingTime / 1))

	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("RESULTS: ")
	fmt.Println("Average Passengers Waiting Time:", minn, "minutes", secc, "secs")
	fmt.Println("Total Passengers Delivered: ", passTotal)
	fmt.Println("Simulation run time: ", duration)
	fmt.Println("-------------------------------------------------------------------------------------------")
	fmt.Println("Simulation has ended...")
}