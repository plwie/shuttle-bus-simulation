package rs

import (
	"math/rand"
)

// Passenger create a passenger object
type Passenger struct {
	Source      string
	Destination string
	next        *Passenger
	waitingTime int
}

func random(min int, max int) int {
	return rand.Intn(max-min) + min
}

//NewPassenger1 Create a new passenger
func NewPassenger() *Passenger {
	var p *Passenger
	p = new(Passenger)
	return p
}

//GnrPsg Generate psg and add to bus stop
func GnrPsg(stopList []*BusStop, random1 int, psgr *Passenger) {
	for i := 1; i < random1; i++ {
		psgr.Source = *&stopList[rand.Intn(10)].Name
		psgr.Destination = *&stopList[rand.Intn((len(stopList)-0-1)+1)].Name
		for i := 0; i < len(stopList)-1; i++ {
			if psgr.Source == *&stopList[i].Name {
				stopList[i].Q.Add(*psgr)
				// fmt.Println(stopList[i].Name)
				// fmt.Println(stopList[i].Q.Size)
			}
		}
	}
}

//GnrPsgAt Generate psg and add to specific bus stop
func GnrPsgAt(stopList []*BusStop, stop string, inputPsg int, psgr *Passenger) {
	for i := 1; i < inputPsg; i++ {
		for i := 0; i < len(stopList)-1; i++ {
			if stop == *&stopList[i].Name {
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
