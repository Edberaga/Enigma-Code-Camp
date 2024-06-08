package main

import "fmt"

type triangle struct {
	base float64
	height float64
}

func(t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

/*Pass By Values
func(t triangle) increaseSize(n float64) {
	t.base += n
	t.height += n
	fmt.Printf("Size has been increased for %.0f\nTriangle size now: \nbase: %.2f\nheight: %.2f\n", n, t.base, t.height)
}
*/

//Pass By References
func(t *triangle) increaseSize(n float64) {
	t.base += n
	t.height += n
	fmt.Printf("Size has been increased for %.0f\nTriangle size now: \nbase: %.2f\nheight: %.2f\n", n, t.base, t.height)
}

func main() {
	triangle1 := triangle{10, 11}
	area := triangle1.area()
	fmt.Println("Area: ", area)

	triangle1.increaseSize(2)
	fmt.Println("triangle now: ", triangle1.base, triangle1.height)
}