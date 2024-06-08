package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Array: ", array)
	fmt.Println("Genap: ", filter(array, isEven))
	fmt.Println("Ganjil: ", filter(array, isOdd))
}

func filter(num []int, function func(int)bool ) []int {
	result := []int{}
	for _, p := range num {
		if(function(p)) {
			result = append(result, p)
		}
	}
	return result
}

func isEven(num int) bool{
	return num % 2 == 0
}

func isOdd(num int) bool {
	return num % 2 != 0
}