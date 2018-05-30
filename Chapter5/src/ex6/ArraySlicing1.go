package main 

import ("fmt")

/**
	A slice is like an array but is expandable, unlike an array which is fixed!
*/
func main() {
	//There is a 3 param which is an number which appears to cap the slice. 
	y := make([]float64, 2)
	//append takes a slice as arg1, afterwards it is variable 
	z := append(y, 5,0,2.2)
	//fmt.Println("slice y: " + y + ", slice z" + z)
	fmt.Print(z)
}