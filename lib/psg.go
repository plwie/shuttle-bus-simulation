package rs

import (
	"math/rand"
)

type Passenger struct {
	source      string
	destination string
}

type BusStation struct {
	name string
}

func NewPassengerAt(busStation) {
	allStation := [7]string{"Engineering", "Science", "Administration", "IT", "Liberal Arts", "Aggricultural", "Train Station"}
	passengerGroup := []*passenger{}
	for i := 1; i < rand.Intn(100); i++ {
		var p *passenger
		p = new(passenger)
		p.source = busStation.name
		p.destination = rand.Intn(len(allStation))
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			passengerGroup = append(passengerGroup, passenger)
		}
	}
	return passengerGroup
}

func NewPassenger() {
	allStation := [7]string{"Engineering", "Science", "Administration", "IT", "Liberal Arts", "Aggricultural", "Train Station"}
	passengerGroup := []*passenger{}
	for i := 1; i < rand.Intn(20); i++ {
		var p *passenger
		p = new(passenger)
		p.source = allStation[rand.Intn(len(allStation))]
		p.destination = allStation[rand.Intn(len(allStation))]
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			passengerGroup = append(passengerGroup, passenger)
		}
	}
	return passengerGroup
}
