package main

import "fmt"

func main(){
	x := []int{25,22,30,50,0}
	//Have to use elipse for the array/slice
	fmt.Printf("Summation is %d", sum(x...))
}

func sum(a ...int) (int){
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}