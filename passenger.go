package main

import (
	"math/rand"
)

type passenger struct {
	source      string
	destination string
}

type busStation struct {
	name string
}

func newPassengerAt(busStation) {
	allStation := []*busStation{busStation{"Engineering"}, busStation{"Science"}, busStation{"Administration"}, busStation{"IT"}, busStation{"Liberal Arts"},busStation{"Aggricultural"},busStation{"Train Station"}}
	passengerGroup := []*passenger{}
	for i := 1; i < rand.Intn(20); i++ {
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
	return &passengerGroup
}

func newPassenger() {
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
	return &passengerGroup
}

