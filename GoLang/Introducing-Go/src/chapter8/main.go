package main

import (
	"fmt"
	// packages can have aliases
	// now we can use go main package "math"
	m "chapter8/math"

)


func main()  {
	xs := []float64{1,2,3,4}
	avg := m.Average(xs)
	fmt.Println(avg)
}

