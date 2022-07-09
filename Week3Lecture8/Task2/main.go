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

type Shapes []Shape

func (s Shapes) LargestArea() float64 {

	pom := s[0]
	for _, s := range s {
		if s.Area() > pom.Area() {
			pom = s
		}
	}
	return pom.Area()

}

func main() {
	c := Circle{radius: 2}
	fmt.Println(c.Area())
	s := Square{2}
	fmt.Println(s.Area())

	sh := Shapes{
		Square{2},
		Circle{3},
		Square{3},
		Square{15},
		Circle{6},
		Square{7},
		Square{5},
		Square{3},
	}
	fmt.Println(sh.LargestArea())

}
