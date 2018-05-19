package main

import (
	"fmt"
)

func main(){

	var x [5]int
	fmt.Println(x)
	
	var y [5]float64
	y[0] = 100
	y[1] = 90
	y[2] = 87
	y[3] = 50
	y[4] = 21
	fmt.Println(y)

	var total float64 = 0
	for _, value := range x {
		total += y[value]
	}	
	
	fmt.Println(total / float64(len(y)))
}