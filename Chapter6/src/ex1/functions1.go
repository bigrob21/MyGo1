package main

import ("fmt")

func main(){
	xs := []float64{98,99,1,100,21}
	total := 0.0
	for _, v:=range xs {
		total +=v
	}
	fmt.Println(total /float64(len(xs)))
}
