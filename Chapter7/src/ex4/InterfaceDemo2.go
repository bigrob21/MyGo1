package main

import "fmt"

/**
	Notice the same signature of the two interface.
**/

type I1 interface {
	M1()
}

type I2 interface {
	M2()
}

type T1 struct { }

func (T1) M1() {fmt.Println("Printing the M1 version")}
func (T1) M2() {fmt.Println("Printing the M2 version")} 

func f1(i I1){i.M1()}
func f2(i I2){i.M2()}

func main() {
	t := T1{}
	f1(t)
	f2(t)
}