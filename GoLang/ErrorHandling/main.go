package main

import (
	"errors"
	"fmt"
)

func main() {
	result, err := divide(5, 5)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println("Hasil pembagian: ", result)
}

func divide(a int, b int) (int, error) {
	if(b == 0) {
		return 0, errors.New("divider can't be 0")
	}
	return a / b, nil
}