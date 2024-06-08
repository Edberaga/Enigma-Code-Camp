package main

import "fmt"


func lessonOne() {
	defer fmt.Println("End") //Defer will be executed at the last order
	fmt.Println("Start")
	fmt.Println("Processing")
}

func lessonSecond() {
	num := 4
	defer fmt.Println("Defer number: ", num) //The number stays 4 as its before the line where the num add 10 times.
	num *= 10
	defer fmt.Println("Defer number after line: ", num) // It will result 40 as this line executed after the num has been added
	fmt.Println("Result number: ", num)
}

func lessonThird() {
	defer add(6, multiply(3, 4)) //multiply(3, 4) called first. this beacause the differ is the add()
	add(11, 10)
}

func lessonFourth() {
	//Defer run is LIFO (Last in First out), which is a descending order where the first defer is the last one executed
	defer fmt.Println("End")
	fmt.Println("Start")
	fmt.Println("Processing")
	defer add(5, 2)
	defer add(11, 12)
	defer multiply(3, 4)
}

func lessonFifth() {
	fmt.Println("Start")
	defer loop()
	fmt.Println("Processing")
}

func add( a, b int) {
	fmt.Println(a + b)
}

func multiply(a, b int) int{
	fmt.Println(a * b)
	return a * b
}

func loop() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done Loop")
}

func lessonSix() {
	defer fmt.Println("Pertama")
	defer fmt.Println("Kedua")
	defer fmt.Println("Terakhir")
}

func main() {
	lessonOne() //Note: Introduction about Defer
	lessonSecond() //Note: How defer only executed according to order, but not with any calculation
	lessonThird() //Note: differ function as parameter will be ignored, and only the main function will be defered.
	lessonFourth() //Note: Show that defer runs in LIFO
	lessonFifth() //Note: show how differ Loop will run LIFO because thats how differ runs.
	lessonSix()
}