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
	waitingTime int
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
	rand.Seed(time.Now().UnixNano())
	random1 := random(3000, 10000)
	for i := 1; i < random1; i++ {
		var p *Passenger
		p = new(Passenger)
		rand.Seed(time.Now().UnixNano())
		rando1 := random(1, 5)
		p.Source = stopList[rando1].Name
		rand.Seed(time.Now().UnixNano())
		rando2 := random(1, 5)
		p.Destination = stopList[rando2].Name
		if p.Destination == p.Source {
			continue
		} else if p.Destination != p.Source {
			passengerGroup = append(passengerGroup, p)
		}
	}
	return passengerGroup
}

//NewPassenger1 Create a new passenger
func NewPassenger1(stopList []*BusStop) *Passenger {
	var p *Passenger
	p = new(Passenger)
	p.Source = "a"
	p.Destination = "b"
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
