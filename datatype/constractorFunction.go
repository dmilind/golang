/*
Struct: Data structure in GO
constructor function:
  When struct is declared and initialized, default values will be assigned to the data inside struct. In order to assign some value
  instead of default value , constructor function can be used
*/
package main

import (
  "fmt"
)

type Circle struct {
  pi float64
  radius float64
}

func DefaultCircleValues (r float64) Circle {
  c := Circle {
    radius: r,
    pi: 3.14,
  }
  return c
}
func main(){
  // var c Circle
  // output would be default values :
  // {0 0}
  fmt.Println(DefaultCircleValues(7))
  // output
  // {3.14 7}
}
