package main

import "fmt"

// Variadic function, the number of arguments is not fixed
func add(args ...int) int {
	// Define variable total, initial valu 0
	total := 0
	// Loop trough arguments and sum values together
	for _, v := range args {
		total += v // total = total + v
	}
	return total //return the end result
}

func main()  {
	fmt.Println(add(1, 2, 3))
	// We can also pass a slice of ints by following the slice with an ellipsis:
	xs := []int{3,4,3}
	fmt.Println(add(xs...))
}
