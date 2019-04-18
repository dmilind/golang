/*
Slice: Data structure in Go

What is slice ?
  Slice is a continuous segment of underlying array. So it can be subset of array. Most of the time slices are used in golang. Array has some limitations.
  Once array is initialized , you already fixed the size , now adding , deleting, any element from this array is not possible. To get this done, slice can
  be helpful. All such operations can be performed on slices since size can be varied in slice.

Syntax:
  var slice_identifier = []data_type{element,element...}

* Size is not mentioned while initializing the slice.
* In golang you can use make function to initialize a slice.
var slice_identifier = make([]int, size)
Value can be assigned to elements from the slice.
*/
package main

import (
  "fmt"
)

func main(){
  // initialize a slice
  var nameSlice = make([]string, 3)
  fmt.Println(nameSlice) // Empty slice has been initialized
  // assigning the values to the slice
  nameSlice[0] = "This"
  nameSlice[1] = "is"
  nameSlice[2] = "slice"
  fmt.Println(nameSlice[0],nameSlice[1],nameSlice[2])

  // using for loop to print the same
  for i:=0; i<3; i++ {
    fmt.Printf("%s ",nameSlice[i])
  }
 // Working on slice : append
 nameSlice = append(nameSlice, "from", "gopher")
 fmt.Println("\n",nameSlice)

 // Adding different type of data in the slice
 nameSlice = append(nameSlice, "200") // this means we can add other type of data_type in slice
 fmt.Println("\n",nameSlice)
}
