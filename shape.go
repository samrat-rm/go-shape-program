package main

import (
	"fmt"
	"log"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side float64
}

func NewSquare(side float64) (*Square, error) {
	if side <= 0 {
		return nil, fmt.Errorf("side length should be greater than 0")
	}
	return &Square{side}, nil
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func (s Square) Perimeter() float64 {
	return 4 * s.side
}

type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length, width float64) (*Rectangle, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length should be greater than 0")
	}
	if width <= 0 {
		return nil, fmt.Errorf("width should be greater than 0")
	}
	return &Rectangle{length, width}, nil
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func main() {
	square, err := NewSquare(5)
	if err != nil {
		log.Fatal(err)
	}
	rectangle, err := NewRectangle(4, 6)
	if err != nil {
		log.Fatal(err)
	}

	shapes := []Shape{square, rectangle}

	for _, shape := range shapes {
		fmt.Printf("Shape: %T\n", shape)
		fmt.Printf("Area: %f\n", shape.Area())
		fmt.Printf("Perimeter: %f\n", shape.Perimeter())
		fmt.Println()
	}
}
