package structs

import "math"

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Length + rectangle.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

type Rectangle struct {
	Width  float64
	Length float64
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
}
