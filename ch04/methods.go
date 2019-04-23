/*
Methods in GO:
What is method ?
  This is little difficult topic to understand.
  To understand methods in golang,we need to remember OOP concept. In OOP , class encapsulates data and functions which belongs together. To use the function
  from that class, and instance is initialized, this instance uses the data and functions from the class. This function is called as a method.

  In golang we dont have OOP concept. We dont have classes as such. But data type struct serves the function like classes. Instances can be created from this
  struct. To get similar feature with OOP, functions are called using instances. This function is called as method. Method acts as behaviour to the data type.
  This method inherits the data element from the data type. Method is similar to functions apart from passing receiver
  value just right after func keyword.

  When any method is linked to a struct, all instances of that struct can avail this method. This method can inherits all the data fields from the struct.

Why dont we use functions here then ?
  lets take an example to undestand this topic.
  First we will try to achieve the goal using function , later we will use methods and will see the difference.

Syntax:
  func (instance data_type) method_identifier (arguments) return_data_type () {
  ...
}
*/
package main

import (
  "fmt"
)

const pi float64 = 3.14

type Rectangle struct {
  Length float64
  Width float64
}

type Circle struct {
  Radius float64
}
// Part I: With functions
// func AreaOfRectangle(l, w float64) float64 {
//   a := l * w
//   return a
// }
//
// func AreaOfCircle(p, r float64) float64 {
//   a := p * r * r
//   return a
// }

// func main(){
//   // initialize struct Rectangle and Circle
//   rect := Rectangle {
//     Length: 3.12,
//     Width: 7.123,
//   }
//
//   cir := Circle{
//     Radius: 6.879,
//   }
//
//   // Calculating area of rectangle and circle:
//   fmt.Println(AreaOfRectangle(rect.Length, rect.Width))
//   fmt.Println(AreaOfCircle(pi, cir.Radius))

// Above program is writter using function. This functions are Calculating the area of rectangle and circle nicely. To get the area of both of these , arguments had
// to passed to the functions.
// Part II : With methods.
func (r Rectangle) AreaOfRectangle() float64 {
  area := r.Length * r.Width
  return area
}
func (c Circle) AreaOfCircle() float64 {
  area := pi * c.Radius * c.Radius
  return area
}

func main(){
  c := Circle {
    Radius: 6.879,
  }
  fmt.Printf("Area of circle is ==> %f\n", c.AreaOfCircle())  // No need to pass element value from the stuct.
  rectangle := Rectangle {
    Length: 3.12,
    Width: 7.123,
  }
  fmt.Println("Area of rectangle is ==>", rectangle.AreaOfRectangle()) // No need to pass element value from the stuct.
  // output
  // Area of circle is ==> 148.586813
  // Area of rectangle is ==> 22.223760000000002
}

/*
In this example two separate methods has been used. We can use same named methods also, which is not possible with functions.
In above program , struct is defined for rectangle and circle.
In part I, functions are being used to calculate the area of rectangle and circle. These functions take arguments , lenght and width for rectangle, and radius
for circl. Return value from these functions is area of the rectangle or circle. When instance of struct initialized, values are passed to function using dot
notation. These values gets injected into function, and area values is returned. This serves out purpose to calculate the area.

To get area of another instance of struct, again parameter values has to be injected in the function every time. This might increase the coding work.

In part II, methods have been used. These methods inherits the data from the struct once receiver of data type is augmented. here we dont need to pass the arguments,
automatically the values will be injected in to the methods and return value is provided. IMO this is the differnce between methods and functions.
*/
