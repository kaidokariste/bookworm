// Person struct with methods of pointer receiver

package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName, LastName string
	Dob                 time.Time
	Email, Location     string
}

//A person method with pointer receiver
func (p *Person) PrintName() {
	fmt.Printf("\n%s %s\n", p.FirstName, p.LastName)
}

//A person method with pointer receiver
func (p *Person) PrintDetails() {
	fmt.Printf("[Date of Birth: %s, Email: %s, Location: %s ]\n", p.Dob.String(), p.Email, p.Location)
}

//A person method with pointer receiver
func (p *Person) ChangeLocation(newLocation string) {
	p.Location = newLocation
}

//A person method with nonpointer receiver (missing *)
func (p Person) ChangeFirstName(newFirstName string) {
	p.FirstName = newFirstName
}

func main() {
	p := &Person{ // to call pointer receiver, the instance should be created with & sign
		"Shiju",
		"Varghese",
		time.Date(1979, time.February, 17, 0, 0, 0, 0, time.UTC),
		"shiju@email.com",
		"Kochi",
	}
	// Location is changed because it changes the name
	p.ChangeLocation("Santa Clara")
	// Although the function exists, but without pointer receiver and the name is not changed
	p.ChangeFirstName("Kaido")
	p.PrintName()
	p.PrintDetails()
}
