package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Circle struct {
	x, y, r float64
}

//This is a method in Go, note the receiver between the input param 'receiver' and an empty arg list, which moved to the receiver portion
func (cPtr *Circle) area() float64 {
	retVal := math.Pi * cPtr.r * cPtr.r
	PrintStuff(retVal)
	return retVal
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	lgth := distance(r.x1, r.y1, r.x2, r.y2)
	wdth := distance(r.x1, r.y1, r.x2, r.y2)
	retVal := lgth * wdth
	PrintStuff(retVal)
	return retVal
}

func PrintStuff(v float64) {
	fmt.Printf("Calculated area is: %f\n", v)
}

func main() {
	c := Circle{0, 0, 7}
	c2 := Circle{x: 2, y: 5.0, r: 8}
	r1 := Rectangle{2, 2, 4, 4}

	c.area()
	c2.area()
	r1.area()
}