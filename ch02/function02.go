/*
Functions in Go
Program which finds the number is odd or even
*/

package main

import (
  "fmt"
)
var (
  numOne int
)
func finder (x int) bool {
  fmt.Println("Using this finder function, evaluating number", x)
  return x % 2 == 0
}

func main (){
  numOne = 13
  fmt.Println("Provided number is", numOne)
  truthValue := finder(numOne)
  fmt.Println("Given number is even, true or flase ? So answer is, it is", truthValue)
}

/*
Expected output:
Provided number is 13
Using this finder function, evaluating number 13
Given number is even, true or flase ? So answer is, it is false
*/
