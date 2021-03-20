package main

import (
	"bufio"
	"fmt"
	"os"
	rs "rs/lib"
	"strconv"
	"strings"
)

var (
	mainInput string
	mainCmd   []string
	err       error
	tStopLst  []*rs.BusStop
	tMin      int
	tHr       int
)

func bstList() bool {
	// This function should print out the list of created bus stop
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

func bstCreate() bool {
	// This function should create a new bus stop and add it to the list
	if len(mainCmd) < 2 {
		fmt.Println("Error: invalid parameter")
		return false
	}
	tStopLst = append(tStopLst, &rs.BusStop{Name: mainCmd[1]})
	return true
}

func bstAdd() bool {
	// This function should add psg into queue of target bus stop
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
	p := rs.NewPassenger()
	for i := 0; i < psgNum; i++ {
		rs.GnrPsgAt(tStopLst, mainCmd[1], psgNum, p)
	}

	// Check feedback
	fmt.Println("Added successsfully")
	return true
}

func bstRmv() bool {
	// This function should remove psg out of queue of target bus stop
	return true
}

func tick() bool {
	fmt.Println("Starting the clock...")
	tMin = 0
	tHr = 0
	for tHr < 24 {
		fmt.Printf("%02v:%02v\n", tHr, tMin)
		rs.TimeTick(&tHr, &tMin)
	}
	fmt.Printf("%02v:%02v\n", tHr, tMin)
	fmt.Println("Clock check successful")
	return true
}

func help() bool {
	fmt.Println("help	-> Display the list of available commands")
	fmt.Println("bstList	-> Print out the list of created bus stop")
	fmt.Println("bstCreate name	-> Create a new bus stop with a name")
	fmt.Println("bstAdd tgStop psgNum	-> Add psgNum passengers into the tgStop bus stop")
	fmt.Println("bstRmv tgStop psgNum 	-> Remove psgNum passengers into the tgStop bus stop")
	fmt.Println("timeTick	-> Run 1 day of time tick and print the clock")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	return true
}

func main() {
	// Setup
	reader := bufio.NewReader(os.Stdin)
	// text = strings.Replace(text, "\n", "", -1)

	var cmdMap = map[string](func() bool){
		"bstList":   bstList,
		"bstCreate": bstCreate,
		"bstAdd":    bstAdd,
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
