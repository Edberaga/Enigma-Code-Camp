package main

import "fmt"

func main() {
	var sum func(int, int) int = add //it needs to be "var"
	fmt.Println("Hasil Penjumlahan:", sum(3, 4))
}

func add(a, b int) int {
	return a + b
}