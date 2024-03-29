package rs

import (
	"math/rand"
)

// Passenger create a passenger object
type Passenger struct {
	Source      string
	Destination string
	Next        *Passenger
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
		psgr.Source = stopList[rand.Intn((len(stopList)-0-1)+1)].Name
		for j := 0; j < len(stopList); j++ {
			if psgr.Source == stopList[j].Name {
				psgr.Destination = stopList[rand.Intn((len(stopList)-0-1)+1)].Name
				for psgr.Source == psgr.Destination {
					psgr.Destination = stopList[rand.Intn((len(stopList)-0-1)+1)].Name
				}
				stopList[j].Q.Add(*psgr)
			}
		}
	}
}

//GnrPsgAt Generate psg and add to specific bus stop
func GnrPsgAt(stopList []*BusStop, stop string, inputPsg int, psgr *Passenger) {
	for i := 0; i < len(stopList); i++ {
		if stop == stopList[i].Name {
			psgr.Source = stop
			for j := 0; j < inputPsg; j++ {
				psgr.Destination = stopList[rand.Intn((len(stopList)-0-1)+1)].Name
				for psgr.Source == psgr.Destination {
					psgr.Destination = stopList[rand.Intn((len(stopList)-0-1)+1)].Name
				}
				stopList[i].Q.Add(*psgr)
				// fmt.Println(stopList[i].Name)
				// fmt.Println(stopList[i].Q.Size)
			}
		}
	}
}

//Random number of int
func Random(min int, max int) int {
	return rand.Intn(max-min) + min
}

// Random code from https://gist.github.com/201Factory/5ef7c2d46cf848db16041cafa17ab054
