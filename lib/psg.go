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

func random(min int, max int) int {
	return rand.Intn(max-min) + min
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
func NewPassenger(stopList []*BusStop) []*Passenger {
	passengerGroup := []*Passenger{}
	rand.Seed(time.Now().Unix())
	random1 := random(5, 10)
	for i := 1; i < random1; i++ {
		var p *Passenger
		p = new(Passenger)
		rand.Seed(time.Now().Unix())
		rando1 := random(1, 5)
		p.Source = stopList[rando1].Name
		rand.Seed(time.Now().Unix())
		rando2 := random(1, 5)
		p.Destination = stopList[rando2].Name
		if p.Destination == p.Source {
			continue
		} else if p.Destination != p.Source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}
