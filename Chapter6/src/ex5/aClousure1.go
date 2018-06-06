package main

import "fmt"

func main(){
	const a = 2
	//Declaring a function within a function (aka Closure)
	multiply := func(a,b int) (int) {
		return a * b
	}
	retval1 := multiply(a,100)
	print(a,100, retval1)
}

/**
Convenience printer
*/
func print(in1, in2, result int) {
	fmt.Printf("%d * %d = %d", in1,in2,result)
}