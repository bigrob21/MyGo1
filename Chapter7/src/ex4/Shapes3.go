package main

//import "fmt"

//Go Interface which is a collection of signatures
type Shape interface {
	area() float64
}

//Go Struct which is analogous to holder of fields.
type Rectangle struct {
	x1, y1, x2, y2 float64
}

//Go Struct which is analogous to holder of fields.
type Circle struct {
	x, y, r float64
}

//Go Struct which is analogous to holder of fields.
type MultiShape struct {
	shapes []Shape
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _,s := range shapes {
		area += s.area()
	}
	return area
}

func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

func main(){ }

// func main() {
// 	multiShapes := MultiShape {
// 		shapes : []Shape {
// 			Circle{0,0,5},
// 			Rectangle{1,2,20,20},
// 		},
// 	}
// 	for _, v := range multiShapes.shapes {
// 		fmt.Printf("Shape calculated area is %d\n", v)
// 	}
// }