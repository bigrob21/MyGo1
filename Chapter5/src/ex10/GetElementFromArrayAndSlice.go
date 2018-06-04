package main

import "fmt"

func main(){
	array := [5]int{1,2,3,4000,5}
	slice := []int{10,20,30,40,50,60,100}

	//Get the 4 element of Array
	fmt.Printf("The 4th element of array is %d\n",array[3])
	fmt.Printf("The 4th element of the slice is %d\n", slice[3])
}