package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y2)
	return l * w
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// structure data zone

type rectangle struct {
	x1, y1, x2, y2 float64
}

func main() {
	// rec := new(rectangle) or
	rec := rectangle{x1: 10, y1: 20, x2: 8, y2: 7}
	fmt.Printf("rec %f", rec)
}
