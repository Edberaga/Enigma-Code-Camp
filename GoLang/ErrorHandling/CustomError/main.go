package main

import (
	"fmt"
)

type Item struct {
	id    int
	name  string
	price int
}

type ItemNotFound struct {
	id int
}

func (i ItemNotFound) Error() string {
	return fmt.Sprintf("Id Item %d is not founded", i.id)
}

var items = []Item{
	{1, "Laptop", 20000},
	{2, "Phone", 8000},
	{3, "Book", 3000},
}

func getItem(i int) (Item, error) {
	for _, val := range items {
		if(i == val.id) {
			return val, nil
		}
	}
	return Item{}, ItemNotFound{id: i}
}

func main() {
	test, err := getItem(4)
	if(err != nil) {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(test)
}