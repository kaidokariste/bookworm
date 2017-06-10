package main

import "fmt"

// Recursion means that function can call itself
// best example is factorial 4! = 1*2*3*4
func factorial(x uint) uint  {
	if x == 0{
		return 1
	}
	return x * factorial(x-1)
}

func main()  {
	fmt.Println(factorial(2))
}

/*
1. Is x == 0? No (xis 2).
2. Find the factorial of x − 1
	a. Is x == 0? No (xis 1).
	b. Find the factorialof x − 1.
		i. Is x == 0? Yes, return 1.
	c. Return 1 * 1.
3. Return 2 * 1.
*/