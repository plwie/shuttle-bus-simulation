package rs

import (
	"time"
)

// TimeTick add 1 min to the global clock
func TimeTick(hour *int, min *int) {
	*min++
	if *min >= 60 {
		*hour += *min / 60
		*min -= 60 * (*min / 60)
	}
}

// ConTimeTick constantly call Time Tick
func ConTimeTick(hour *int, min *int) {
	for *hour < 24 {
		TimeTick(hour, min)
		time.Sleep(time.Millisecond * 50)
	}
}
