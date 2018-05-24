package main

import ("fmt")

func main(){
	x := []int{1,10,25,30}	
	fmt.Println(x)
	xx := append(x,100)
	fmt.Println(xx)
	
	//Now this will be interesting
	y := []int{1,2}
	fmt.Println(" And the size of y is " + string(len(y)))
//	yy := append(y,3,4,5)
//	fmt.Println(yy)
	
}