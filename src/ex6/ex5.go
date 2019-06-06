package main

import (
	"fmt"
	"math"
)

type Square struct {
	name string
	a    float64
}

type Circle struct {
	name string
	r    float64
}

func (s Square) area() float64 {
	return (s.a * 4)
}

func (s Circle) area() float64 {
	return (math.Round(100*math.Sqrt(s.r)*math.Pi) / 100)
}

type Shape interface {
	area() float64
}

func info(s Shape) {
	var area float64
	var name string
	area = s.area()
	switch s.(type) {
	case Square:
		name = s.(Square).name
	case Circle:
		name = s.(Circle).name
	}
	fmt.Println("Area of", name, "is", area)
}

func main() {

	square1 := Square{
		name: "Square1",
		a:    4,
	}
	circle1 := Circle{
		name: "Circle1",
		r:    3.2,
	}

	info(square1)
	info(circle1)

}
