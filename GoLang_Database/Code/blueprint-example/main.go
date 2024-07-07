package main

import (
	"fmt"
	"github.com/Edberaga/Geomerty-Bin/"
)

func main() {
	firstFloor := shape.Rectangle{Width: 7.5, Length: 10.0}
	secondFloor := shape.Rectangle{Width: 4.5, Length: 7.5}

	totalArea := firstFloor.Area() + secondFloor.Area()

	fmt.Println("Total Area: ", totalArea)
}