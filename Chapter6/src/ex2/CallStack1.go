package main

import "fmt"

func func1(){
	print("func1()")
}

func func2(){
	print("func2()")
}

func main(){
	print("main()")
	func1()
	func2()
}

func print(str string) {
	fmt.Printf("Calling from %s\n", str)
}