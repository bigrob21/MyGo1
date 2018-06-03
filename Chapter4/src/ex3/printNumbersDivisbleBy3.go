package main

import "fmt"

func main(){
	x := []int{}
	//Build up a list of numbers from 1 to 100
	for z := 1; z<=100; z++ {
		x = append(x,z)
	}
	//Print out the numbers that are divisble by 3
	for i := 0; i<len(x); i++ {
		if x[i] % 3 == 0 {
			fmt.Printf("%d is divisble by 3 \n", x[i])
		}
	}
}