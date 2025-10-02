package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

func (p Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Hypot(dx, dy)
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.x, p.y)
}

func main() {
	p1 := NewPoint(1.0, 1.0)
	p2 := NewPoint(5.0, 4.0)

	distance := p1.Distance(p2)

	fmt.Printf("Point A: %s\n", p1)
	fmt.Printf("Point B: %s\n", p2)
	fmt.Printf("Distance between A and B: %.2f\n", distance)
}