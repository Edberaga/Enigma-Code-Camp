package main

import "fmt"

func main() {
	a := 10
	var b int

	b = a
	fmt.Println("Nilai B: ", b)
	b +=a //b = b + a
	fmt.Println("Nilai B: ", b)
	b -= a //b = b - a
	fmt.Println("Nilai B: ", b)
	b *= a //b = b * a
	fmt.Println("Nilai B: ", b)
	b /= a //b = b / a
	fmt.Println("Nilai B: ", b)
	b %= a //b = b % a
	fmt.Println("Nilai B: ", b)
}