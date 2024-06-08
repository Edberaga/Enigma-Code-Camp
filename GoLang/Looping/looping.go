package main 

import "fmt"

func main() {
	/*for i := 1; i <= 5; i++ {
		fmt.Println("No ke ", i)
	}*/

	var n = 5
	for(n <= 10) {
		fmt.Println("Angka ", n)
		if n == 7 {
			fmt.Println("di pertengahan")
		}
		n++
	}
	fmt.Println("====================================")
	for i := 0; i <= 19; i++ {
		if i % 3 == 0 {
			fmt.Println("Faktor 3")
			continue
		}
		if i == 16 {
			fmt.Println("Angka ke - 16!")
			break
		}
		fmt.Println("Angka: ", i)
	}
}