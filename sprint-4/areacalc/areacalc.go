package areacalc

import (
	"strings"
)

const pi = 3.14159

type Shape interface {
	Area() float64
	Type() string
}

type Rectangle struct {
	name string
	side []float64
}

func NewRectangle(a, b float64, s string) *Rectangle {
	return &Rectangle{name: s, side: []float64{a, b}}
}

func (r *Rectangle) Area() float64 {
	return r.side[0] * r.side[1]
}

func (r *Rectangle) Type() string {
	return r.name
}

type Circle struct {
	name   string
	radius float64
}

func NewCircle(r float64, s string) *Circle {
	return &Circle{radius: r, name: s}
}

func (c *Circle) Area() float64 {
	return pi * c.radius * c.radius
}

func (c *Circle) Type() string {
	return c.name
}

func AreaCalculator(figures []Shape) (string, float64) {
	var sum float64
	var desc []string
	for _, f := range figures {
		sum += f.Area()
		desc = append(desc, f.Type())
	}
	return strings.Join(desc, "-"), sum
}
