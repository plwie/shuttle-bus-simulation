package rs

import (
	"fmt"
	"sync"
)

var (
	mut sync.Mutex
)

// ClassEnd --------------------------------------------
func ClassEnd(graph *Graph, lst []*BusStop, p *Passenger) {
	n := Random(150, 200)
	GnrPsg(lst, n, p)
	// generate all roads
	graph.GenerateTraffic(CarGroup(), nil, nil)

	fmt.Println("Event generate:", n, "Passengers")

}

// Train ----------------------------------------
func Train(graph *Graph, lst []*BusStop, p *Passenger) {
	n := Random(50, 100)
	GnrPsgAt(lst, "hBuilding", n, p)
	// generate roads that head to train station
	graph.GenerateTraffic(CarGroupTr(), lst[7], nil) // lst[7] -> &hBuilding
	fmt.Println("Event generate At HBuilding:", n, "Passengers")
}

// // AfterWork ---------------------------------------------
func AfterWork(graph *Graph, lst []*BusStop, p *Passenger) {
	n := Random(350, 500)
	GnrPsg(lst, n, p)
	// generate all roads
	graph.GenerateTraffic(CarGroupBusy(), nil, nil)
	fmt.Println("Event generate:", n, "Passengers")
}

func Event(graph *Graph, lst []*BusStop, p *Passenger, time int, wg *sync.WaitGroup) {
	if time != 0 {
		if time <= 480 && time%120 == 60 { // 60, 180, 300, 420
			fmt.Println("At time:", time)
			ClassEnd(graph, lst, p)
		} else if time <= 480 && time%120 == 0 { // (120, 240, 360, 480) x 2
			fmt.Println("At time:", time)
			Train(graph, lst, p)
			ClassEnd(graph, lst, p)
		} else if time >= 480 && time%120 == 60 { // 540
			fmt.Println("At time:", time)
			AfterWork(graph, lst, p)
		}
	}
	wg.Done()
}
