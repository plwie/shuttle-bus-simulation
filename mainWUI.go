package main

import (
	"fmt"
	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	rs "rs/lib"
	"strconv"
	"sync"
	"time"

	"github.com/kbinani/screenshot"

	ui "github.com/gizak/termui/v3"
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
	//pos "does not refer to (x,y) coordinates of the bus"
	//Since the route is in 1D this Pos variable is to keep the bus in its track
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
				"distance": 2
			},
			{
				"source": "bBuilding",
				"destination": "cBuilding",
				"distance": 1
			},
			{
				"source": "cBuilding",
				"destination": "dBuilding",
				"distance": 3
			},
			{
				"source": "dBuilding",
				"destination": "eBuilding",
				"distance": 2
			},
			{
				"source": "eBuilding",
				"destination": "fBuilding",
				"distance": 2
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
				"distance": 3
			},
			{
				"source": "iBuilding",
				"destination": "jBuilding",
				"distance": 1
			},
			{
				"source": "jBuilding",
				"destination": "aBuilding",
				"distance": 1
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
		g := widgets.NewGauge()
		g.Title = "Bus " + strconv.Itoa(i) + ": Traveling from " + newBus.CurrStop + " to " + newBus.NextStop
		g.SetRect(50, i*8, 85, i*8+3)
		g.BarColor = ui.ColorRed
		g.TitleStyle.Fg = ui.ColorYellow

		renBus = append(renBus, g)

		l := widgets.NewList()
		l.Title = "Bus" + strconv.Itoa(i+1) + "Step: 0"
		l.TitleStyle.Fg = ui.ColorYellow
		l.WrapText = false
		l.SetRect(0, i*8, 48, i*8+8)
		l.Rows = []string{
			"Current Stop:" + BusArr[i].CurrStop,
			"Next Stop: " + BusArr[i].NextStop,
			"Psg on Bus:" + strconv.FormatInt(int64(BusArr[i].PassOn), 10),
			"Available Seats:" + strconv.Itoa(BusArr[i].AvailSeats),
			"Distance until next stop:" + strconv.FormatFloat(BusArr[i].DistToNext, 'f', -1, 64),
			"Psg down on next stop:" + strconv.FormatInt(int64(BusArr[i].M[BusArr[i].CurrStop]), 10),
		}
		renAt = append(renAt, l)
	}

	// Init termui
	if err := ui.Init(); err != nil {
		log.Fatalf("failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// PSG queue chart
	bstbc := widgets.NewBarChart()
	bstbc.Title = "Passenger in Queue"
	bstbc.SetRect(87, 0, 138, 15)
	bstbc.BarGap = 2
	var stopName []string
	for _, v := range stopList {
		stopName = append(stopName, v.Name[0:1])
	}
	bstbc.Labels = stopName
	// Global Timer
	tp := widgets.NewParagraph()
	tp.Title = "Current Time"
	tp.SetRect(140, 0, 170, 5)
	tp.TitleStyle.Fg = ui.ColorGreen

	//passenger delivered
	pd := widgets.NewParagraph()
	pd.Title = "Passenger Delivered"
	pd.TitleStyle.Fg = ui.ColorGreen
	pd.SetRect(140, 10, 170, 5)
	pd.TextStyle.Fg = ui.ColorWhite
	pd.BorderStyle.Fg = ui.ColorWhite

	//Event Log
	el := widgets.NewList()
	el.Title = "Event Log"
	el.TitleStyle.Fg = ui.ColorCyan
	el.WrapText = false
	el.SetRect(87, 16, 138, 35)

	ui.Render(el)

	drawEvent := func(lst []string) {
		el.Rows = lst
		el.ScrollDown()
		ui.Render(el)
	}
	drawBus := func(n int, step int) {
		now := BusArr[n].CurrStop
		next := BusArr[n].NextStop
		distTo := float64(getDist(now, next))
		if distTo == 0 {
			distTo = 1
		}
		distNow := float64(BusArr[n].DistToNext * 100)
		if distNow == 0 {
			distNow = 0
		}
		renBus[n].Title = "Bus " + strconv.Itoa(n+1) + ": " + now + " to " + next
		distFin := int(distNow / distTo)
		if distFin > 100 {
			distFin = 100
		}
		renBus[n].Percent = 100 - distFin

		renAt[n].Title = "Bus" + strconv.Itoa(n+1) + " Step:" + strconv.Itoa(step)
		renAt[n].Rows[0] = "Current Stop: " + BusArr[n].CurrStop
		renAt[n].Rows[1] = "Next Stop: " + BusArr[n].NextStop
		renAt[n].Rows[2] = "Psg on Bus: " + strconv.FormatInt(int64(BusArr[n].PassOn), 10)
		renAt[n].Rows[3] = "Available Seats: " + strconv.Itoa(BusArr[n].AvailSeats)
		renAt[n].Rows[4] = "Distance until next stop (KM): " + strconv.FormatFloat(BusArr[n].DistToNext, 'f', -1, 32)
		mutx.Lock()
		renAt[n].Rows[5] = "Psg down on next stop: " + strconv.FormatInt(int64(BusArr[n].M[BusArr[n].CurrStop]), 10)
		mutx.Unlock()

		ui.Render(renAt[n])
		ui.Render(renBus[n])
	}
	drawBST := func() {
		var psgNum []float64
		for _, v := range stopList {
			psgNum = append(psgNum, float64(v.Q.Size))
		}
		psgNum = append(psgNum, float64(1))
		bstbc.Data = psgNum
		ui.Render(bstbc)
	}
	drawTimer := func(n int) {
		tp.Text = strconv.Itoa(n/60) + " HR: " + strconv.Itoa(n%60) + " MIN"
		ui.Render(tp)
	}
	drawPassDev := func() {
		pd.Text = strconv.Itoa(passTotal) + " people"
		ui.Render(pd)
	}

	// Main simulation step
	event := ui.PollEvents()
	for worldTime < inputStep {
		var bwg sync.WaitGroup
		bwg.Add(1)
		go rs.Event(&graph, stopList, psgr, worldTime, &bwg, g)
		bwg.Wait()
		worldTime++
		if worldTime == g.AtTime+1 {
			info := ("At Time" + "_" + strconv.Itoa(g.AtTime) + "_" + "Event generate:" + "_" + strconv.Itoa(g.PsgAdded) + "_" + "Passengers")
			infoes = append(infoes, info)
		}
		drawEvent(infoes)

		for i := 0; i < inputNoBus; i++ {
			bwg.Add(1)
			go Busc(i, stopList, BusArr[i], &bwg)
			select {
			case e := <-event:
				switch e.ID {
				case "q", "<C-c>":
					return
				case "p":
					for f := range ui.PollEvents() {
						if f.Type == ui.KeyboardEvent {
							break
						}
					}
				}

			default:
				drawBus(i, worldTime)
				// drawBattributes(i, worldTime)
			}
		}
		bwg.Wait()
		rs.IncreasePassengerWaitingTime(stopList)
		time.Sleep(time.Millisecond)
		drawBST()
		drawTimer(worldTime)
		drawPassDev()
		//call screenshot function
		// getScreen(worldTime)
		// time.Sleep(time.Second / 2)
	}

	// Calculating simulation results
	duration := time.Since(start)
	waitingTime = (totalTime / float64(passTotal)) / 60
	secc := math.Round((((math.Mod(waitingTime, 1)) * 60) * 1000) / 1000)
	minn := (math.Floor(waitingTime / 1))

	// Print out result
	rsp := widgets.NewParagraph()
	rsp.Title = "RESULTS"
	rsp.TitleStyle.Fg = ui.ColorRed
	rl1 := "Average Passengers Waiting Time: " + strconv.FormatFloat(minn, 'f', -1, 32) + " minutes " + strconv.FormatFloat(secc, 'f', -1, 32) + " secs\n"
	rl2 := "Total Passengers Delivered: " + strconv.Itoa(passTotal) + "\n"
	rl3 := "Simulation run time: " + duration.String() + "\n"
	rl4 := "Simulation has ended...\n"
	rsp.Text = rl1 + rl2 + rl3 + rl4
	rsp.SetRect(87, 36, 170, 44)
	ui.Render(rsp)
	// fmt.Println("-------------------------------------------------------------------------------------------")
	// fmt.Println("RESULTS: ")
	// fmt.Println("Average Passengers Waiting Time:", minn, "minutes", secc, "secs")
	// fmt.Println("Total Passengers Delivered: ", passTotal)
	// fmt.Println("Simulation run time: ", duration)
	// fmt.Println("-------------------------------------------------------------------------------------------")
	// fmt.Println("Simulation has ended...")
	for e := range ui.PollEvents() {
		if e.Type == ui.KeyboardEvent {
			break
		}
	}
}
