package main

import "fmt"

type Celcius float64

func (c Celcius) toFahrenhiet() float64 {
	//return c *9 / 5 + 32 Note: error for not able to use Celcius data even if its from float64
	return float64(c * 9 / 5 + 32)
}

func (c *Celcius) add(value float64) {
	//*c += value Note: its error because it only can be add with the same type...
	*c += Celcius(value) 
}

func main() {
	var derajat Celcius = 20.0
	fmt.Println("Derajat Celcius: ", derajat)
	fmt.Println("Derajat Fahrenhit: ", derajat.toFahrenhiet())
	derajat.add(5)
	fmt.Println("Derjat Sekarang: ", derajat.toFahrenhiet())
}
