package main

import "fmt"

type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type rectangle struct {
	width float64
	height float64
}

type square struct {
	shape float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) getPerimeter() float64 {
	return 2 * (r.width * r.height)
}

func (s square) getArea() float64 {
	return s.shape * 2
}

func (s square) getPerimeter() float64 {
	return s.shape * 4
}

func main(){
	fmt.Println("=========UPSCALLING WAY=========")
	var Shape1 Shape = square{shape: 5}
	var Shape2 Shape = rectangle{width: 10, height: 3}
	fmt.Printf("%#v Memiliki Luas: %.1f dan Keliling: %.1f\n", Shape1, Shape1.getArea(), Shape1.getPerimeter())
	fmt.Printf("%#v Memiliki Luas: %.1f dan Keliling: %.1f\n", Shape2, Shape2.getArea(), Shape2.getPerimeter())
	//fmt.Printf("Square Shape: %.1f", Shape1.shape) Error: undefined

	fmt.Println("=========DOWNSCALLING WAY=========")
	var Square1 square = Shape1.(square)
	fmt.Printf("%#v Memiliki Luas: %.1f dan Keliling: %.1f\n", Square1, Square1.getArea(), Square1.getPerimeter())
	fmt.Printf("Square Shape: %.1f\n", Square1.shape)

	//Interface data type can be anything and changes:
	fmt.Println("=== Interface data type can be anything and changes ===")
	var anything interface{}

	anything = "Hello World"
	fmt.Printf("Type: %T, Value: %v\n", anything, anything)

	anything = 1008
	fmt.Printf("Type: %T, Value: %v\n", anything, anything)

	anything = true
	fmt.Printf("Type: %T, Value: %v\n", anything, anything)

	anything = []string{"Edbert", "SW", "Andre", "YZ"}
	fmt.Printf("Type: %T, Value: %v\n", anything, anything)
}