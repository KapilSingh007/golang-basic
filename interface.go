package main

import "fmt"

type Shape interface {
	getPerimeter() uint
}
type Square struct {
	side uint
}
type Triangle struct {
	a uint
	b uint
	c uint
}

func (s Square) getPerimeter() uint {
	return s.side * 4
}

func (s Triangle) getPerimeter() uint {
	return s.a + s.b + s.c
}
func main() {

	var shapes []Shape = []Shape{Triangle{a: 10, b: 5, c: 12}, Square{4}}

	perimeter := uint(0)

	for _, shape := range shapes {
		perimeter += shape.getPerimeter()
	}

	fmt.Println(perimeter)

}
