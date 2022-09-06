package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Lee",
		contact: contactInfo{
			email: "jim@exmpale.com",
			zipCode: 94000,
		},
	}

	//jimPointer := &jim
	//jimPointer.updateName("Jimmy2")
	//or
	jim.updateName(("Jimmy"))
	jim.print()

	fmt.Println("\n--------------------")
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

func updateSlice(slice []string) {
	slice[0] = "Bye"
}

func (pointer *person) updateName(newFirstName string) {
	(*pointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}