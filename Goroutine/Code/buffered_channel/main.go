package main

import (
	"fmt"
)

func main() {
	channel := make(chan int, 3)

	channel <- 1
	channel <- 2
	channel <- 3
	
	fmt.Println("Channel current length: ", len(channel))
	fmt.Println(<- channel) //it's 1, because buffered channel is FIFO, its output the first one in
	//channel <- 4
	fmt.Println("Channel current length: ", len(channel)) //after every output, the channel length will decrease. which is now 2, while it was 3.
	fmt.Println(<- channel)
	fmt.Println(<- channel)
	//fmt.Println(<- channel) //if try again more than capacity it will error deadlock

	fmt.Println("Channel cap: ", cap(channel))

	fmt.Println("It is done")
}