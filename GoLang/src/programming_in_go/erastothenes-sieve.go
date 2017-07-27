package main

import "fmt"

// Erathosthenes famous algorithm for finding prime numbers
// Returns intgere slice and total count of primes
func erathosthenes_sieve(limit_number int64) ([]int64, int64) {

	// Define base slice with first prime number 2
	base_slice := []int64{2}

	// Next we loop trough all numbers starting from 2 till user given limit
	for i := int64(2); i <= limit_number; i++ {

		// Loop trough the slice
		for index, number := range base_slice {
			// For every iteration we determine the lenght of current slice
			ln_of_slice := len(base_slice)

			if i % number == 0 {
				//If the remaining of the iteration number and some slice number is 0
				// then this number is not a prime
				break
			} else {
				// If the remaining has not been 0 and we have reached end of the slice
				// then current number is prime and we add it at the end of slice
				// looping index starts from 0, slice length count starts from 1
				if index + 1 == ln_of_slice {
					base_slice = append(base_slice, i)
				}

			}

		}

	}
	// Return slice and the count of them
	return base_slice, int64(len(base_slice))

}


func main()  {
	prime_numbers, total := erathosthenes_sieve(99999)
	fmt.Println("Prime numbers", prime_numbers, "total of them", total)

}