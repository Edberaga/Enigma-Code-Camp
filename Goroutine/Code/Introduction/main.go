package main

import (
	"fmt"
	"time"
)

func main() {
	go repeat("Potion")
	//time.Sleep(time.Second * 2)
	fmt.Scanln() // Will finish counting if you press enter
	fmt.Println("Done")
}

func simpletest() {
	fmt.Println("Satu")
	fmt.Println("Dua")
	time.Sleep(time.Second * 3)
	fmt.Println("Tiga")
	fmt.Println("Empat")
}

func repeat(word string) {
	for i:= 0; i <= 10; i++ {
		fmt.Println(i, word)
		time.Sleep(time.Millisecond * 500)
	}
}