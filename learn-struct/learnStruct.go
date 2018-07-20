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

func main() {
	fmt.Print("Hello Gopher!")
	res := distance(1, 2, 3, 4)
	res = rectangleArea(1, 2, 3, 4)
	fmt.Printf("result %f", res)
}
