package main 

import "fmt"

func main() {
	var button int
	fmt.Print("Enter Button Number: ")
	fmt.Scanln(&button)

	switch button {
		case 1, 2, 3: {
			fmt.Println("Beginner Level")
		}
		case 4, 5, 6: {
			fmt.Println("Normal Level")
		}
		case 7, 8, 9: {
			fmt.Println("Expert Level")
		}
		default: {
			fmt.Println("Unknown")
		}
	}
}