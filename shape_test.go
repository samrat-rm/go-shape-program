package main

import (
	"fmt"
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
func TestTriangle(t *testing.T) {
	tests := []struct {
		name     string
		a        float64
		b        float64
		c        float64
		area     float64
		perimeter float64
		err      error
	}{
		{
			name:      "valid triangle",
			a:         3,
			b:         4,
			c:         5,
			area:      6,
			perimeter: 12,
			err:       nil,
		},
		{
			name:      "zero sides",
			a:         0,
			b:         0,
			c:         0,
			area:      0,
			perimeter: 0,
			err:       fmt.Errorf("sides should be greater than 0"),
		},
		{
			name:      "negative sides",
			a:         -2,
			b:         3,
			c:         4,
			area:      0,
			perimeter: 0,
			err:       fmt.Errorf("sides should be greater than 0"),
		},
		{
			name:      "invalid triangle",
			a:         1,
			b:         2,
			c:         4,
			area:      0,
			perimeter: 0,
			err:       fmt.Errorf("sides do not form a valid triangle"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			triangle, err := NewTriangle(tt.a, tt.b, tt.c)

			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf("Expected error: %v, but got: %v", tt.err, err)
			} else if err == nil && tt.err != nil {
				t.Errorf("Expected error: %v, but got: %v", tt.err, err)
			} else if err == nil && tt.err == nil {
				if triangle.Area() != tt.area {
					t.Errorf("Expected area: %f, but got: %f", tt.area, triangle.Area())
				}
				if triangle.Perimeter() != tt.perimeter {
					t.Errorf("Expected perimeter: %f, but got: %f", tt.perimeter, triangle.Perimeter())
				}
			}
		})
	}
}
