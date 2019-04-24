/*
Struct: Data structur in GO
==> Ways to initialize struct
*/
package main

import (
  "fmt"
)

type Employee struct {
  Firstname string
  Middlename string
  Lastname string
  Age int
  Height float32
  Address string
  Role string
}

func main(){
  // Create an instance from Person for Employee
  var empOne Employee     // One way of creating instance/initializing a variable for struct
  empOne.Firstname = "Tom"
  empOne.Middlename = "K"
  empOne.Lastname = "Jeff"
  empOne.Age = 51
  empOne.Height = 5.8
  empOne.Address = "89 Cathey Dr, California"
  empOne.Role = "Software Engineer"

  // printing a data for empOne
  fmt.Println(empOne)
  // output:
  // {Tom K Jeff 51 5.8 89 Cathey Dr, California Software Engineer}

  // Crating other instance from Employee struct

//  empTwo := new(Employee) // new function is used to initialize instance for struct Employee
// if we use new method the output would be
// &{Jim B Khatri 59 5.2 San Mateo Dr, California UI Developer} ===> why &
  var empTwo Employee
  fmt.Printf("%+v\n", empTwo) // Method to debug the struct initilization

  // Assigning values to the elements for instace empTwo
  empTwo.Firstname = "Jim"
  empTwo.Middlename = "B"
  empTwo.Lastname = "Khatri"
  empTwo.Age = 59
  empTwo.Height = 5.2
  empTwo.Address = "San Mateo Dr, California"
  empTwo.Role = "UI Developer"

  // printing a data for empOne
  fmt.Println(empTwo)
  // output:
  // {Jim B Khatri 59 5.2 San Mateo Dr, California UI Developer}

  // Crating one more instance from Employee struct
  empThree := Employee {
    Firstname: "Hung",
    Middlename: "Chi",
    Lastname: "Sung",
    Age: 41,
    Height: 5.3,
    Address: "Cornell Ave ,California",
    Role: "Senior Manager II",
  }
  // printing a data for empOne
  fmt.Println(empThree)
  // output
  // {Hung Chi Sung 41 5.3 Cornell Ave ,California Senior Manager II}
}
