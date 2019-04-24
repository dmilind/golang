/*
Array: Data Structure in Go
What is array ?
  Array is a data structure in golang where similar type of data is collected. When I am saying similar kind of data that means, data type has to be identical
  for all the elements in the array.
  When array is initialized , it gets stored in continuous memory location. And all these location can be considered as a block which is represented by the
  identifier used for array.
  For example:
    if we want to store a range of numbers from 1 to 10. We can either use individual variable for storing each value. like a = 1, b =2, c=3 , d=4 etc.
    But how many many variable do we use and how many initialization we do for this ? Other way is to use array. In other word, store all these numbers
    in a continuous memory location and assign an identifier to a block of these memory location.

    The idea of array is to point many objects with a single variable. Array is a fixed number of object. While initializing of array size must be provided.

Syntax:
  var array_identifier = [size_of_array]data_type{object1,object2,...object_size}
  example:
  var a = [10]int{1,2,3,4,5,6,7,8,9,10}
*/
package main

import (
  "fmt"
)

func main(){
  // initialization an array of numbers
  var num = [10]int{1,2,3,4,5,6,7,8,9,10}
  // printing array elements: way 1
  fmt.Println("Displaying array")
  fmt.Println(num)  // this prints all array since num variable is assigned with all elements.
  fmt.Println("Displaying array using indexing")
  fmt.Printf("%d %d %d %d %d %d %d %d %d %d\n" ,num[0],num[1],num[2],num[3],num[4],num[05],num[6],num[7],num[8],num[9])
  // in above statement , all array elements are called by index. For long list this is really difficult to code.
  // For this we can use for loop to go through the each element and print the number
  // Pseudo code
  /*
  1. set index at 0,
  2. print the number
  3. increament the index to 1
  repeat 2 , 3 until we get EOD
  */
  for i :=0 ; i < 10 ; i ++ {
    fmt.Printf("%d\n",num[i])
  }
  fmt.Println("Using len function")
  // other way to do this using standard function from go
  for i := 0 ; i < len(num) ; i++ {
    fmt.Printf("%d\n",num[i])
  }
  fmt.Println("Using range function")
  for _,n := range num {
    fmt.Printf("%d\n", n)
  }
}
