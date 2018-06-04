package main

import (
	"fmt"
)

func main() {
	y := make(map[string]string)
	y["He"] = "Helium"
	y["O"] = "Oxygen"
	y["C"] = "Carbon"
	y["H"] = "Hydrogen"

	for k, v := range y {
		fmt.Printf("Element: [%s]. Long Name Element: [%s]\n", k, v)
	}

	fmt.Println("======================")

	if value, present := y["ABA"]; present {
		fmt.Printf("%s was present but sould not be. elements returned %b", value, present)
	} else {
		fmt.Printf("%s was not present!!", "ABA", present)
	}

}
