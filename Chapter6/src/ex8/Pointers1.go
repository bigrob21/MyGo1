package main

import "fmt"

func main() {
	x := 54
	fmt.Printf("The initial value of x is %d\n", x)
	zero1(x)
	fmt.Printf("Testing Pass By Value, the value after func call is still %d\n", x)
	//Note I'm using a Pointer reference to x
	zero2(&x)
	fmt.Printf("Testing Pass By Reference using Pointer, the value after func call is %d\n", x)
}

func zero1(v int) {
	v = 0
}

func zero2(vPtr *int) {
	*vPtr = 0
}
