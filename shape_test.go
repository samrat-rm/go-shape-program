package main

import (
	"testing"
)

func TestSquareArea(t *testing.T) {
	s := Square{side: 5}
	expectedArea := 25.0
	if s.Area() != expectedArea {
		t.Errorf("Square area calculation incorrect. Got %f, expected %f.", s.Area(), expectedArea)
	}
}

func TestSquarePerimeter(t *testing.T) {
	s := Square{side: 5}
	expectedPerimeter := 20.0
	if s.Perimeter() != expectedPerimeter {
		t.Errorf("Square perimeter calculation incorrect. Got %f, expected %f.", s.Perimeter(), expectedPerimeter)
	}
}

func TestRectangleArea(t *testing.T) {
	r := Rectangle{length: 4, width: 6}
	expectedArea := 24.0
	if r.Area() != expectedArea {
		t.Errorf("Rectangle area calculation incorrect. Got %f, expected %f.", r.Area(), expectedArea)
	}
}

func TestRectanglePerimeter(t *testing.T) {
	r := Rectangle{length: 4, width: 6}
	expectedPerimeter := 20.0
	if r.Perimeter() != expectedPerimeter {
		t.Errorf("Rectangle perimeter calculation incorrect. Got %f, expected %f.", r.Perimeter(), expectedPerimeter)
	}
}

func TestNewSquare(t *testing.T) {
	_, err := NewSquare(-1)
	if err == nil {
		t.Errorf("Expected error for negative side length")
	}

	_, err = NewSquare(0)
	if err == nil {
		t.Errorf("Expected error for side length of 0")
	}
}

func TestNewRectangle(t *testing.T) {
	_, err := NewRectangle(-1, 5)
	if err == nil {
		t.Errorf("Expected error for negative length")
	}

	_, err = NewRectangle(4, -2)
	if err == nil {
		t.Errorf("Expected error for negative width")
	}

	_, err = NewRectangle(0, 5)
	if err == nil {
		t.Errorf("Expected error for length of 0")
	}

	_, err = NewRectangle(4, 0)
	if err == nil {
		t.Errorf("Expected error for width of 0")
	}
}