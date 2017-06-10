package main

import "fmt"
/*
	Create function to calculate average.
	The average function will need to take in a slice of float64s and return one float64.
*/
func average(xs []float64) float64  {
	// Define variable total to calculate sum of slice elements
	total := 0.0
	// Looping over slice to calculate sum of elements, _ is placeholder
	// for index parameter that we don't need
	for _, v := range xs{
		total += v // total = total + v
	}
	/*
	The return statement causes the function to immediately stop and return the value
	after it to the function that called this one.
	*/
	return total / float64(len(xs))
}

// Functions can return more than one value
func morethanonevalue() (int, int) {
	return 5 ,6
}

func main()  {
	// Define slice
	xs := []float64{98,93,97,82,83}
	fmt.Println(xs)
	fmt.Println("Average of these numbers are", average(xs))

	// Here we call function to return more than one value
	x, y := morethanonevalue()
	fmt.Println(x, y)

	// notice that we are not using colon because we already have variable.
	//We just overwrite it
	x = 0

	// It is possible to create local functions like function inside function
	// Local function can access local variables
	increment := func() int{
		x++ // add 1 to variable
		return x
	}
	fmt.Println("First incremental", increment())
	fmt.Println("Second incremental",increment())
}