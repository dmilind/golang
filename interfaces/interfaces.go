/*
Interfaces in Go
What is interface ?
  Interfaces is little hard to understand. Better way to understand it from the old program that we did in methods. So lets refer to methods.go program.
  From that program we try to understand why do we need interface.
  So interface is a collection of method set. What that means ? Consider a data type which encapsulates set of methods. Those methods belongs to some type, lets
  say struct. So indirectly those structs implements interface. When instance is created from that struct. That instance can implement interface, and
  set of methods can be applied to that instance.

  Hope this is understood now. For more details lets walk through the program by comparing wih methods.go program

Syntax :
  type Interface_indetifier interface {
    methods_identifier01(arguments) return_data_type
    methods_identifier02(arguments) return_data_type
    methods_identifier03(arguments) return_data_type
    methods_identifier04(arguments) return_data_type
}
*/
// from last program methods.go , area of circle and area of rectangle is calculated using methods. In this program we add one more method to calculate perimeters

package main

import (
  "fmt"
//  "log"
  "math"
)

// declare interface
type Geometry interface {
  // set of methods:
  Area() float64
  Perimeter() float64
}

// define struct
// struct for circle
type Circle struct {
  Radius float64
}
// struct for Rectangle
type Rectangle struct {
  Length float64
  Width float64
}
// define methods
// method for calculating area of circle
func (c Circle) Area() float64 {
  area := math.Pi * c.Radius * c.Radius
  return area
}

// method for calculating perimeter of cirle
func (c Circle) Perimeter() float64 {
  perimtr := 2 * math.Pi * c.Radius
  return perimtr
}

// method to calculate area of rectangle
func (r Rectangle) Area() float64 {
  area := r.Length * r.Width
  return area
}

// method to calculate perimeter of rectangle
func (r Rectangle) Perimeter() float64 {
  perimtr := (2 * r.Length) + (2 * r.Width)
  return perimtr
}
func Calculate(g Geometry) (geometryArea, geometryPerimeter float64) {
   geometryArea = g.Area()
   geometryPerimeter = g.Perimeter()
   return geometryArea, geometryPerimeter
}
func main () {
  // initialize structs
  c := Circle {
    Radius: 8.9,
  }
  r := Rectangle {
    Length: 3.4,
    Width: 7.1,
  }
  // Calling function calculate which will call Geometry interface
  cicleArea, circlePerimeter := Calculate(c)
  rectangleArea, rectanglePerimeter := Calculate(r)
  fmt.Printf("Area and Perimeter of a circle is %f and %f respe\n", cicleArea, circlePerimeter )
  fmt.Printf("Area and Perimeter of a rectangle is %f and %f\n", rectangleArea, rectanglePerimeter)
}
// output
//Area and Perimeter of a circle is 248.845554 and 55.920349 respe
//Area and Perimeter of a rectangle is 24.140000 and 21.000000a
