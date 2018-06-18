package main

import "fmt"

type Circle struct {
	x,y,r float64
}

func main(){
	var c *Circle
	//Note the new(<Type>) actually returns a pointer not a object reference
	c = new (Circle)
	c2 := Circle{x: 0.0, y: 0.0, r: 5.0}
	
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c2.x, c2.y, c2.r)
	
}
