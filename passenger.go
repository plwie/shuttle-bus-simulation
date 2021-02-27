package main

import (
	"math/rand"
)

type passenger struct {
	source      string
	destination string
}

func newPassenger(BusStation) {
	allStation := [7]string{"Engineering", "Science", "Administration", "IT", "Liberal Arts", "Aggricultural", "Train Station"}
	passengerGroup := []*passenger{}
	for i := 1; i < 100; i++ {
		var p *passenger
		p = new(passenger)
		p.source = rand.Intn(len(BusStation))
		p.destination = rand.Intn(len(allStation))
		if p.destination == p.source {
			continue
		} else if p.destination != p.source {
			passengerGroup = append(passengerGroup, passenger)
		}
	}
	return &passengerGroup
}

func main() {

}
