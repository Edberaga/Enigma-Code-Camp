package main 

import (
	"fmt"
)

func main() {
	animal := make(chan string)
	fruits := make(chan string)

	go func() {
		animal <- "Cat"
		animal <- "Dog"
		animal <- "Bird"
		animal <- "Chicken"
		animal <- "Fish"
		close(animal)
	}()

	go func() {
		fruits <- "Apple"
		fruits <- "Banana"
		fruits <- "Grapes"
		fruits <- "Orange"
		fruits <- "Watermelon"
	}()

	var animalStat bool
	var fruitsStat bool

	for {
		select {
		case data, status := <- animal:
			animalStat = status
			if animalStat {
				fmt.Println("Animal: ", data)
			} 
		case data, status := <- fruits: 
			fruitsStat = status
			if fruitsStat {
				fmt.Println("Fruits: ", data)
			} 
		}
		if (!animalStat && !fruitsStat) {
			break
		}
	}
}