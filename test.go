package main

import (
	"bufio"
	"fmt"
	"os"
	rs "rs/lib"
	"strconv"
	"strings"
	"time"
)

var (
	mainInput string
	mainCmd   []string
	err       error
	tStopLst  []*rs.BusStop
	tMin      int
	tHr       int
)

// This function should print out the list of created bus stop
func bstList() bool {
	if len(tStopLst) == 0 {
		fmt.Println("Error: no item in bus stop list")
		return false
	}
	fmt.Printf("List of bus stops: ")
	for _, v := range tStopLst {
		fmt.Printf("%v ", v.Name)
	}
	fmt.Printf("\n")
	return true
}

// This function should print out the data of the target bus stop
func bstGet() bool {
	// Check parameters
	if len(mainCmd) < 2 {
		fmt.Println("Error: invalid parameter")
		return false
	}

	// Find and print
	for _, v := range tStopLst {
		if v.Name == mainCmd[1] {
			temp := *v
			fmt.Printf("Bus Stop Name: %v\n", temp.Name)
			fmt.Printf("Total Passengers: %v\n", temp.Q.Size)
			fmt.Printf("List of passengers: ")
			for i := temp.Q.Head; i != nil; i = i.Next {
				fmt.Printf("%v ", i)
			}
			fmt.Printf("\n")
			return true
		}
	}
	return false
}

// This function should create a new bus stop and add it to the list
func bstCreate() bool {
	// Check parameters
	if len(mainCmd) < 2 {
		fmt.Println("Error: invalid parameter")
		return false
	}
	tStopLst = append(tStopLst, &rs.BusStop{Name: mainCmd[1]})
	return true
}

// This function should add psg into queue of target bus stop
func bstAdd() bool {
	// Check parameters
	if len(mainCmd) < 3 {
		fmt.Println("Error: invalid parameter")
		return false
	}
	// Convert second parameter into int
	psgNum, err := strconv.Atoi(mainCmd[2])
	if err != nil {
		fmt.Printf("Error: invalid argument %v\n", mainCmd[2])
		return false
	}
	// Create a new passenger
	start := time.Now()
	p := rs.NewPassenger()
	for i := 0; i < psgNum; i++ {
		rs.GnrPsgAt(tStopLst, mainCmd[1], psgNum, p)
	}
	end := time.Since(start)

	// Check feedback
	// Amount of psg in bst src and dst dif
	totalCorrect := 0
	for _, v := range tStopLst {
		if v.Name == mainCmd[1] {
			temp := *v
			for temp.Q.Size != 0 {
				check := temp.Q.Pop()
				if check.Source != check.Destination {
					totalCorrect++
				}
			}
			fmt.Printf("Added %v passengers successsfully\n", totalCorrect)
			fmt.Printf("Time taken: %v\n", end)
			return true
		}
	}
	fmt.Println("Error: adding failed")
	return false
}

// This function should add psg into queue of random bus stop
func bstAddRd() bool {
	// Check parameters
	if len(mainCmd) < 3 {
		fmt.Println("Error: invalid parameter")
		return false
	}
	// Convert second parameter into int
	psgNum, err := strconv.Atoi(mainCmd[2])
	if err != nil {
		fmt.Printf("Error: invalid argument %v\n", mainCmd[2])
		return false
	}
	// Create a new passenger
	start := time.Now()
	p := rs.NewPassenger()
	for i := 0; i < psgNum; i++ {
		rs.GnrPsg(tStopLst, psgNum, p)
	}
	end := time.Since(start)

	// Check feedback
	fmt.Println("Added successsfully")
	fmt.Printf("Time taken: %v\n", end)
	return true
}

// This function should remove psg out of queue of target bus stop
func bstRmv() bool {
	return true
}

func tick() bool {
	start := time.Now()
	fmt.Println("Starting the clock...")
	tHr = 0
	tMin = 0
	// fmt.Printf("%02v:%02v\n", tHr, tMin)
	rs.ConTimeTick(&tHr, &tMin)
	if tHr == 24 && tMin == 0 {
		end := time.Since(start)
		fmt.Println("Clock check successful")
		fmt.Printf("Time taken: %v\n", end)
		return true
	} else {
		fmt.Println("Error: clock check failed")
		return false
	}
}

func help() bool {
	fmt.Println("help	-> Display the list of available commands")
	fmt.Println("bstList	-> Print out the list of created bus stop")
	fmt.Println("bstGet tgStop	-> Get data of the tgStop bus stop")
	fmt.Println("bstCreate name	-> Create a new bus stop with a name")
	fmt.Println("bstAdd tgStop psgNum	-> Add psgNum passengers into the tgStop bus stop")
	fmt.Println("bstAddRd tgStop psgNum	-> Add psgNum passengers into the tgStop bus stop")
	fmt.Println("bstRmv tgStop psgNum 	-> Remove psgNum passengers into the tgStop bus stop")
	fmt.Println("timeTick	-> Run 1 day of time tick and print the clock")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	return true
}

func main() {
	// Setup
	tStopLst = append(tStopLst, &rs.BusStop{Name: "cheatBst"})
	p := rs.NewPassenger()
	rs.GnrPsgAt(tStopLst, "cheatBst", 100, p)
	/*
		for i := 0; i < 100; i++ {
			tStopLst[0].Q.Add(p)
		}
	*/
	reader := bufio.NewReader(os.Stdin)
	// text = strings.Replace(text, "\n", "", -1)

	var cmdMap = map[string](func() bool){
		"bstList":   bstList,
		"bstGet":    bstGet,
		"bstCreate": bstCreate,
		"bstAdd":    bstAdd,
		"bstAddRd":  bstAddRd,
		"bstRmv":    bstRmv,
		"timeTick":  tick,
		"help":      help,
	}

	fmt.Println(cmdMap)
	fmt.Println("Test Drive Initiated...!")

	// Simple shell
	for {
		// Take Input and Split Argument
		fmt.Printf("> ")
		mainInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error: input unsuccessful")
			continue
		}
		mainInput = strings.TrimSpace(mainInput)
		mainCmd = strings.Split(mainInput, " ")
		exeTarget, ok := cmdMap[mainCmd[0]]
		if !ok {
			fmt.Println("Error: command does not exist; Try help for list of commands")
			continue
		}
		exeTarget()
	}
}
