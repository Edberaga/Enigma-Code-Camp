package main

import "fmt"

func main() {
	var array = [...]string{ "Warrior", "Mage", "Paladin", "Thief"}

	for i := 0; i < len(array); i++ {
		fmt.Printf("Party %d : %s\n", i + 1, array[i] )
	}
	fmt.Println("========================================================================")

	for i, value := range array {
		fmt.Printf("Party %d : %s\n", i + 1, value )
	}
	fmt.Println("========================================================================")

	for _, value := range array {
		fmt.Printf("We have %s in our party\n", value)
	}
	fmt.Println("========================================================================")

	for i := range array {
		fmt.Printf("Party Member: %d\n", i)
	}
	fmt.Println("========================================================================")

	var numbers = [5]int{5, 8, 11, 14, 7}

	for _, value := range numbers {
		if(value %2 == 0) {
			fmt.Println(value)
		}
	}
	fmt.Println("========================================================================")

	fmt.Println("Sebelum: ", numbers)
	for i := range numbers {
		numbers[i] *= 2
	}
	fmt.Println("Sesudah: ", numbers)
}