package structs

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 20.0}

	got := Perimeter(rectangle)
	want := 60.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	// create a slice of type struct {...} with inside the values of Rectangle and Circle
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 10, Width: 20}, hasArea: 100},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 5, Height: 10}, hasArea: 25},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
