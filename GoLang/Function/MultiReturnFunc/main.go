package main

import "fmt"

func main() {
	number := []int{2, 5, 6, 7, 3}
	fmt.Println(findMinMax(number))
}

//Returning multiple values will need to be around () brackets.
//And return value may contain (int, int) or declared it immediately by (min int, max int)
func findMinMax(num []int)  (min int, max int) {
	max = num[0]
	min = num[0]

	for _, n := range num {
		if(n < min) {
			min = n
		}
		if(n > max) {
			max = n
		}
	}
	return min, max
}