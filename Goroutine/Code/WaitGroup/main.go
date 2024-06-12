package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Start")
	wg.Add(1)
	go repeat("Cat", 200)
	go repeat("Dog", 1000)
	wg.Wait()
	fmt.Println("Done")
}

func repeat(word string, delay time.Duration) {
	for i := 0; i <= 10; i++ {
		fmt.Println(i, word)
		time.Sleep(time.Millisecond * delay)
	}
	wg.Done()
}