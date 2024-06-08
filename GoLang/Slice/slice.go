package main

import "fmt"

func main() {
	numSlice := []int{1, 2, 3, 4}
	fmt.Println(len(numSlice))
	fmt.Println(cap(numSlice))

	fmt.Println("========================================================================")

	scores := make([]int, 4, 10)
	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	for i := range scores {
		scores[i] = 4 * i //Only 4 times, reach the length only not the capacity
	}
	fmt.Println(scores)
	fmt.Println(len(scores))
	fmt.Println(cap(scores))

	//scores[4] = 108 Error: index out of range [4] with length 4

	fmt.Println("========================================================================")

	heroes := [4]string{"Ed", "Sydna", "Hiro", "Ellen"}
	fmt.Println("Heroes: ", heroes)
	//heroes[5] = "Hiya" error: index no bound

	villain := make([]string, 3, 5)
	villain[0] = "Thanos"
	villain[1] = "Joker"
	villain[2] = "Sasuke"
	fmt.Println("Villain: ", villain, "\nPanjang: ", len(villain), "\nKapasitas: ", cap(villain))

	villain = append(villain, "Ultron", "Voldemort", "Homelnder")
	villain[0] = "Thanos"
	villain[1] = "Joker"
	villain[2] = "Sasuke"
	fmt.Println("Villain: ", villain, "\nPanjang: ", len(villain), "\nKapasitas: ", cap(villain)) //Slice cap x2 when length is more then cap
}