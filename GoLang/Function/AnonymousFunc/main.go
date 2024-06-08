package main

import "fmt"

func main() {
	//Anonymous Function
	func() {
		fmt.Println("This is an Anonymous Function")
	}()

	//Anonymous Function through Variables
	greet := func() {
		fmt.Println("Anonymous function through variables")
	}
	greet()

	//Anonymous Function with parameters and passing arguements
	func(name string) {
		fmt.Println("Hello", name)
	}("Edbert Lim")

	//Passing arguement with variables anonymous function
	hello := func(name string) {
		fmt.Println("Hello,", name)
	}
	hello("Edbert")

	//Passing Anonymous Function as arguement
	greetEnglish := func(name string) string {
		return "Hello " + name
	}

	greeting("Edbert", greetEnglish)

	//Another practice, passing anonymous func as arguement
	x := 5
	y := 15
	add := func(a, b int) int {
		return a + b
	}

	minus := func(a, b int) int {
		return a - b
	}

	multiply := func(a, b int) int {
		return a * b
	}

	divide := func(a, b int) int {
		return b / a
	}

	fmt.Println(calculate(x, y, add))
	fmt.Println(calculate(x, y, minus))
	fmt.Println(calculate(x, y, multiply))
	fmt.Println(calculate(x, y, divide))
}

func greeting(name string, function func(name string) string) {
	fmt.Println(function(name))
}

func calculate(a, b int, function func(int, int) int) int {
	return function(a, b)
}
