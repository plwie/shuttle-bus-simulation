package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	mainInput string
	mainCmd   []string
	err       error
)

func testA() {
	fmt.Println("Succesfully called function A...")
}

func testD() {
	fmt.Println("Succesfully called function D...")
}

func main() {
	// Setup
	reader := bufio.NewReader(os.Stdin)
	// text = strings.Replace(text, "\n", "", -1)

	var cmdMap = map[string]func(){
		"FunctionA": testA,
		"FunctionD": testD,
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
