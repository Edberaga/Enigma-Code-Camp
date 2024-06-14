package main

import(
	"fmt"
)

func main() {
	channel := make(chan string)
	//defer close(channel)

	go passData(channel, "Hello, Edbert!")

	fmt.Println(<- channel)
}

func passData (channel_par chan string, word string) {
	channel_par <- word
}