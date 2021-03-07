package rs

import (
	"math/rand"
)

// Passenger create a passenger object
type Passenger struct {
	source      string
	destination string
	next        *Passenger
}

//NewPassengerAt add passenger to specific bus stop
func NewPassengerAt(stopList []BusStop, start BusStop) []*Passenger {
	// passengerGroup := []*Passenger{}
	for i := 1; i < rand.Intn(100); i++ {
		var p *Passenger
		p = new(Passenger)
		p.source = start.Name
		rando := rand.Intn(len(stopList))
		p.destination = stopList[rando].Name
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			if p.source == BusStop.Name {
				BusStop.q.Add(psgr[i])
			}
		}
	}
	// return passengerGroup
}

//NewPassenger add passenger to random bus stop
func NewPassenger(stopList []BusStop) []*Passenger {
	// passengerGroup := []*Passenger{}
	for i := 1; i < rand.Intn(20); i++ {
		var p *Passenger
		p = new(Passenger)
		rando1 := rand.Intn(len(stopList))
		p.source = stopList[rando1].Name
		rando2 := rand.Intn(len(stopList))
		p.destination = stopList[rando2].Name
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			if p.source == BusStop.Name {
				BusStop.q.Add(psgr[i])
			}
		}
	}
	// return passengerGroup
}
