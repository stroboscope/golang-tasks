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
	sideA float64
	sideB float64
}

func NewRectangle(a, b float64, s string) *Rectangle {
	return &Rectangle{name: s, sideA: a, sideB: b}
}

func (r *Rectangle) Area() float64 {
	return r.sideA * r.sideB
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
