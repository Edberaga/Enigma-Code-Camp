package main

import (
	"golng-introduction/AccessModifier/Library"
	"fmt"
)

func main() {
	fmt.Println("Each hour in a day: ",library.DayInHour)
	//fmt.Println(library.secondInMinute) Error: unable to access if their name low case that means they aren't exported in package
	//fmt.Println(library.minuteInHour)

	fmt.Println(library.MinuteToHour(180))
	//fmt.Println(library.secondToMinute(180))
}