package main
import (
	"fmt"
)

type Student struct {
	//Isi struct ini
	name string
	age int
}

func (s *Student) Introduction() {
   //Tulis kode disini
   s.name = "Sammy"
   s.age = 17
   fmt.Printf("Hello, my name is %s. Im %d years old", s.name, s.age)
}

func main() {
	// Tulis kode disini
	var student1 Student
	student1.Introduction()
}