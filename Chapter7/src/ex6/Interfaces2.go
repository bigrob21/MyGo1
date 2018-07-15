package main

import "fmt"

//Struct definition of a 'Logic' class
type Logic struct{ /* */ } 
//Struct definition of a 'SmallLogic' class
type SmallLogic struct { /* */ }

//statement makes a reference pointer to Logic and assigns it to the super aggregate interface - Interface3. 
//This insures that the programmer must implement all interfaces to satisfy the Interface
var sumOfAllInterfaces Interface3 = &Logic{}
var smallSubsetOfInterface1 Interface1 = &SmallLogic{}

type Interface1 interface{
	SayHiInterface1() (string)
}

type Interface2 interface{
	SayHiInterface2() (string)
}

type Interface3 interface {
	//Note the "embeddedness of two separate interface. It just means that this contract will combine the other 2"
	Interface1
	Interface2
	SayHiInterface3() (string)
}

func (*Logic) SayHiInterface1() (string) {
	return fmt.Sprintf("Hello from INTERFACE #1....")
}

func (*Logic) SayHiInterface2() (string) {
	return fmt.Sprintf("Hello from INTERFACE #2....")
}

func (*Logic) SayHiInterface3() (string) {
	return fmt.Sprintf("Hello from INTERFACE #3....")
}

func (*SmallLogic) SayHiInterface1() (string) {
	return fmt.Sprintf("I satisfy Interface#1 only....")
}

func main() {
	fmt.Println("---------------------------------------")
	fmt.Println("\t Invoking the Super Implementation of Logic (a.k.a BigLogic)")
	fmt.Println("---------------------------------------")
	fmt.Println(sumOfAllInterfaces.SayHiInterface1())
	fmt.Println(sumOfAllInterfaces.SayHiInterface2())
	fmt.Println(sumOfAllInterfaces.SayHiInterface3())
	fmt.Println("---------------------------------------")
	fmt.Println("\t Invoking the implementation of SmallLogic")
	fmt.Println("---------------------------------------")
	fmt.Println(smallSubsetOfInterface1.SayHiInterface1())
}