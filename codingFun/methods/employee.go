/*
Using method to find all active employees from a company
*/

package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	empId     int
	age       int
	contact   int64
	active    bool
}

func (e Employee) CompanyEmployee() {
	// return only active employee data
	if e.active {
		fmt.Println("Active Employee :", e.firstName, e.lastName)
	} else {
		fmt.Println("Inactive Employee :", e.firstName, e.lastName)
	}
}

// creating an object from struct
func main() {
	var eone Employee
	eone = Employee{
		firstName: "John",
		lastName:  "Doe",
		empId:     101,
		age:       21,
		contact:   1231231234,
		active:    true,
	}

	etwo := Employee{
		firstName: "Jim",
		lastName:  "Victor",
		empId:     103,
		age:       27,
		contact:   1211211212,
		active:    false,
	}
	ethree := Employee{
		firstName: "Mogambo",
		lastName:  "Pat",
		empId:     999,
		age:       78,
		contact:   9191919191,
		active:    true,
	}
	efour := Employee{
		firstName: "Victor",
		lastName:  "Pande",
		empId:     312,
		age:       47,
		contact:   7867865642,
		active:    false,
	}

	eone.CompanyEmployee()
	etwo.CompanyEmployee()
	ethree.CompanyEmployee()
	efour.CompanyEmployee()

}
