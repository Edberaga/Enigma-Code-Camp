package main

import "fmt"

func main() {
	total := add(252, 213)
	fmt.Println("Hasil Penjumlahan: ", total)
	sum := multiply(16, 85)
	fmt.Println("Hasil Pekalihan: ", sum)
	fmt.Println("Hasil Kalkulasi: ", add(1008, multiply(12, 11)))
}

func add(x int, y int) int{
	sum := x + y
	return sum
}

func multiply(x, y int) int {
	return x * y
}