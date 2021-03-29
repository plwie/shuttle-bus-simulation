package rs

// TimeTick add 1 min to the global clock
func TimeTick(hour *int, min *int) {
	*min++
	if *min >= 60 {
		*hour += *min / 60
		*min -= 60 * (*min / 60)
	}
	// fmt.Printf("%02v:%02v\n", *hour, *min)
}

// ConTimeTick constantly call Time Tick
func ConTimeTick(hour *int, min *int, lst []*BusStop, p *Passenger) int {
	eventCalled := 0
	initTime := 0
	for *hour < 24 {
		// fmt.Printf("%02v:%02v\n", *hour, *min)
		TimeTick(hour, min)
		time.Sleep(time.Millisecond * 30)
		// Event Class End --------------------------------------------
		if initTime != *hour {
			GnrPsg(lst, Random(150, 200), p)
			eventCalled++
			initTime++
			//rs.GnrTrf(CarGroup())
		}

		// Event train ----------------------------------------
		if *hour%2 == 0 && *min == 0 {
			GnrPsgAt(lst, "hBuilding", Random(50, 100), p)
			eventCalled++
			//rs.GnrTrf(CarGroupTr())
		}

		// // Event After 4pm ---------------------------------------------
		if (*hour == 16 || *hour == 17 || *hour == 18) && *min == 0 {
			GnrPsg(lst, Random(150, 200), p)
			eventCalled++
			//rs.GnrTrf(CarGroupBusy())
		}
	}
	return eventCalled
}
