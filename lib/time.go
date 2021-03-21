package rs

<<<<<<< HEAD
import (
	"time"
)

=======
>>>>>>> e2e087412e6ad4783e18e7d55c90b2fb51432c96
// TimeTick add 1 min to the global clock
func TimeTick(hour *int, min *int) {
	*min++
	if *min >= 60 {
		*hour += *min / 60
		*min -= 60 * (*min / 60)
	}
<<<<<<< HEAD
	// fmt.Printf("%02v:%02v\n", *hour, *min)
=======
>>>>>>> e2e087412e6ad4783e18e7d55c90b2fb51432c96
}

// ConTimeTick constantly call Time Tick
func ConTimeTick(hour *int, min *int) {
	for *hour < 24 {
		TimeTick(hour, min)
		// time.Sleep(time.Millisecond * 50)
	}
}
