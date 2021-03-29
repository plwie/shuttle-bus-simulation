package rs

import (
	"sync"
)

var (
	mut sync.Mutex
)

// ClassEnd --------------------------------------------
func ClassEnd(lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(150, 200), p)
	//rs.GnrTrf(CarGroup())

}

// Train ----------------------------------------
func Train(lst []*BusStop, p *Passenger) {
	GnrPsgAt(lst, "hBuilding", Random(50, 100), p)
	//rs.GnrTrf(CarGroupTr())
}

// // AfterWork ---------------------------------------------
func AfterWork(lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(350, 500), p)
	//rs.GnrTrf(CarGroupBusy())
}

func Event(lst []*BusStop, p *Passenger, time int, wg *sync.WaitGroup) {
	if time != 0 {
		if time <= 480 && time%120 == 60 {
			ClassEnd(lst, p)
		} else if time <= 480 && time%120 == 0 {
			Train(lst, p)
			ClassEnd(lst, p)
		} else if time >= 480 && time%120 == 60 {
			AfterWork(lst, p)
		}
	}
	wg.Done()
}
