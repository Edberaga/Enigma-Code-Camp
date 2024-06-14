package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	channel := make(chan int, 3)

	wg.Add(3)
	go send(channel, 7)
	go send(channel, 11)
	go send(channel, 3)

	wg.Wait()
	fmt.Println(sum(channel))
	fmt.Println("Done")
}

func send(channel chan int, num int) {
	channel <- num
	wg.Done()
}

func print (channel chan int) {
	fmt.Println(<- channel)
	wg.Done()
}

func sum (channel chan int) int {
	result := 0
	for len(channel) > 0 {
		result += <- channel
	}
	return result
}