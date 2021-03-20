package rs

import (
	"fmt"
	"time"
)

// TimeTick add 1 min to the global clock
func TimeTick(hour *int, min *int) {
	*min++
	if *min >= 60 {
		*hour += *min / 60
		*min -= 60 * (*min / 60)
	}
	fmt.Printf("%02v:%02v\n", *hour, *min)
}

// ConTimeTick constantly call Time Tick
func ConTimeTick(hour *int, min *int) {
	for *hour < 24 {
		TimeTick(hour, min)
		time.Sleep(time.Millisecond * 50)
	}
}
