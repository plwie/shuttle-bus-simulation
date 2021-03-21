package main

import (
	"fmt"
	"math"
	"math/rand"
	rs "rs/lib"
	"time"
)

var (
	stopList   []*rs.BusStop
	inputNoBus int
	globalHour int
	globalMin  int
	//for putting busc in main
	countPos  int     = 0
	count     int     = 0
	graph             = rs.Graph{}
	totalTime float64 = 0
	passTotal int     = 0
	waitingTime float64 = 0
)

// Bus Struct
// type Bus struct {
// 	availSeats int
// 	passOn     int
// 	currStop   string
// 	nextStop   string
// }

//busc threading function---------------------------------------------------------------
func Busc(name string, path []*rs.BusStop) {
	//need to declare global count = 0
	// graph := rs.Graph{}
	m := make(map[string]int)
	pos := countPos
	countPos++
	var len int = len(path)
	var count int = 0
	var countPass int = 0
	var localTimeHour int = 0
	var localTimeMin int = 0
	var spd float64
	var dist float64
	var calcTime float64

	//create bus struct instance
	busStruct := rs.Bus{
		AvailSeats: 30,
		PassOn:     0,
		CurrStop:   *&path[pos].Name,
		NextStop:   *&path[pos+1].Name,
	}
	for i := 0; i < 10; i++ {
		m[path[i].Name] = 0
	}
	//code for bus traveling (busstop to another busstop)
	for {
		if pos < len && name != "test" {
			time.Sleep(time.Millisecond * 1000)
			busStruct.CurrStop = *&path[pos].Name
			busStruct.NextStop = *&path[(pos+1)%len].Name

			busStruct.PassOn -= m[busStruct.CurrStop]
			busStruct.AvailSeats += m[busStruct.CurrStop]
			fmt.Println("Passenger of", name, "off at", busStruct.CurrStop, "is:", m[busStruct.CurrStop])
			m[busStruct.CurrStop] = 0

			fmt.Println(count, name, busStruct.CurrStop, busStruct.NextStop, busStruct.AvailSeats, busStruct.PassOn)
			// fmt.Println(globalHour, globalMin)

			fmt.Println("G:H", globalHour, "G:M", globalMin)
			fmt.Println("L:H", localTimeHour, "L:M", localTimeMin)

			if localTimeHour <= globalHour && localTimeMin <= globalMin {
				spd = float64(graph.GetSpeed(path[pos], path[(pos+1)%len]))
				dist = float64(graph.Edges[pos].Cost)
				calcTime = float64(math.Round(((dist/spd)*3600)*100) / 100)
				for i := 0; i < busStruct.AvailSeats; i++ {
					if path[i%10].Q.Size != 0 {
						// m[path[i%10].Q.Pop().Destination]++
						rs.GetPass(m, path, i)
						busStruct.PassOn++
						countPass++
						busStruct.AvailSeats--

					}
				}

				fmt.Println(m)

				if localTimeMin <= 60 {
					localTimeMin = globalMin + (int(calcTime) / 60)
				}
				if localTimeMin >= 60 {
					localTimeMin = localTimeMin - 60*(localTimeMin/60)
					localTimeHour++
				}
				totalTime += (calcTime * float64(countPass))
				pos++
				count++
			}
			fmt.Println("|distance:", dist, "|speed:", spd, "|time:", calcTime, "sec", "|totalTime:", totalTime)
			passTotal += countPass
			fmt.Println("|countpass", countPass, "|passTotal", passTotal, "totaltime: ", totalTime)
			// pos++
			// count++
			countPass = 0
			} else {
				pos = 0
			}
			if globalHour == 1{
				waitingTime = ((totalTime)/float64(passTotal))/60
				secc := math.Round((((math.Mod(waitingTime,1))*60)*1000)/1000)
				minn := (math.Floor(waitingTime/1))
				fmt.Println("Waiting Time:", minn, "minutes", secc, "secs")
				break
			}
			}
			}

//End busc--------------------------------------------------------------------------------------------------------

func main() {

	//WaitingTime
	//waitingTime := totalTime/float64(passTotal)

	// graph := rs.Graph{}

	aBuilding := rs.BusStop{Name: "aBuilding"}
	bBuilding := rs.BusStop{Name: "bBuilding"}
	cBuilding := rs.BusStop{Name: "cBuilding"}
	dBuilding := rs.BusStop{Name: "dBuilding"}
	eBuilding := rs.BusStop{Name: "eBuilding"}
	fBuilding := rs.BusStop{Name: "fBuilding"}
	gBuilding := rs.BusStop{Name: "gBuilding"}
	hBuilding := rs.BusStop{Name: "hBuilding"}
	iBuilding := rs.BusStop{Name: "iBuilding"}
	jBuilding := rs.BusStop{Name: "jBuilding"}

	stopList := graph.StopList

	stopList = append(stopList, &aBuilding)
	stopList = append(stopList, &bBuilding)
	stopList = append(stopList, &cBuilding)
	stopList = append(stopList, &dBuilding)
	stopList = append(stopList, &eBuilding)
	stopList = append(stopList, &fBuilding)
	stopList = append(stopList, &gBuilding)
	stopList = append(stopList, &hBuilding)
	stopList = append(stopList, &iBuilding)
	stopList = append(stopList, &jBuilding)

	graph.AddEdge(&aBuilding, &bBuilding, 1)
	graph.AddEdge(&bBuilding, &cBuilding, 2)
	graph.AddEdge(&cBuilding, &dBuilding, 1)
	graph.AddEdge(&dBuilding, &eBuilding, 2)
	graph.AddEdge(&eBuilding, &fBuilding, 3)
	graph.AddEdge(&fBuilding, &gBuilding, 1)
	graph.AddEdge(&gBuilding, &hBuilding, 2)
	graph.AddEdge(&hBuilding, &iBuilding, 1)
	graph.AddEdge(&iBuilding, &jBuilding, 3)
	graph.AddEdge(&jBuilding, &aBuilding, 2)

	//---------------------------------------------------------------
	graph.GenerateTraffic(rs.CarGroup(), &aBuilding, &bBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &bBuilding, &aBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &bBuilding, &cBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &cBuilding, &bBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &cBuilding, &dBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &dBuilding, &cBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &dBuilding, &eBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &eBuilding, &dBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &eBuilding, &fBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &fBuilding, &eBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &fBuilding, &gBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &gBuilding, &fBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &gBuilding, &hBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &hBuilding, &gBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &hBuilding, &aBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &iBuilding, &hBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &iBuilding, &jBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &jBuilding, &iBuilding)
	graph.GenerateTraffic(rs.CarGroup(), &jBuilding, &aBuilding)
	//-----------------------------------------------------------------

	fmt.Printf("Initiated bus stop list: %v\n", stopList)
	fmt.Println("How many bus?")
	fmt.Scanln(&inputNoBus)
	var inputPsg int
	fmt.Println("How many passenger?")
	fmt.Scanln(&inputPsg)

	psgr := rs.NewPassenger()
	rand.Seed(time.Now().UnixNano())
	//Passenger Generated -------------------------
	// random1 := rs.Random(150, 200)
	// rand.Seed(time.Now().UnixNano())
	// random2 := rs.Random(50, 100)
	// rand.Seed(time.Now().UnixNano())
	// random3 := rs.Random(150, 200)

	//Cars Generated ------------------------------
	// cars1 := rs.CarGroup()

	// rs.TimeTick(globalHour, globalMin)

	// Init -------------------------------------------------
	fmt.Println("Total Passenger :", inputPsg)
	rs.GnrPsg(stopList, inputPsg, psgr)
	// // Event Class End --------------------------------------------
	// if (globalMin % 60) == 0 {
	// 	rs.GnrPsg(stopList, inputPsg, psgr)
	// }

	// Event train ----------------------------------------
	// if (globalMin % 120) == 0 {
	// 	rs.GnrPsgAt(stopList, "hBuilding", random2, psgr)
	// 	cars := rs.CarGroup()
	// 	fmt.Println("Total cars")
	// 	fmt.Println(len(cars))
	// }

	// // Event After 4pm ---------------------------------------------
	// if globalHour == 16 {
	// 	rs.GnrPsg(stopList, random3, psgr)
	// }

	fmt.Println("#,BusName,CurrentStop,NextStop,AvailableSeats,TotalPassengerOnBus ")
	for i := 0; i < inputNoBus; i++ {
		go Busc("bus"+fmt.Sprint((i+1)), stopList)
	}
	go rs.ConTimeTick(&globalHour, &globalMin)
	Busc("test", stopList)
	fmt.Println("Ending main package...")
}
