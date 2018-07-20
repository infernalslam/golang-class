package main

import (
	"fmt"
	"math"
)

// refactoer code function distance change struct parameter
func distance(rec rectangle) float64 {
	a := rec.x2 - rec.x1
	b := rec.y2 - rec.y1
	return math.Sqrt(a*a + b*b)
}

// is so fucking error cause struct rectangle
func rectangleArea(rec rectangle) float64 {
	// l := distance(x1, y1, x2, y2) // error: too many arguments in call to distance
	// w := distance(x1, y1, x2, y2) // error: too many arguments in call to distance
	l := distance(rec)
	w := distance(rec)
	return l * w
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

// structure data zone
type rectangle struct {
	x1, y1, x2, y2 float64
}

type circle struct {
	x1, x2, r float64
}

// call methods circle
func (c *circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	// rec := new(rectangle) or
	rec := rectangle{x1: 10, y1: 20, x2: 8, y2: 7}
	// call structure circle
	// send to distance function
	res := distance(rec) // so, change parameter
	res = rectangleArea(rec)

	fmt.Printf("rec %f\n", res)
	fmt.Println("Gopher is sccuess!")

	cir := circle{0, 0, 4}
	fmt.Printf("circle %f\n", cir)
	fmt.Println(cir.area())
}
