package main 

import "fmt"

func main() {
	var array = [2][3]int{{24, 15, 16}, {22, 32, 64}}
	fmt.Println(array)

	for i := range array {
		for j := range array[i] {
			fmt.Println(array[i][j])
		}
	} 
}