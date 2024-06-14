package main

import(
	"fmt"
)

func main() {
	channel := make(chan string)
	defer close(channel)
	go func() {
		channel <- "Hello World!"
	}()

	// words := <- channel //if ignore the <- will output the memory address
	// fmt.Println(words)

	fmt.Println(<- channel)

	fmt.Println("Done")
}