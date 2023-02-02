package main

import "fmt"

type contactInfo struct {
	phone   int
	email   string
	address string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	ian := person{
		firstName: "Ian",
		lastName:  "Baliawa",
		contact: contactInfo{
			phone:   787444814,
			address: "Kawaala",
			email:   "inabalijjawa@gmail.com",
		},
	}

	ian.updateFirstName("puff")
	ian.print()

	age := 12

	fmt.Println("address ", &age)

}

func (p person) print() {
	fmt.Printf("%+v \n", p)
}

func (p *person) updateFirstName(newFirstname string) {
	(*p).firstName = newFirstname
}
