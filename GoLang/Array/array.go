package main

import "fmt"

func main() {
	var fruits = [4]string{"Apel", "Banana", "Orange", "Mie"}
	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fruits[1] = "Jeruk"
	fmt.Println(fruits)
	fmt.Println("========================================================================")

	var scores [3]int
	scores[0] = 252

	fmt.Println(scores)
	scores[1] = 213
	scores[2] = 85
	//scores[3] = 114 Error: Array\array.go:19:9: invalid argument: index 3 out of bounds [0:3]

	fmt.Println(scores)
	fmt.Println("========================================================================")

	var days = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	fmt.Println(days)
	//days[7] = "Libur" Error: index out of bonds
	
	fmt.Println("Jumlah Elemen: ", len(days))
}
