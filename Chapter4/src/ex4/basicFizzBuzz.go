package main

import "fmt"

func main(){
	arr := generateNumbers1To100()
	for _, z := range arr {
		if isFizz(z) {
			print("FIZZ", z)
		}
		if isBuzz(z) {
			print("BUZZ", z)
		}
		if isFizzBuzz(z){
			print("FIZZBUZZ", z)
		}
	}
}

func print(x string, y int) {
	fmt.Printf("Number %d is %s\n", y, x)
}

func generateNumbers1To100() ([]int) {
	x := []int{}
	for i := 1; i <= 100; i ++ {
		x =  append(x,i)
	}
	return x
}

func isFizz(i int) bool {
	return  i % 3 == 0
}

func isBuzz(i int) bool {
	return  i % 5 == 0
}

func isFizzBuzz(i int) bool {
	return isFizz(i) && isBuzz(i)
}