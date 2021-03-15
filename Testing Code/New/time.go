package main

import (
	"fmt"
	"os"
)

var (
	globalHour   int = 0
	globalMinute int = 0
	globalDay    int = 0
)

func main() {
	fmt.Printf("Input the starting hour(0-24): ")
	fmt.Scanln(&globalHour)
	fmt.Printf("%v Days -> %v:%v\n", globalDay, globalHour, globalMinute)
	for {
		timeTick()
		fmt.Printf("%v Days -> %v:%v\n", globalDay, globalHour, globalMinute)
	}
}

func timeTick() {
	globalMinute++
	if globalMinute == 60 {
		globalHour++
		globalMinute = 0
	}
	if globalHour == 24 {
		globalDay++
		globalHour = 0
	}
	if globalDay == 8 {
		os.Exit(0)
	}
}
