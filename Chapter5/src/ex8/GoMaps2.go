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
		fmt.Printf("Element: [%s]. Long Name Element: [%s]\n", k,v)
	}

}
