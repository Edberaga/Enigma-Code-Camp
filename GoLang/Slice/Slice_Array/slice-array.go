package main

import "fmt"

func main() {
	//Array is based on value, while Slice is based on references
	
	//We will try array Pass by Value
	numbers := [5]int{2, 1, 4, 8, 7}
	anotherNum := numbers
	fmt.Println("Array Values: ", numbers, "\nDuplicate Values: ", anotherNum)

	anotherNum[3] = 16
	fmt.Println("Array Values: ", numbers, "\nDuplicate Values: ", anotherNum)
	fmt.Println("========================================================================")

	//Now try to do it without set the length, which going to be Pass by reference
	prices := []int{2, 1, 4, 8, 7}
	anotherPrices := prices
	fmt.Println("Array Values: ", prices, "\nDuplicate Values: ", anotherPrices)

	anotherPrices[3] = 16
	fmt.Println("Array Values: ", prices, "\nDuplicate Values: ", anotherPrices)
	fmt.Println("========================================================================")

	//Now lets try Slice
	ages := [4]string{"Warrior", "Mage", "Rogue", "Cleric"}
	anotherAge := ages[0:4] //Here's how Slice duplicate an array, can write ages[0:len(ages)] but this better
	fmt.Println("Array Values: ", ages)
	fmt.Println("Slice Values: ", anotherAge)

	anotherAge[2] = "Archer" //Because Slice is based on reference, if its changed it will also change the one it duplicates
	fmt.Println("Array Values: ", ages)
	fmt.Println("Slice Values: ", anotherAge)
}