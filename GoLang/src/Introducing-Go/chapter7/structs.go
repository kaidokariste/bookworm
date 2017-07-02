package main

import (
	"fmt"
	"math"
)

// define new type Circle and it is struct type
type Circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c *Circle) float64 { // using without a pointer func circleArea(c *Circle) float64
	c.r = c.r * 2
	fmt.Println("Function modified variable", c.r)
	return math.Pi * c.r*c.r
}


func main()  {
	// Initialization
	c := Circle{x: 0, y: 0, r: 6}
	fmt.Println("Original variable", c.r)
	fmt.Println(circleArea(&c)) // using without a pointer:  fmt.Println(circleArea(c))
	fmt.Println("Still original variable", c.r)
}

// Results without pointer
/*
Original variable 6
Function modified variable 12
452.3893421169302
Still original variable 6
*/

// Results with pointer
/*
Original variable 6
Function modified variable 12
452.3893421169302
Still original variable 12*/

// We can see that without pointer, the arguments are only copied. Even when they are modified inside function
// they are not changed inside defined struct. Using pointers, they are.