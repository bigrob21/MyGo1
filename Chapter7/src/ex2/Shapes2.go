package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
}

//Normal function with Pointer
// func calculateCircleArea2(cPtr *Circle) float64 {
// 	retVal := math.Pi * cPtr.r * cPtr.r
// 	PrintStuff(retVal)
// 	return retVal
// }

//This is a method in Go, note the receiver between the input param 'receiver' and an empty arg list, which moved to the receiver portion
func (cPtr *Circle) calculateCircleArea2() float64 {
	retVal := math.Pi * cPtr.r * cPtr.r
	PrintStuff(retVal)
	return retVal
}

func calculateCircleArea(c Circle) float64 {
	retVal := math.Pi * c.r * c.r
	PrintStuff(retVal)
	return retVal
}

func PrintStuff(v float64) {
	fmt.Printf("Calculated area is: %f\n", v)
}

func main() {
	c := Circle{0, 0, 7}
	c2 := Circle{x: 2, y: 5.0, r: 8}
	//Takes the reference
	calculateCircleArea(c)
	//Takes the Pointer Referece
	c2.calculateCircleArea2()
	//calculateCircleArea2(&c2)
}
