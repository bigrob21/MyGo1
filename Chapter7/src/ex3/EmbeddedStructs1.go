package main

import "fmt"

type Person struct {
	Name string
}

type Android struct{
	//Note the missing type, this allows the android to leverage te receiver method communicate
	Person
	Model string
}

func (p *Person) communicate(){
	fmt.Printf("Hi! I'am %s\n", p.Name)
}

func main() {
	p1 := Person{Name:"Robert"}
	p2 := Person{Name: "Arnold, the android"}

	a1 := Android{Person: p2, Model: "T-800"}
	a2 := new(Android)

	p1.communicate()
	a1.Person.communicate()

	a2.communicate()
}