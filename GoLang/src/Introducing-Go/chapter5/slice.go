package main

import "fmt"

func main()  {
	sliceappend()
	copyslice()
}

func sliceappend(){
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
}

func copyslice() {
	slice1 := []int{1,2,3}
	slice2 := make([]int,2)
	// copy (destination, source)
	copy(slice2,slice1)
	fmt.Println(slice2, slice1)
}