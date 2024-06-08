package main

import "fmt"

//lastName := "Lim" Global variables need to be declared with "var"
var lastName = "Lim" 

func main() {
	fmt.Println("Default Last name: ", lastName)
	greetings("Edbert", 24)
}

func greetings(name string, age int) {
	lastName := "Kirisame" //If there is same global variable name, the local will take it place
	fmt.Println("Hello My name is", name, lastName, "I am", age, "years old")
}