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
	for i := 1; i < rand.Intn(10000); i++ {
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
	random1 := random(3000, 10000)
	for i := 1; i < random1; i++ {
		var p *Passenger
		p = new(Passenger)
		p.Source = stopList[4].Name
		p.Destination = stopList[3].Name
		if p.Destination == p.Source {
			continue
		} else if p.Destination != p.Source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}

// Random code from https://gist.github.com/201Factory/5ef7c2d46cf848db16041cafa17ab054
