package main

import "fmt"

func main(){
	a,b,c := multiIntReturn()
	fmt.Printf("Received the following 3 values %d,%d,%d", a,b,c )
}

func multiIntReturn() (int, int, int) {
	return 1, 100, 1000
}