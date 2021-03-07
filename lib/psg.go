package rs

import (
	"math/rand"
	"ftm"
)

// Passenger create a passenger object
type Passenger struct {
	source      string
	destination string
}

//NewPassengerAt add passenger to specific bus stop
func NewPassengerAt(BusStop) {
	passengerGroup := []*passenger{}
	for i := 1; i < rand.Intn(100); i++ {
		var p *passenger
		p = new(passenger)
		p.source = BusStop.name
		rando := rand.Intn(len(allStation))
		p.destination = BusStop[rando]
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}

//NewPassenger add passenger to random bus stop
func NewPassenger(BusStop) {
	passengerGroup := []*passenger{}
	for i := 1; i < rand.Intn(20); i++ {
		var p *passenger
		p = new(passenger)
		rando1 := rand.Intn(len(allStation))
		p.source = BusStop[rando1]
		rando2 := rand.Intn(len(allStation))
		p.destination = BusStop[rando2]
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}
