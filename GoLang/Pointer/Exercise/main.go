package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	//Input:
	scanner.Scan()
	giver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	receiver, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	marbles, _ := strconv.Atoi(scanner.Text())

	calculate(&giver, &receiver, marbles)
	fmt.Println(giver)
	fmt.Println(receiver)
}

func calculate(x *int, y *int, z int) (int, int) {
	if *x < z {
		fmt.Println("Kelereng tidak cukup untuk dibagikan")
	} else {
		*x -= z
		*y += z
	}
	
	return *x, *y
}