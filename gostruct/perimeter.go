package structs

import "math"

// Rectangle struct
type Rectangle struct {
	Height float64
	Width  float64
}

// Area Rectangle method
func (r Rectangle) Area() float64 {
	return (r.Height * r.Width) / 2
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area Circle method
func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

// Triangle struct
type Triangle struct {
	Base   float64
	Height float64
}

// Area Triangle method
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

// Shape interface
type Shape interface {
	Area() float64
}

// Perimeter function
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}
