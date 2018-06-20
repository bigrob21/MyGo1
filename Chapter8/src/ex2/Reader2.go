package main

import ("fmt"
		"io/ioutil"
		"log")

/**
*	Simpler method of reading a file using the ioutil package. 
*/
func main() {
	bytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Printf(err.Error())
		return
	}
	fmt.Println(string(bytes))
}