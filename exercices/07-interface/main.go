package main

import (
	"log"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	height float64
	width  float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2

	return math.Pow((s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC)), 0.5)
}
func (t Triangle) Perimeter() float64 {
	s := (t.SideA + t.SideB + t.SideC)
	return s
}

func (t Rectangle) Area() float64 {

	return t.height * t.width
}
func (t Rectangle) Perimeter() float64 {
	return t.height + t.width
}

func (t Circle) Area() float64 {

	return math.Pi * math.Pow(t.radius, 2)
}
func (t Circle) Perimeter() float64 {
	return float64(2) * math.Pi * t.radius
}

func PrintShapeDetails(s interface{}) {
	x, is := s.(Shape)

	if is {
		log.Println("Area of the Shape ", x.Area())
		log.Println("Area of the Shape ", x.Perimeter())
	} else {
		log.Fatal("Its not a shape")
	}
}

func main() {

	PrintShapeDetails(Triangle{4, 5, 3})

}
