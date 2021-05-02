package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	rs "rs/lib"
	"strconv"
	"sync"
	"time"

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
	renBusT        []*widgets.List
	infoes         []string
	baList         []string
	renBusCheck    []*widgets.List
	passCheck      int
	totalCheck     int
	gnrPsg         int
	busState       int
	busStateTot    int
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
	var prevCount int

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
		rs.DropPass(BusArr, &countPass)
		rs.GetPassngr(path, BusArr, &calculatedT)
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
	prevCount = BusArr.PassOn
	mutx.Unlock()
	bwg.Done()

	if prevCount != BusArr.PassOn {
		BusArr.State = "Passenger has been dropped off/pick up on the road"
	} else {
		BusArr.State = "Passed"
	}

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

func main() {
	buildingInputJson := `{
		"busStopList": [
			{
				"source": "aBuilding",
				"destination": "bBuilding",
				"distance": 2,
				"speedlimit": 40
			},
			{
				"source": "bBuilding",
				"destination": "cBuilding",
				"distance": 1,
				"speedlimit": 50
			},
			{
				"source": "cBuilding",
				"destination": "dBuilding",
				"distance": 3,
				"speedlimit": 40
			},
			{
				"source": "dBuilding",
				"destination": "eBuilding",
				"distance": 2,
				"speedlimit": 50
			},
			{
				"source": "eBuilding",
				"destination": "fBuilding",
				"distance": 2,
				"speedlimit": 50
			},
			{
				"source": "fBuilding",
				"destination": "gBuilding",
				"distance": 1,
				"speedlimit": 30
			},
			{
				"source": "gBuilding",
				"destination": "hBuilding",
				"distance": 2,
				"speedlimit": 30
			},
			{
				"source": "hBuilding",
				"destination": "iBuilding",
				"distance": 3,
				"speedlimit": 50
			},
			{
				"source": "iBuilding",
				"destination": "jBuilding",
				"distance": 1,
				"speedlimit": 50
			},
			{
				"source": "jBuilding",
				"destination": "aBuilding",
				"distance": 1,
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

	// Init termui
	if err := ui.Init(); err != nil {
		log.Fatalf("Failed to initialize termui: %v", err)
	}
	defer ui.Close()

	// Generating Passenger
	start := time.Now()
	psgr := rs.NewPassenger()
	rand.Seed(time.Now().UnixNano())
	random1 := rs.Random(150, 200)
	totalPassenger = inputPsg

	if inputPsg != 0 {
		rs.GnrPsg(stopList, inputPsg, psgr)
		gnrPsg = inputPsg
	} else {
		rs.GnrPsg(stopList, random1, psgr)
		gnrPsg = random1
	}

	g := rs.NewGlobDis()

	// Create bus instance and put in array
	for i := 0; i < inputNoBus; i++ {
		// Create Bus
		newBus := &rs.Bus{AvailSeats: 30, PassOn: 0}
		BusArr = append(BusArr, newBus)

		if i >= 7 {
			continue
		}
		// Create Bus Info Box
		bc := widgets.NewList()
		bc.Title = "Bus" + strconv.Itoa(i+1) + " Step: 0"
		bc.TitleStyle.Fg = ui.ColorBlue
		bc.SetRect(0, i*8, 29, i*8+8)
		bc.Rows = []string{
			"PassengerOn: ",
			"AvailableSeats: ",
			"Psg down on next stop: ",
			"SeatsStatus: ",
			"BusStatus",
		}
		renBusCheck = append(renBusCheck, bc)

		//  Create Travel Gauge
		g := widgets.NewGauge()
		g.Title = newBus.CurrStop + " to " + newBus.NextStop
		g.SetRect(30, i*8, 60, i*8+3)
		g.BarColor = ui.ColorRed
		g.TitleStyle.Fg = ui.ColorYellow
		renBus = append(renBus, g)

		// Create Travel Info
		bti := widgets.NewList()
		bti.SetRect(30, i*8+3, 60, i*8+8)
		renBusT = append(renBusT, bti)
	}

	// Create Seat Check Per Step Log
	cs := widgets.NewList()
	cs.Title = "Seat Check Per Step"
	cs.TitleStyle.Fg = ui.ColorMagenta
	cs.Rows = []string{}
	cs.SetRect(61, 0, 89, 18)

	// Create Passenger drop Check Per Step Log
	cpd := widgets.NewList()
	cpd.Title = "Psg Drop Check Per Step"
	cpd.TitleStyle.Fg = ui.ColorMagenta
	cpd.Rows = []string{}
	cpd.SetRect(90, 0, 118, 18)

	// Create Passenger drop Check Conclusion Log
	cp := widgets.NewParagraph()
	cp.Title = "Psg Drop Check Results"
	cp.TitleStyle.Fg = ui.ColorGreen
	cp.SetRect(90, 18, 118, 22)

	// Create Seat Check Conclusion Log
	c := widgets.NewParagraph()
	c.Title = "Seat Check Results"
	c.TitleStyle.Fg = ui.ColorGreen
	c.SetRect(61, 18, 89, 22)

	// Create PSG queue chart
	bstbc := widgets.NewBarChart()
	bstbc.Title = "Passenger in Queue"
	bstbc.SetRect(119, 0, 170, 15)
	bstbc.BarGap = 2
	var stopName []string
	for _, v := range stopList {
		stopName = append(stopName, v.Name[0:1])
	}
	bstbc.Labels = stopName

	// Create Event Log
	el := widgets.NewList()
	el.Title = "Event Log (Total Passenger Generated: " + strconv.Itoa(gnrPsg) + ")"
	el.TitleStyle.Fg = ui.ColorCyan
	el.WrapText = false
	initInfo := ("At Time" + "_" + strconv.Itoa(g.AtTime) + "_" + "Event generate:" + "_" + strconv.Itoa(gnrPsg) + "_" + "Passengers")
	infoes = append(infoes, initInfo)
	el.SetRect(119, 16, 170, 35)

	// Create Psg Delivered Log
	pd := widgets.NewParagraph()
	pd.Title = "Passenger Delivered"
	pd.TitleStyle.Fg = ui.ColorGreen
	pd.SetRect(61, 22, 89, 26)
	pd.TextStyle.Fg = ui.ColorWhite
	pd.BorderStyle.Fg = ui.ColorWhite

	// Create Global Timer
	tp := widgets.NewParagraph()
	tp.Title = "Current Time"
	tp.SetRect(90, 22, 118, 26)
	tp.TitleStyle.Fg = ui.ColorGreen

	// Create Test Log
	tl := widgets.NewList()
	tl.Title = "TEST RESULTS"
	tl.TitleStyle.Fg = ui.ColorRed
	tl.SetRect(61, 26, 118, 35)
	ui.Render(tl)

	// Create Result Log
	rsp := widgets.NewParagraph()
	rsp.Title = "SIMULATION RESULTS"
	rsp.TitleStyle.Fg = ui.ColorRed
	rsp.SetRect(61, 36, 118, 44)

	// Create Error Log
	erl := widgets.NewList()
	erl.Title = "ERROR LOG"
	erl.TitleStyle.Fg = ui.ColorRed
	erl.SetRect(119, 36, 170, 44)
	ui.Render(erl)
	var erlst []string

	// Draw function
	drawEvent := func(lst []string) {
		el.Rows = lst
		el.Title = "Event Log (Total Passenger Generated: " + strconv.Itoa(gnrPsg) + ")"
		el.ScrollDown()
		ui.Render(el)
	}
	drawBus := func(n int) {
		if n >= 7 {
			return
		}
		// Gauge
		now := BusArr[n].CurrStop
		next := BusArr[n].NextStop
		distTo := float64(getDist(now, next))
		distNow := float64(BusArr[n].DistToNext * 100)
		dt := float64(getDist(now, next)) - BusArr[n].DistToNext
		distFin := int(distNow / distTo)
		if distFin > 100 {
			distFin = 100
		}
		renBus[n].Title = now + " to " + next
		renBus[n].Percent = 100 - distFin
		ui.Render(renBus[n])
		// Box
		var tl []string
		tl = append(tl, "Current: "+now)
		tl = append(tl, "Next: "+next)
		if dt < 0 {
			dt = 0
		}
		tl = append(tl, "Dist: "+strconv.FormatFloat(dt, 'f', -1, 32)+"/"+strconv.Itoa(int(distTo))+" KM")
		renBusT[n].Rows = tl
		ui.Render(renBusT[n])
	}
	drawBusCheck := func(n int, step int, check *int, totalCheck *int, stateCheck *int, totalState *int) {
		var checkResult string
		if BusArr[n].PassOn+BusArr[n].AvailSeats == 30 {
			checkResult = "SeatsStatus: Passed"
			mutx.Lock()
			*check++
			*totalCheck++
			mutx.Unlock()
		} else {
			checkResult = "SeatsStatus: Failed"
			erlst = append(erlst, "ERROR: Invalid Bus Seat at Step "+strconv.Itoa(step))
			erlst = append(erlst, "(On: "+strconv.Itoa(BusArr[n].PassOn)+", Empty: "+strconv.Itoa(BusArr[n].AvailSeats)+")")
			erl.Rows = erlst
			ui.Render(erl)

		}
		if BusArr[n].State == "Passenger has been dropped off/pick up on the road" {
			erlst = append(erlst, "At time "+strconv.Itoa(step), BusArr[n].State)
		} else {
			*stateCheck++
			*totalState++
		}
		if n >= 7 {
			return
		}
		renBusCheck[n].Title = "Bus" + strconv.Itoa(n) + "Step: " + strconv.Itoa(step)
		renBusCheck[n].Rows[0] = "PassengerOn: " + strconv.Itoa(BusArr[n].PassOn)
		renBusCheck[n].Rows[1] = "AvailableSeats: " + strconv.Itoa(BusArr[n].AvailSeats)
		renBusCheck[n].Rows[2] = "Psg down on next stop: " + strconv.FormatInt(int64(BusArr[n].M[BusArr[n].CurrStop]), 10)
		renBusCheck[n].Rows[3] = checkResult
		renBusCheck[n].Rows[4] = "Bus Status: " + BusArr[n].State
		ui.Render(renBusCheck[n])
	}
	drawStepCheck := func(step int) {
		cs.Rows = append(cs.Rows, "Step: "+strconv.Itoa(step)+" Passed: "+strconv.Itoa(passCheck)+"/"+strconv.Itoa(inputNoBus))
		cs.ScrollDown()
		ui.Render(cs)
	}
	drawPassenCheck := func(step int) {
		cpd.Rows = append(cpd.Rows, "Step: "+strconv.Itoa(step)+" Status: "+strconv.Itoa(busState)+"/"+strconv.Itoa(inputNoBus))
		cpd.ScrollDown()
		ui.Render(cpd)
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
	drawBusResult := func() {
		c.Text = "Results: " + strconv.Itoa(totalCheck) + "/" + strconv.Itoa(inputNoBus*inputStep)
		ui.Render(c)
	}
	drawPsgResult := func() {
		cp.Text = "Results: " + strconv.Itoa(busStateTot) + "/" + strconv.Itoa(inputNoBus*inputStep)
		ui.Render(cp)
	}

	// Main simulation step
	event := ui.PollEvents()
	ui.Render(rsp)
	for worldTime < inputStep {
		var bwg sync.WaitGroup
		bwg.Add(1)
		go rs.Event(&graph, stopList, psgr, worldTime, &bwg, g)
		bwg.Wait()
		worldTime++
		for i := 0; i < inputNoBus; i++ {
			bwg.Add(1)
			go Busc(i, stopList, BusArr[i], &bwg)
			drawBusCheck(i, worldTime, &passCheck, &totalCheck, &busState, &busStateTot)
			drawBus(i)
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
				break
			}

		}
		bwg.Wait()
		rs.IncreasePassengerWaitingTime(stopList)

		// Render and Update
		if worldTime == g.AtTime+1 && worldTime > 59 {
			gnrPsg += g.PsgAdded
			info := ("At Time" + "_" + strconv.Itoa(g.AtTime) + "_" + "Event generate:" + "_" + strconv.Itoa(g.PsgAdded) + "_" + "Passengers")
			infoes = append(infoes, info)
		}
		drawEvent(infoes)
		drawTimer(worldTime)
		drawPassDev()
		drawBusResult()
		drawPsgResult()
		drawBST()
		drawStepCheck(worldTime)
		drawPassenCheck(worldTime)
		passCheck = 0
		busState = 0

		//call screenshot function
		// getScreen(worldTime)

		// time.Sleep(time.Second)
	}

	// Calculating simulation results
	duration := time.Since(start)
	waitingTime = (totalTime / float64(passTotal)) / 60
	secc := math.Round((((math.Mod(waitingTime, 1)) * 60) * 1000) / 1000)
	minn := (math.Floor(waitingTime / 1))

	// Print out simulation result
	rl1 := "Average Passengers Waiting Time: " + strconv.FormatFloat(minn, 'f', -1, 32) + " minutes " + strconv.FormatFloat(secc, 'f', -1, 32) + " secs\n"
	rl2 := "Total Passengers Delivered: " + strconv.Itoa(passTotal) + "\n"
	rl3 := "Simulation run time: " + duration.String() + "\n"
	rsp.Text = rl1 + rl2 + rl3
	ui.Render(rsp)

	// Print out testing result
	var trlst []string
	var inBus int
	var inQ int
	initPsg := passTotal
	for _, v := range stopList {
		inQ += v.Q.Size
	}
	for _, v := range BusArr {
		inBus += v.PassOn
	}
	totalPsgTrack := initPsg + inBus + inQ
	tr1 := "PSG_Tracked/PSG_Generated: " + strconv.Itoa(totalPsgTrack) + " / " + strconv.Itoa(gnrPsg)
	tr2 := "(Init / In_Bus / In_Stop: " + strconv.Itoa(initPsg) + " / " + strconv.Itoa(inBus) + " / " + strconv.Itoa(inQ) + ")\n"
	tr3 := "Seat Check Status: " + strconv.Itoa(totalCheck) + " / " + strconv.Itoa(inputNoBus*inputStep)
	tr4 := "Bus Check Status: " + strconv.Itoa(busStateTot) + " / " + strconv.Itoa(inputNoBus*inputStep)
	trlst = append(trlst, tr1)
	trlst = append(trlst, tr2)
	trlst = append(trlst, tr3)
	trlst = append(trlst, tr4)
	tl.Rows = trlst
	if totalPsgTrack != gnrPsg {
		erlst = append(erlst, "ERROR: Final Passenger Incorrect")
		erl.Rows = erlst
		ui.Render(erl)
	}
	if totalCheck != (inputNoBus * inputStep) {
		erlst = append(erlst, "ERROR: Final Seat Incorrect")
		erl.Rows = erlst
		ui.Render(erl)
	}
	ui.Render(tl)

	// Wait for keyboard exit
	for exitEvent := range event {
		if exitEvent.Type == ui.KeyboardEvent {
			break
		}
	}
}
