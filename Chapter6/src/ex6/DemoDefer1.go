package main 

import "fmt"

func main() {
	fmt.Println("Running main() normally -- BEGIN")
	first()
	second()
	fmt.Println("Running main() normally -- END")
	fmt.Println("=====================================")
	fmt.Println("Demoing deferment, Not first is deferred even though it is called before second!!")
	defer first()
	second()
}

func first() {
	fmt.Println("FIRST")
}

func second() {
	fmt.Println("SECOND")
}