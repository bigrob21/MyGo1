package main

import "fmt"

func main(){
	elements := map[string]map[string]string{
		"H": map[string]string {
			"name":"Hydrogen",
			"state":"gas",
		},
		"O": map[string]string {
			"name":"Oxygen",
			"state":"gas",
		},
		"PB": map[string]string {
			"name":"Lead",
			"state":"solid",
		},
		"BR": map[string]string {
			"name":"Bromine",
			"state":"liquid",
		},
	}
	fmt.Println(elements)
	fmt.Println("================\n")
	if elem, ok := elements["O"]; ok {
		fmt.Println(elem["name"], elem["state"])
	}
}