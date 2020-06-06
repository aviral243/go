package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	aviral := person{
		firstName: "Aviral",
		lastName:  "Gangwar",
		contactInfo: contactInfo{
			email:   "aviralgangwar24@gmail.com",
			pincode: 243122,
		},
	}
	aviral.updateName("avi")
	aviral.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
