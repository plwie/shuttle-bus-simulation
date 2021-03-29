package rs

// ClassEnd --------------------------------------------
func ClassEnd(lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(150, 200), p)
	//rs.GnrTrf(CarGroup())

}

// Train ----------------------------------------
func Train(lst []*BusStop, p *Passenger) {
	GnrPsgAt(lst, "hBuilding", Random(50, 100), p)
	//rs.GnrTrf(CarGroupTr())
}

// // AfterWork ---------------------------------------------
func AfterWork(lst []*BusStop, p *Passenger) {
	GnrPsg(lst, Random(350, 500), p)
	//rs.GnrTrf(CarGroupBusy())
}
