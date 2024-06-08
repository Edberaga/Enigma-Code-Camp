package main

import "fmt"

func main() {
	fmt.Println("Data type")
	var positiveNum uint = 90
	var negativeNum int = -90
	fmt.Printf("Positive num: %d\n", positiveNum)
	fmt.Printf("Negative num: %d\n", negativeNum)

	var decimalNum = 3.90
	fmt.Printf("Decimal num: %f\n", decimalNum)
	fmt.Printf("Decimal num: %.2f\n", decimalNum)

	var trueBoolean = true
	var falseBoolean = false
	fmt.Printf("True Boolean type \"%t\"\n", trueBoolean)
	fmt.Printf("False Boolean type \"%t\"\n", falseBoolean)

	var backTicMes = `This is a message from 	backtick message,
	the data is assembles according to the code,
	you create new line		and you tab,   you spaces`
	fmt.Println(backTicMes)

	var message = "Halo"
	fmt.Printf("Print String variables %s\n", message)

	fmt.Println(len(message))
}