package math

/*
To see documentation, run godoc cmd//chapter8/math Average in terminal
Comment befor function adds the documentation
*/

// Function to calculate average
func Average(xs []float64) float64  {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}