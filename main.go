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
	converted := timeToDigits(time.Now())
    for i:=0;i<5;i++{
		fmt.Printf("%s %s %s %s %s %s %s %s\n",
			digits[converted[0]][i], digits[converted[1]][i], separator[i],
			digits[converted[2]][i], digits[converted[3]][i], separator[i],
			digits[converted[4]][i], digits[converted[5]][i],
		)
	}
}
