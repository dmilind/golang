/*
Methods With Pointers in Go.
Explaination:
When data type is defined. Method can be written for it. receiver takes a value from the data type. This value can be pass by value or pass by reference.
If receiver take value using pass by value, then copy of data type is created and all operations are performed on the copy not on the actual data type.
If you want to work on actual data type then receiver value has to be passed using pass by reference. In pass by reference pointer is used which points the memory
address of the data type.

Syntax:
  Pass by value:
  func (variable data_type) method_identifier (arguments) return_data_type {
  ...
  }

  Pass by reference:
  func (variable *data_type) method_identifier (arguments) return_data_type {
  ...
  }

Example:
  Write a program which can find the area of the triagle using method. Use both pass by value and pass by reference concepts.
*/
package main

import (
  "fmt"
)

type Triangle struct {
  Base float64
  Height float64
  }

// method using pass by value ==> value is received into the method.
func (t Triangle) Area() float64 {
  // this methods returns the area of triangle
  area := 0.5 * (t.Base * t.Height)
  return area
}
func (t Triangle) AreaWithNoPointer(b float64) float64 {
  // trying to change the value of base of triagle
  t.Base = b
  area := 0.5 *(t.Base * t.Height)
  return area
}
func (t *Triangle) AreaWithPointer(b float64) float64 {
  // change the value of base
  t.Base = b
  area := 0.5 * (t.Base * t.Height)
  return area
}
func main(){
  // initialize struct
  T1 := Triangle {
    Base: 7,
    Height: 3,
  }
  // call method to calculate area of triangle t1
//  fmt.Println("Area of triangle T1 is", T1.Area())
  // output :
  // Area of triangle T1 is 10.5
  // now print actual values of T1 triangle
  fmt.Printf("Triangle T1's base is %f and Height is %f\n", T1.Base, T1.Height)
  // Try to change the value of base now.
  fmt.Printf("Area of triagle is %f\n", T1.AreaWithNoPointer(19))
  // Triangle T1's base is 7.000000 and Height is 3.000000
  // Area of triagle is 28.500000
  /*
  if you see second area method , Base value has been changed, but still in the output base value is taken as 7 instead of 19. Why is this so ?
  When value is being passed to method, the changes has been took place on the copy of the struct. But actual struct remained same.In order to change the value
  of original struct then pointer should be used.
  */
  fmt.Printf("Area of triagle is %f\n", T1.AreaWithPointer(19))
  fmt.Printf("Triangle T1's base is %f and Height is %f\n", T1.Base, T1.Height)
  // output:
  // Triangle T1's base is 19.000000 and Height is 3.000000

  // Here we have changed the value of Base from struct. Pointer directly refer to memory location.
}
