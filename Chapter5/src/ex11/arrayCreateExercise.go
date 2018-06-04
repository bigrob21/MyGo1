package main

import "fmt"

func main() {
	slice := make([]int, 3,9)
	fmt.Printf("The size of the slice using 'make([]int, 3,9)' is %d", len(slice))
}