package main

import "fmt"

func main() {
	total := sum(1, 2, 3, 4)
	fmt.Println(total)
}

func sum(num ...int) int{
	total := 0
	fmt.Printf("%T \n", num) //the num is a slice type []int
	for _, number := range num {
		total += number
	}

	return total
}