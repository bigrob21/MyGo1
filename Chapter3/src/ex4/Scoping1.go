package main

import ("fmt")

//This is final
const a string = "A Global Immutable Variable"
var z string = "A mutable global"

func main(){
	b := "A mutable variable in the main function"
	fmt.Println("printing..." + b)
	aFunc()
	aFunc2()
}

func aFunc(){
	c := "A call from aFunc()"
	fmt.Println("=======================================\n")
	fmt.Println("printing my variable --- " + c)
	fmt.Println("=======================================\n")
}

func aFunc2(){
	fmt.Println("Printing the global scoped variable 'Z' ---->> " + z)
	z = "Overriding Z, no seriously!!"
	fmt.Println("Reprinting Z, but checkout the overriding... " + z)
}

