package main

import "fmt"

func main() {
	var x int = 4
	var y *int = &x

	fmt.Println("x:", x)
	fmt.Println("Alamat x:", &x)
	fmt.Println("y:", y)
	fmt.Println("Alamat y:", &y)
	fmt.Println("Reference y:", *y)

	*y += 1
	fmt.Println("y:", y)
	fmt.Println("*y:", *y)
	fmt.Println("x:", x)

	//Tess pass by reference through function
	var a int = 4
	fmt.Println("Declaring a:", a)
	increase(&a)
	fmt.Println("After Increasing a: ", a)
}

func increase(n *int) {
	*n++
	fmt.Println("Increasing a: ", *n)
}