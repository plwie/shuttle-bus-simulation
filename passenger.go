package main

import (
	"fmt"
	"math/rand"
)

allStation := [7]string{"Engineering", "Science", "Administration", "IT", "Liberal Arts", "Aggricultural","Train Station"}

type passenger struct {
	source      string
	destination string
}

func newPassenger(BusStation, allStation) {
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
	randomStation := rand.Intn(len(allStation))
	p := passenger(BusStation, randomStation)
}

func main() {

}
