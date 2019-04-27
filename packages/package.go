/*
Package in Golang
Packages are something that are being called in the go pragram which does some functionality on the code.
Packages contains the functions which can be used in our program instead of writing them from the scratch. It is kind of plugin program from the open
community or standard golang library. Library is collection of all such packages.

In out example , we will create a package math. This math is different from actual standard math package. Package creation is little different than
executable program in golang. Executatino golang uses main package and main function , while package neither use main package not main function.
A Package is a collection of function. Each function is responsible for doing somethis.
*/

package main

import (
  "fmt"
  "practice/golang/packages/math"     // importing math package from go path.this would be under pkg folder. math is a object file
)

func main () {
  // We want to find out the average of two different numbers , we can create a function here to do so. If we want this function to be used in other golang code,
  // which is out if the scope of this file, then it is not possible. So to do so we can make that function available in math package.
  // take two numbers
  a := 8.0
  b := 12.0
  // call a function to find an average of these numbers
  average := math.Average(a, b)
  fmt.Println("Average of these two numbers is", average)
  // At this point I have not written a package named math, so if you run this program , you will find an error:
  // package.go:16:3: cannot find package "src/practice/golang/packages/math" in any of:
      //  /usr/local/go/src/src/practice/golang/packages/math (from $GOROOT)
      //  /Users/mdhoke/go/src/src/practice/golang/packages/math (from $GOPATH)
  // lets create a package
  // Output of above code :
  //Calculating Average:  this line is from the package
  // Average of these two numbers is 10 // return value is coming from package

  sum := math.Sum(a, b)
  fmt.Println("Sum of two numbers is", sum)
  // output
  // Calculating Average:
   // Average of these two numbers is 10
  // Sum of two numbers is 20
  subtract := math.Subtract(a,b)
  fmt.Println("Subtraction of two numbers is", subtract)

}
