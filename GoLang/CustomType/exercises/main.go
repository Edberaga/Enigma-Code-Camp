package main

import (
	"fmt"
)

type Meter string
type Temperature int

func(t Temperature) Display() {
	fmt.Printf("Suhu derajat: %d", t)
}

func main() {
	//var distances Meter = 50
	//fmt.Println("Distance: ", distances)
	temp := Temperature(25)
	temp.Display()
}
