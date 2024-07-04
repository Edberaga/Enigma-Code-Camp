package main

import (
	"fmt"
	"geometry-bin/shape"
)

func main() {
	field1 := shape.Rectangle{Width: 15.0, Length: 11.0}
	field2 := shape.Rectangle{Width: 12.5, Length: 11.5}

	harvestField := field1.Area() + field2.Length

	fmt.Println("Harvest Field Area: ", harvestField)
}