package main

import (
	"math"
	"fmt"
)

// I had to change it to Circle2 because Circle1 is in structs
// file and Go sees it package wide
type Circle2 struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// special type of function called method
// (c *Circle) is called a receiver
func (c *Circle2) area() float64{
	return math.Pi * c.r * c.r
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main()  {
	var R Circle2 // create an instance of our new Circletype, by default set to zero.
	R.r = 5 // Set radius value from zero to 5
	fmt.Println(R) // Print out Circle values {0 0 5}
	fmt.Println(R.area()) // Print Circle area 78.53981633974483
	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())
}

