package main

import "fmt"

func main() {
	arr := []int{48,59,66,30,100,50,99,2}
	smallest := 0
	for x := 0; x < len(arr); x++ {
		if x == 0 || arr[x] < smallest  {
			smallest = arr[x]
		}
	}
	fmt.Printf("The smallest number detected is: %d", smallest)
}