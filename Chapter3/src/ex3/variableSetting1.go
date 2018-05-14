package main 

import "fmt"

func main() {
	var x string
	x = "Hello"
	fmt.Println(x)
	x = "Developer"
	fmt.Println(x)
	x = x + "- CONCATTED"
	fmt.Println(x)
	//Note this is shorthand setting
	y := "NewVar 'Y'"
	fmt.Println(y)
}

