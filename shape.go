package main

import (
	"fmt"
	"log"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	side float64
}
type Rectangle struct {
	length float64
	width  float64
}
type Triangle struct {
	a float64
	b float64
	c float64
}

func NewSquare(side float64) (*Square, error) {
	if side <= 0 {
		return nil, fmt.Errorf("side length should be greater than 0")
	}
	return &Square{side}, nil
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
func NewTriangle(a, b, c float64) (*Triangle, error) {
	if a <= 0 || b <= 0 || c <= 0 {
		return nil, fmt.Errorf("sides should be greater than 0")
	}
	if a+b <= c || a+c <= b || b+c <= a {
		return nil, fmt.Errorf("sides do not form a valid triangle")
	}
	return &Triangle{a, b, c}, nil
}

func (s Square) Area() float64 {
	return s.side * s.side
}
func (s Square) Perimeter() float64 {
	return 4 * s.side
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (t Triangle) Area() float64 {
	s := (t.a + t.b + t.c) / 2
	return math.Sqrt(s * (s - t.a) * (s - t.b) * (s - t.c))
}
func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
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
	triangle, err := NewTriangle(3, 4, 5)
	if err != nil {
		log.Fatal(err)
	}

	shapes := []Shape{square, rectangle, triangle}

	for _, shape := range shapes {
		fmt.Printf("Shape: %T\n", shape)
		fmt.Printf("Area: %f\n", shape.Area())
		fmt.Printf("Perimeter: %f\n", shape.Perimeter())
		fmt.Println()
	}
}
