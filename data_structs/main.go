package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName, lastName string
	contact             contactInfo
}

func main() {

	alex := person{"Alex", "Anderson", contactInfo{"jim@jim.com", 12345}}
	alexPointer := &alex
	alexPointer.updateName("Jim")
	alex.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {

	fmt.Println(p)
}
