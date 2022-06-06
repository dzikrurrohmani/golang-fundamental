package data

import "math"

type Circle struct {
	Radius float32
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(float64(c.Radius), 2)
}

func (c *Circle) Perimeter() float32 {
	return 2 * math.Pi * c.Radius
}
