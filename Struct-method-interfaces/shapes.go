package shapes

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Perimeter(rectangle Rectangle) float64 {

	return 2 * (rectangle.width + rectangle.height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.width * rectangle.height
}