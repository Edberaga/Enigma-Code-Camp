package main

import "fmt"

func main() {
	x := 4
	y := x //Pass by Value
	y++

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("Memory x: ", &x)
	fmt.Println("Memory y: ", &y)

	a := 3
	fmt.Println(increase(a)) //Pass by Value through arguement
	fmt.Println("a: ", a)
	fmt.Println("Memory a: ", &a)
}

func increase(n int) int{
	fmt.Println("Memory n: ", &n) //As always pass by value will locate to another address
	return n + 1
}