package main

import (
	"fmt"
	"math"
)

// shape interface
type Shape interface {
	area() float64
	circumf() float64
}

type Circle struct {
	radius float64
}

type Square struct {
	length float64
}

// Square methods
func (s Square) area() float64 {
	return s.length * s.length
}

func (s Square) circumf() float64 {
	return s.length * 4
}

// Circle methods
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

func printShapeInfo(s Shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func generateShapes() {
	shapes := []Shape{
		Circle{radius: 7.5},
		Square{length: 15.2},
		Circle{radius: 12.3},
		Square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("-----")
	}
}
