package main

import "fmt"

func main() {
	a := []float64{1.0, 200.4, 21.2, 99.97}
	fmt.Printf("Average generated is %3f", average(a))
}

func average(xs []float64) float64 {
	total := 0.0
	for _, v:= range xs {
		total += v		
	}
	return total / float64(len(xs))
}
