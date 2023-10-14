package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson"}
	fmt.Println("--------------------------")
	fmt.Println(alex)
	fmt.Println("--------------------------")
	fmt.Printf("%+v", alex)
	fmt.Println("--------------------------")
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: "78689",
		},
	}
	fmt.Println("--------------------------")
	fmt.Println(jim)
	fmt.Println("--------------------------")
	fmt.Printf("%+v", jim)
	fmt.Println("--------------------------")
    jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
	fmt.Println("--------------------------")
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string){
(*pointerToPerson).firstName = newFirstName
}
