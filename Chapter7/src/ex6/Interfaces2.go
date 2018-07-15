package main

import "fmt"

type Logic struct{ /* */ } 

//var _ broker.Interface = &BusinessLogic{}
var sumOfAllInterfaces Interface3 = &Logic{}

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

func main() {
	fmt.Println(sumOfAllInterfaces.SayHiInterface1())
	fmt.Println(sumOfAllInterfaces.SayHiInterface2())
	fmt.Println(sumOfAllInterfaces.SayHiInterface3())
}