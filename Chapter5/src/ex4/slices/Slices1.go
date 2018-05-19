package main

import ("fmt")

func main(){
	var x []float64
	//Create an empty array using the make() op
	var y = make([]float64, 5,10)
	fmt.Println(x)
	fmt.Println(y)
	
	var z = []float64{1,2,3,4,5,6,7,8}
	var zz = z[0:5]
	fmt.Println(z)
	fmt.Println(zz)
}