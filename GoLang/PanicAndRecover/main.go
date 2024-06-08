package main

import "fmt"

func main() {
	defer func() {
		if r:= recover(); r != nil {
			fmt.Println("Terjadi Panic: ", r)
		}
	}()

	name := ""
	fmt.Print("Write your name: ")
	fmt.Scanln(&name)
	isEmpty(name)
	fmt.Println("The name is: ", name)
	fmt.Println("Is it still continue?")
}

func isEmpty(name string) bool {
	if(name == "") {
		panic("Name can't be empty!")
	}
	return false
}