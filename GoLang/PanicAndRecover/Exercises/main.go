package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("Defered function")
	}()

	fmt.Println("Before panic")
	panic("Panic Occured")
}