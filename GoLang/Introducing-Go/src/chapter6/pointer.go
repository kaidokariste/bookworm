package main

import "fmt"

// Usual case
/*
func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // x is still 5
}*/

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

/*Pointers reference a location in memory where a value is stored rather than the value
itself. By using a pointer (*int), the zero  function is able to modify the original vari‐
able

In Go, a pointer is represented using an asterisk (*) followed by the type of the stored
value. In the zerofunction, xPtris a pointer to an int.
*/