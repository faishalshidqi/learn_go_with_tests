package structsMethodsInterfaces

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasPerimeter: 40},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6, sideC: 6}, hasPerimeter: 24},
	}
	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			checkPerimeter(t, tt.shape, tt.hasPerimeter)
		})
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 10.0, Height: 10.0}, hasArea: 100},
		{name: "Circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6, sideC: 6}, hasArea: 36},
	}
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			checkArea(t, tt.shape, tt.hasArea)
		})
	}
}

func checkArea(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Area()
	if got != want {
		t.Errorf("%#v got %g, want Area %g", shape, got, want)
	}
}

func checkPerimeter(t *testing.T, shape Shape, want float64) {
	t.Helper()
	got := shape.Perimeter()
	if got != want {
		t.Errorf("%#v got %g, want Perimeter %g", shape, got, want)
	}
}
