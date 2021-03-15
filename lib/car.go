package rs

import (
	"math/rand"
	"time"
)

// Car create a car object
type Car struct {
	id string
}

//CarGroup create group of cars
func CarGroup() []*Car {
	CarGroup := []*Car{}
	rand.Seed(time.Now().UnixNano())
	random1 := random(15, 25)
	for i := 1; i < random1; i++ {
		var p *Car
		p = new(Car)
		CarGroup = append(CarGroup, p)
	}
	return CarGroup
}
