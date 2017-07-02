package main

import "fmt"

func secondary()  {
	fmt.Println("I am secondary")
	// Create array length of 5
	var x [5]float64
	// Fill up every element
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	// set up loop to compute total score
	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	// calculate the average
	// total is float 64 but len(x) by default is int, needs conversion
	fmt.Println(total/ float64(len(x)))

	// we can modify also for loop if we deal with arrays
	var total2 float64 = 0

	// Looping over array
	for k, value := range x {
		fmt.Println(k, value)
	}
	// Go does not let you use variable that you don't use
	// this is why we see underscore (_)there. it is array index placeholder
	for _, value := range x {
		total2 += value
	}
	fmt.Println(total2 / float64(len(x)))
}

func tertiary()  {
	//shorter form to creat array
	x := [5]float64{
		44,
		66,
		89,
		75,
		6,
	}

	for n, value := range x {
		fmt.Println("Index: ", n , "value: ", value)
	}
}

func main()  {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	// print out the secondary function results
	secondary()
	// tertiary function
	tertiary()
}



