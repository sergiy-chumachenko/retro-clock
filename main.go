package main

import (
	"fmt"
	"github.com/inancgumus/screen"
	"time"
)

func main() {
	screen.Clear()
	for {
		screen.MoveTopLeft()
		currentTime := time.Now()
		hours, minutes, seconds, microseconds := currentTime.Hour(), currentTime.Minute(), currentTime.Second(), currentTime.Nanosecond()
		clock := [...]placeholder{
			digits[hours/10], digits[hours%10],
			sep,
			digits[minutes/10], digits[minutes%10],
			sep,
			digits[seconds/10], digits[seconds%10],
			dot,
			digits[microseconds/10e7],
		}

		alarmed := seconds%10 == 0

		for line := range clock[0] {
			if alarmed {
				clock = alarm
			}
			for index, digit := range clock {
				next := clock[index][line]
				if seconds%2 == 0 && (digit == sep || digit == dot) {
					next = "   "
				}
				fmt.Printf("%s  ", next)
			}
			fmt.Printf("\n")
		}
		time.Sleep(time.Nanosecond)
	}
}
