package main

import (
	"fmt"
	"math"
)

type circle struct {
	color  string
	radius float64
}

type rectangle struct {
	color  string
	width  float64
	length float64
}

//////////////////////////
//// Definging method ////
//////////////////////////

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.length
}

//////////////////////////

// for defining an interface, just like structs, but instead of defining fields we define a "method set"
type shape interface {
	area() float64
}

func returnShapesArea(s shape) float64 {
	return s.area()
}

func main() {

	// 2 ways of using interfaces

	var c1 shape = circle{color: "red", radius: 12}
	var r1 shape = rectangle{color: "blue", width: 123, length: 34}
	fmt.Println(returnShapesArea(c1), c1.area())
	fmt.Println(returnShapesArea(r1), r1.area())

	c2 := circle{color: "red", radius: 12}
	r2 := rectangle{color: "blue", width: 123, length: 34}
	fmt.Println(returnShapesArea(c2))
	fmt.Println(returnShapesArea(r2))
}
