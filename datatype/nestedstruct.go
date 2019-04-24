/*
Struct: Data structure in golang
Nested Struct:
  Some times data structure can be complex, in this case creating simple struct can be difficult. Complex data structure means data within data.
  Example of this complex data structure can be data in xml file , json file. To organize this data to be injected to program for further processing,
  nested stuct can be used. Nested struct means one struct is referenced in another super struct.

Example:
type superStruct struct {
  data_field data_type
  data_field data_type
  struct_filed structName
}

type structName struct {
  data_field data_type
  data_field data_type
}

In above example , structName is a part of data stracture superStruct. That means all complex data can be initialized bu using superStruct only. So data will be
modeled in a way where the data from structName will be sorted first, then this data will be passed to superStruct. Now at this level whole data can be
referenced and used in further processing in program.
*/
package main

import (
  "fmt"
)
// declare struct of Employee where each employee has name, phone number and address. But this address is has apt number, street name , state and zip code.
// If we see this is kind of complex data which is difficult to manage by using simple struct. Other way we can say that address is structure in innner
// elements(data) of different types. To get different type of data, we can write a separate struct for address. And since address is a part of employee
// we have to refer this struct in Employee struct which is a superstruct.

type Employee struct {
  Fullname string
  Contact int
  Address EmployeeAdress           // Struct EmployeeAdress is nested
}

type EmployeeAdress struct {
  Apt int
  Street string
  State string
  Zip int
}

// consider above data can be in xml format as below (example only to understant. Getting data from xml file is different topic where xml functionalities required
/*
<employee>
<fullname>Allen Walker</fullname>
<cotact>1234567890</contact>
<address>
  <apt>1680</apt>                    // this is data within data, data of address where address is a data of employee :) [Hope this is clear]
  <street>Cornell Drive</stree>
  <state>california</state>
  <zip>98989</zip>
</address>
*/

func main(){
  var emp Employee
  fmt.Printf("%+v\n",emp)  // To check that struct has been initialized with default values properly.
  // output:
  // {Fullname: Contact:0 Address:{Apt:0 Stree: State: Zip:0}} This confirm that EmployeeAdress struct got embedded while initializing Employee struct.
  // lets start using struct to create an instances from Employee struct
  emp = Employee {
    Fullname: "Allen Walker",
    Contact: 1234567890,
    Address: EmployeeAdress{
      Apt: 1680,
      Street: "Cornell Drive",
      State: "California",
      Zip: 98989,
    },
  }
  fmt.Println(emp)  // emp already initialized
  // output:
  // {Allen Walker 1234567890 {1680 Cornell Drive California 98989}} // This specifies that nested struct worked here.
  // lets call each element of struct using . (dot) notation.
  fmt.Println("Employee Name =", emp.Fullname)
  fmt.Println("Contact Number =", emp.Contact)
  fmt.Println("Full Address =", emp.Address)  // This return the whole address like {1680 Cornell Drive California 98989}
  fmt.Println("Apt Number =", emp.Address.Apt)  // dot notation till we get embedded element from embedded struct.
  fmt.Println("Stree Name =", emp.Address.Street)
  fmt.Println("State =", emp.Address.State)
  fmt.Println("Zip =", emp.Address.Zip)

  // output for above line
  /*
  Employee Name = Allen Walker
  Contact Number = 1234567890
  Full Address = {1680 Cornell Drive California 98989}
  Apt Number = 1680
  Stree Name = Cornell Drive
  State = California
  Zip = 98989
*/
}

/*
Program Output:
{Fullname: Contact:0 Address:{Apt:0 Street: State: Zip:0}}
{Allen Walker 1234567890 {1680 Cornell Drive California 98989}}
Employee Name = Allen Walker
Contact Number = 1234567890
Full Address = {1680 Cornell Drive California 98989}
Apt Number = 1680
Stree Name = Cornell Drive
State = California
Zip = 98989
*/
