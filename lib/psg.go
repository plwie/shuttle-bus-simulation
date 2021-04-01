package rs

import (
	"math/rand"
)

var M = make(map[string]int)

// Passenger create a passenger object
type Passenger struct {
	Source      string
	Destination string
	Next        *Passenger
	WaitTime    int
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

//NewPassenger1 Create a new passenger
func NewPassenger() *Passenger {
	var p Passenger
	return &p
}

//GnrPsg Generate psg and add to bus stop
func GnrPsg(stopList []*BusStop, random1 int, psgr *Passenger) {
	for i := 0; i < random1; i++ {
		psgr.WaitTime = 0
		src := Random(0, len(stopList))
		psgr.Source = stopList[src].Name
		psgr.Destination = stopList[Random(0, len(stopList))].Name
		if psgr.Source == psgr.Destination {
			i--
			continue
		} else {
			stopList[src].Q.Add(*psgr)
		}
	}
}

//GnrPsgAt Generate psg and add to specific bus stop
func GnrPsgAt(stopList []*BusStop, stop string, inputPsg int, psgr *Passenger) {
	var temp int
	for j := 0; j < len(stopList); j++ {
		if stop == stopList[j].Name {
			temp = j
		}
	}
	for i := 0; i < inputPsg; i++ {
		psgr.WaitTime = 0
		psgr.Source = stop
		psgr.Destination = stopList[Random(0, len(stopList))].Name
		if psgr.Source == psgr.Destination {
			i--
			continue
		} else {
			stopList[temp].Q.Add(*psgr)
		}
	}
}

//Random number of int
func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// Random code from https://gist.github.com/201Factory/5ef7c2d46cf848db16041cafa17ab054
