package main

import ("fmt")

func main(){
	x := []int{1,10,25,30}	
	fmt.Println(x)
	xx := append(x,100)
	fmt.Println(xx)
}