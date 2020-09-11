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
		hours, minutes, seconds := currentTime.Hour(), currentTime.Minute(), currentTime.Second()

		clock := [...]placeholder{
			digits[hours/10], digits[hours%10],
			sep,
			digits[minutes/10], digits[minutes%10],
			sep,
			digits[seconds/10], digits[seconds%10],
		}
		for line := range clock[0] {
			for index, digit := range clock {
				next := clock[index][line]
				if digit == sep && seconds%2 == 0 {
					next = "   "
				}
				fmt.Printf("%s  ", next)
			}
			fmt.Printf("\n")
		}
		time.Sleep(time.Second)
	}

}
