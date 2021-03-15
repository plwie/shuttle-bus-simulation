package main

import (
	"fmt"
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
		timeTick(&globalHour, &globalMinute)
		fmt.Printf("%v Days -> %v:%v\n", globalDay, globalHour, globalMinute)
	}
}

func timeTick(hour *int, min *int) {
	*min++
	if *min >= 60 {
		*hour += *min / 60
		*min -= 60 * (*min / 60)
	}
}
