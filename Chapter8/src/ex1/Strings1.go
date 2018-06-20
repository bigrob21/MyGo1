package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "Welcome to the Jungle Baby"
	fmt.Printf("String cotains 'Jungle'? => %t\n", strings.Contains(str, "Jungle"))
	fmt.Printf("String contains 'Developer'? => %t\n",strings.Contains(str, "Developer"))
}
