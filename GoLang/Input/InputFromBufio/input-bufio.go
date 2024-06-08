package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv" // Convert string into integer
)

func main() {
	var fullName string
	var address string
	var age int

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enigma Camp Detail")

	//Camper Name
	fmt.Print("Enter your name: ")
	scanner.Scan()
	fullName = scanner.Text()

	//Camper Age
	fmt.Print("Enter your age: ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text()) //Atoi = ASCII into INTEGER

	//Camper Address
	fmt.Print("Enter your Address: ")
	scanner.Scan()
	address = scanner.Text()

	fmt.Printf("Name: %s \n Age: %d \n Addres: %s\n", fullName, age, address)
}
