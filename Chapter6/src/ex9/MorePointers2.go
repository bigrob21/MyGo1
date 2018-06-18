package main

import "fmt"

func main() {
	x := 1.5
	fmt.Printf("X's original value is %f\n", x)
	square(&x)
	fmt.Printf("New value of X squared %f\n",x)
}

func square(x *float64) {
	*x = *x * *x
}