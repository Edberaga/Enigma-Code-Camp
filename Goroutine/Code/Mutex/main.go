package main

import (
	"fmt"
	"sync"
)

var x = 0

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mutex)
	}
	wg.Wait()
	fmt.Println("Counter = ", x)
}

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	for i := 1; i <= 100; i++ {
		m.Lock()
		x = x + 1
		m.Unlock()
	}
	wg.Done()
}