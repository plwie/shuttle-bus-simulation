package rs

import (
	"sync"
)

var (
	mut sync.Mutex
)

// Globdis ---------------------------------------------
type Globdis struct {
	PsgAdded int
	AtTime   int
}

//NewPassenger1 Create a new passenger
func NewGlobDis() *Globdis {
	var g Globdis
	return &g
}

// ClassEnd --------------------------------------------
func ClassEnd(graph *Graph, lst []*BusStop, p *Passenger, g *Globdis) {
	n := Random(150, 200)
	GnrPsg(lst, n, p)
	graph.GenerateTraffic(CarGroup(), nil, nil)
	g.PsgAdded = n
	// fmt.Println("Event generate:", n, "Passengers")

}

// Train ----------------------------------------
func Train(graph *Graph, lst []*BusStop, p *Passenger, g *Globdis) {
	n := Random(50, 100)
	n1 := Random(150, 200)
	GnrPsgAt(lst, "hBuilding", n, p)
	GnrPsg(lst, n1, p)
	graph.GenerateTraffic(CarGroupTr(), lst[7], nil)
	g.PsgAdded = n + n1
	// fmt.Println("Event generate:", n+n1, "Passengers")
}

// // AfterWork ---------------------------------------------
func AfterWork(graph *Graph, lst []*BusStop, p *Passenger, g *Globdis) {
	n := Random(350, 500)
	GnrPsg(lst, n, p)
	graph.GenerateTraffic(CarGroupBusy(), nil, nil)
	g.PsgAdded = n
	// fmt.Println("Event generate:", n, "Passengers")
}

func Event(graph *Graph, lst []*BusStop, p *Passenger, time int, wg *sync.WaitGroup, g *Globdis) {
	if time != 0 {
		if time <= 480 && time%120 == 60 {
			// fmt.Println("At time:", time)
			ClassEnd(graph, lst, p, g)
			g.AtTime = time
		} else if time <= 480 && time%120 == 0 {
			// fmt.Println("At time:", time)
			Train(graph, lst, p, g)
			g.AtTime = time
		} else if time%120 == 60 && time >= 480 {
			// fmt.Println("At time:", time)
			AfterWork(graph, lst, p, g)
			g.AtTime = time
		}
	}
	wg.Done()
}
