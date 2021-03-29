package rs

// Event Class End --------------------------------------------
func ClassEnd(lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(150, 200), p)
	//rs.GnrTrf(CarGroup())

}

// Event train ----------------------------------------
func Train(lst []*BusStop, p *Passenger) {
	GnrPsgAt(lst, "hBuilding", Random(50, 100), p)
	//rs.GnrTrf(CarGroupTr())
}

// // Event After 4pm ---------------------------------------------
func AfterWork(lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(350, 500), p)
	//rs.GnrTrf(CarGroupBusy())
}
