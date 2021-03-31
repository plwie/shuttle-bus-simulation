package rs

import (
	"sync"
)

var (
	mut sync.Mutex
)

// ClassEnd --------------------------------------------
func ClassEnd(graph *Graph,lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(150, 200), p)
	// generate all roads
	graph.GenerateTraffic(CarGroup(), nil, nil)

}

// Train ----------------------------------------
func Train(graph *Graph,lst []*BusStop, p *Passenger) {
	GnrPsgAt(lst, "hBuilding", Random(50, 100), p)
	// generate roads that head to train station
	graph.GenerateTraffic(CarGroupTr(), lst[7], nil) // lst[7] -> &hBuilding
}

// // AfterWork ---------------------------------------------
func AfterWork(graph *Graph, lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(350, 500), p)
	// generate all roads
	graph.GenerateTraffic(CarGroupBusy(), nil, nil)
}

func Event(graph *Graph, lst []*BusStop, p *Passenger, time int, wg *sync.WaitGroup) {
	if time != 0 {
		if time <= 480 && time%120 == 60 { // 60, 180, 300, 420
			ClassEnd(graph, lst, p)
		} else if time <= 480 && time%120 == 0 { // (120, 240, 360, 480) x 2
			Train(graph, lst, p)
			ClassEnd(graph, lst, p)
		} else if time >= 480 && time%120 == 60 { // 540
			AfterWork(graph, lst, p)
		}
	}
	wg.Done()
}