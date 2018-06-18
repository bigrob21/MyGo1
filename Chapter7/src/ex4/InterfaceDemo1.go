package main

import "fmt"

type Z interface {
	Z() string
}

type X struct {
	name string
}

func (x X) Z() string {
	return x.name
}

func Hello(z Z) {
	fmt.Printf("Hi, my name is %s\n", z.Z())
}

func main() {
	Hello(X{name: "Robert Da Developer"})
}