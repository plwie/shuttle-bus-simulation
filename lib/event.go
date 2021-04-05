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
	graph.GenerateTraffic(CarGroup(), nil, nil)
	fmt.Println("Event generate:", n, "Passengers")

}

// Train ----------------------------------------
func Train(graph *Graph, lst []*BusStop, p *Passenger) {
	n := Random(50, 100)
	n1 := Random(150, 200)
	GnrPsgAt(lst, "hBuilding", n, p)
	GnrPsg(lst, n1, p)
	graph.GenerateTraffic(CarGroupTr(), lst[7], nil)
	fmt.Println("Event generate:", n+n1, "Passengers")
}

// // AfterWork ---------------------------------------------
func AfterWork(graph *Graph, lst []*BusStop, p *Passenger) {
	n := Random(350, 500)
	GnrPsg(lst, n, p)
	graph.GenerateTraffic(CarGroupBusy(), nil, nil)
	fmt.Println("Event generate:", n, "Passengers")
}

func Event(graph *Graph, lst []*BusStop, p *Passenger, time int, wg *sync.WaitGroup) {
	if time != 0 {
		if time <= 480 && time%120 == 60 {
			fmt.Println("At time:", time)
			ClassEnd(graph, lst, p)
		} else if time <= 480 && time%120 == 0 {
			fmt.Println("At time:", time)
			Train(graph, lst, p)
		} else if time%60 == 0 && time >= 480 {
			fmt.Println("At time:", time)
			AfterWork(graph, lst, p)
		}
	}
	wg.Done()
}
