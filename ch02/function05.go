/*
Functions in GO
Functions returning named variables.

Objective :
    Return actual named variables from the function.

*/

package main

import (
  "fmt"
)

var (

)

func namedReturn()(f, m, l string) {
  f = "Milind"
  m = "A"
  l = "Dhoke"
  return
}
func main () {
  // take values from function and assign to variables here.
  firstName, middleName, lastName := namedReturn()
  fmt.Printf("My name is %s %s %s\n",firstName, middleName, lastName)
}

/*
Expected Output:

My name is Milind A Dhoke

*/
