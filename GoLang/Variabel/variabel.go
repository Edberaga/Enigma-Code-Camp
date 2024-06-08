package main

import "fmt"

func main() {
	var firstname string = "Edbert"
	var lastname string = "Lim"

	fmt.Println("Hello,", firstname, lastname)
	fmt.Printf("%s %s Kalimat Kedua \n", firstname, lastname)
	fmt.Printf("Kalimat Ketiga\n")

	var (
		age int
		address string
	)

	age = 24
	address = "Padang"

	fmt.Printf("Hello my name is %s %s, I'm %d years old. I'm from %s \n", firstname, lastname, age, address)

	var (
		campName, campNameAddress = "Enigma Camp", "Jakarta"
	)
	fmt.Println(campName, campNameAddress)

	day := "Monday"
	date := 24
	month := "July"

	fmt.Println("Today date is: ", day, date, month + "2024")

	var number = 20
	number = 21

	const phi = 3.14

	fmt.Println(number, phi)
}