package main

import "fmt"

func main(){
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Flipping out in Panic Mode Now!")
}

/*func run1(){
	panic("Executing Panic Mode")	
	str := recover()
	fmt.Println(str)
}*/