package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jr := person{
		firstName: "José",
		lastName:  "Rodríguez",
		contactInfo: contactInfo{
			email:   "jrm2087@gmail.com",
			zipCode: 11101,
		},
	}
	//jrPointer := &jr // memory address
	jr.updateName("Darío")
	jr.print()
}

// * value
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
