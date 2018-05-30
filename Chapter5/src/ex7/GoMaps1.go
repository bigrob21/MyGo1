package main

import (
	"fmt"
)

func main() {
	x := make(map[int]string)
	x[0] = "Rob"
	x[1] = "A."
	x[2] = "Foobar"
	x[3] = "A Map"

	fmt.Println(x)
	fmt.Println("Deleting key 3 from Map")

	delete(x, 3)
	fmt.Println("Viewing the same map again after the delete of the 3 key")
	fmt.Println(x)
}
