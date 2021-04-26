if prevCount != BusArr.PassOn {
	fmt.Println("Bus", name, " pickup or drop off on the road")
} else {
	if BusArr.DistToNext <= 0 {
		correctness3++
		fmt.Println("Bus", name, "status", "Dropping/Get Passenger")
	} else {
		correctness3++
		fmt.Println("Bus", name, "status", "Travelling")
		// stepCor++
	}
}