package main

import ("fmt")

func main(){

	var u = [3]float64{1,100,1000}
	var total float64 = 0.0
	for _, value:= range u {
		total += value
	}
	fmt.Println(float64(total))
	
}