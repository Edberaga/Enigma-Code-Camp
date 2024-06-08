package main

import "fmt"

func main() {
	var score int

	fmt.Println("Enter the student score: ")
	fmt.Scanln(&score)
	if score == 100 {
		fmt.Println("PERFECT SCORE!")
	} else if score <= 99 && score >=80 {
		fmt.Println("MEMUASKAN!")
	} else if score <= 79 && score >= 50 {
		fmt.Println("CUKUP MEMUASKAN")
	} else {
		fmt.Println("KURANG MEMUASKAN")
	}
}