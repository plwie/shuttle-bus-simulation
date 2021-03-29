package rs

import (
	"math/rand"
	"time"
)

// Car create a car object
type Car struct {
	id string
}

//CarGroupTr create group of cars in train arrival event
func CarGroupTr() []*Car {
	CarGroup := []*Car{}
	rand.Seed(time.Now().UnixNano())
	random1 := random(25, 30)
	for i := 0; i < random1; i++ {
		var p *Car
		p = new(Car)
		CarGroup = append(CarGroup, p)
	}
	return CarGroup
}

//CarGroup represent typical car generates
func CarGroup() []*Car {
	CarGroup := []*Car{}
	rand.Seed(time.Now().UnixNano())
	random1 := random(15, 25)
	for i := 0; i < random1; i++ {
		var p *Car
		p = new(Car)
		CarGroup = append(CarGroup, p)
	}
	return CarGroup
}

//CarGroupBusy create group of cars in busy event
func CarGroupBusy() []*Car {
	CarGroup := []*Car{}
	rand.Seed(time.Now().UnixNano())
	random1 := random(30, 35)
	for i := 0; i < random1; i++ {
		var p *Car
		p = new(Car)
		CarGroup = append(CarGroup, p)
	}
	return CarGroup
}
