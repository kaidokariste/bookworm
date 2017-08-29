package main

import "fmt"

func doPanic(){
	// anonymous function has been added to the deferred list, in which recover is called
	// to regain control from the panicking function
	defer func(){
		// Recover is typically used inside deferred function that regain control
		// of a panicking function
		if e:= recover(); e != nil {
			fmt.Println("Recover with: ", e)
		}
	}()
	panic("Just panicking for the sake of demo")
	fmt.Println("This will never be called")
}

func main(){
	fmt.Println("Starting to panic")
	doPanic()
	fmt.Println("Program regains control after panic recover")
}
