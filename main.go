package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func timeToDigits(currentTime time.Time) []int {
	convertedToString := strconv.Itoa(currentTime.Hour()) + strconv.Itoa(currentTime.Minute()) + strconv.Itoa(currentTime.Second())
	split := strings.Split(convertedToString, "")
	return convertToDigits(split)
}

func convertToDigits(strArr []string) []int {
	var intArr []int
	for _, v := range strArr {
		j, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		intArr = append(intArr, j)

	}
	return intArr
}

func drawTime(convertedTime []int, separator [5]string, digits [10][5]string){
	for i:=0;i<5;i++{
		fmt.Printf("%s %s %s %s %s %s %s %s\n",
			digits[convertedTime[0]][i], digits[convertedTime[1]][i], separator[i],
			digits[convertedTime[2]][i], digits[convertedTime[3]][i], separator[i],
			digits[convertedTime[4]][i], digits[convertedTime[5]][i],
		)
	}
}

func main(){
	digits := [10][5]string{
		{
			"███",
			"█ █",
			"█ █",
			"█ █",
			"███",
		},
		{
			"██ ",
			" █ ",
			" █ ",
			" █ ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"█  ",
			"███",
		},
		{
			"███",
			"  █",
			"███",
			"  █",
			"███",
		},
		{
			"█ █",
			"█ █",
			"███",
			"  █",
			"  █",
		},
		{
			"███",
			"█  ",
			"███",
			"  █",
			"███",
		},
		{
			"███",
			"█  ",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"  █",
			"  █",
			"  █",
			"  █",
		},
		{
			"███",
			"█ █",
			"███",
			"█ █",
			"███",
		},
		{
			"███",
			"█ █",
			"███",
			"  █",
			"███",
		},
	}

	separator := [5]string{
		"   ",
		" ░ ",
		"   ",
		" ░ ",
		"   ",
	}
	separator2 := [5]string{
		"   ",
		"   ",
		"   ",
		"   ",
		"   ",
	}

	flag := true
	for {
		converted := timeToDigits(time.Now())
		var sep [5]string
		if flag {
			sep = separator
			flag = false
		} else {
			sep = separator2
			flag = true
		}
		time.Sleep(time.Second)
		drawTime(converted, sep, digits)
		fmt.Printf("\n\n\n\n\n\n")
	}
}
