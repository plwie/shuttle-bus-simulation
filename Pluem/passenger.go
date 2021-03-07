package rs

import (
	"math/rand"
	"time"
)

// Passenger create a passenger object
type Passenger struct {
	Source      string
	Destination string
	next        *Passenger
}

//NewPassengerAt add passenger to specific bus stop
func NewPassengerAt(stopList []BusStop, start BusStop) []*Passenger {
	passengerGroup := []*Passenger{}
	for i := 1; i < rand.Intn(100); i++ {
		var p *Passenger
		p = new(Passenger)
		p.Source = start.Name
		rand.Seed(time.Now().Unix())
		rando := rand.Intn(len(stopList))
		p.Destination = stopList[rando].Name
		if p.Destination == p.Source {
			continue
		} else if p.Destination != p.Source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}

//NewPassenger add passenger to random bus stop
func NewPassenger(stopList []BusStop) []*Passenger {
	passengerGroup := []*Passenger{}
	// rand.Seed(time.Now().Unix())
	for i := 1; i < rand.Intn(20); i++ {
		var p *Passenger
		p = new(Passenger)
		rand.Seed(time.Now().Unix())
		rando1 := rand.Intn(len(stopList) - 1)
		p.Source = stopList[rando1].Name
		// rand.Seed(time.Now().Unix())
		rando2 := rand.Intn(len(stopList) - 1)
		p.Destination = stopList[rando2].Name
		if p.Destination == p.Source {
			continue
		} else if p.Destination != p.Source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}
