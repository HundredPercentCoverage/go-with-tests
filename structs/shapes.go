package structs

import "math"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Shape interface {
	Area() float64
	// Perimeter() float64
}

// Calculates the perimeter of a rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Calculates the area of a rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// func (t Triangle) Perimeter() float64 {
// 	return 2 * math.Pi * c.Radius
// }

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2.0
}
