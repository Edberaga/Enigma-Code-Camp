package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	account := BankAccount{Balance : 2500}
	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go account.invest(10)
	}
	wg.Wait()
	fmt.Printf("Bank account current balance: %.2f\n", account.Balance)
}

type BankAccount struct {
	Balance float64
	mutex sync.Mutex
}

func (acc *BankAccount) invest(ammount float64) {
	acc.mutex.Lock()
	acc.Balance += ammount * 0.0255
	wg.Done()
	acc.mutex.Unlock()
}