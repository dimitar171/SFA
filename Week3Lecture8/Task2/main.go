package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

func (sq Square) Area() float64 {
	return sq.side * sq.side
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	Area() float64
}

type Shapes struct {
	shapes []Shape
}

func (s Shapes) LargestArea() float64 {

	pom := s.shapes[0].Area()
	for _, s := range s.shapes {
		if s.Area() > pom {
			pom = s.Area()
		}
	}
	return pom

}

func main() {
	c := Circle{radius: 2}
	fmt.Println(c.Area())
	s := Square{2}
	fmt.Println(s.Area())
	sh := Shapes{
		shapes: []Shape{
			Square{2},
			Circle{3},
			Square{3},
			Square{15},
			Circle{6},
			Square{7},
			Square{5},
			Square{3},
		}}
	fmt.Println(sh.LargestArea())

}
