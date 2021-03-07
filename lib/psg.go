package rs

import (
	"math/rand"
	"ftm"
)

// Passenger create a passenger object
type Passenger struct {
	source      string
	destination string
	next *Passenger
}

//NewPassengerAt add passenger to specific bus stop
func NewPassengerAt(BusStop) {
	passengerGroup := []*Passenger{}
	for i := 1; i < rand.Intn(100); i++ {
		var p *Passenger
		p = new(Passenger)
		p.source = BusStop.name
		rando := rand.Intn(len(BusStop))
		p.destination = BusStop[rando].name
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
	passengerGroup := []*Passenger{}
	for i := 1; i < rand.Intn(20); i++ {
		var p *Passenger
		p = new(Passenger)
		rando1 := rand.Intn(len(BusStop))
		p.source = BusStop[rando1].name
		rando2 := rand.Intn(len(BusStop))
		p.destination = BusStop[rando2].name
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}
