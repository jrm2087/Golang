package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	jr := person{
		firstName: "José",
		lastName:  "Rodríguez",
		contact: contactInfo{
			email:   "jrm2087@gmail.com",
			zipCode: 11101,
		},
	}

	fmt.Printf("%+v", jr)
}

//func main() {
//	jose := person{"José", "Rodríguez"}
//	jose := person{firstName: "José", lastName: "Rodríguez"}
//	fmt.Println(jose)
//
//	var jose person
//	jose.firstName = "José"
//	jose.lastName = "Leal"
//
//	fmt.Println(jose)
//	fmt.Printf("%+v", jose)
//}
