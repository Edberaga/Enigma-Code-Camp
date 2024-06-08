package main

import "fmt"

type hero struct {
	name string
	strength int
	dexterity int
	intellect int
}

func main() {
	//1st Way: simple variable declaration
	var stanly hero
	stanly.name = "Stanly"
	stanly.strength = 37
	stanly.dexterity = 75
	stanly.intellect = 0
	fmt.Println(stanly)

	fmt.Println("------------------")

	//2nd Way: declaration with {} with manual assign
	var andre = hero{}
	andre.name = "Andre"
	andre.strength = 33
	andre.dexterity = 13
	andre.intellect = 13
	fmt.Println(andre)

	fmt.Println("------------------")

	//3rd Way: {} declaration {} with assign value inside the brackets (it requires to correct position property)
	var ray = hero{"Ray", 1, 1, 1}
	fmt.Println(ray)

	fmt.Println("------------------")

	//4th Way: declaration {} and assign value inside the brackets, with property name. Position assign can be different
	var edbert = hero{intellect: 999, strength: 999, dexterity: 999, name: "Edbert"}
	fmt.Println(edbert)

	fmt.Println("------------------")
}