package main

import "fmt"

type Person struct {
	Name string
}

// Method talk
func (p Person) Talk() {
	fmt.Println("Hi my name is", p.Name)
}

/*
Go supports relationships like this by using embedded types
(sometimes also referred to as anonymous fields
*/
type Android struct {
	Person
	Model string
}

func main()  {
	a := new(Android)
	a.Name = "Kaido"
	// the Persons struct can be accessed using the type name
	a.Person.Talk()
	// or shorter form
	a.Talk()
}