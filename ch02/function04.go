/*
Functions in Go
variadic functions

Objective:
    Get all the arguments from the called block and do some operations on data.

Problem:
    Taking huge data of same data type as an argument in function may need to write list of variables in the function. This may slower the code performance.
    Also runtime memory will be busy for variables used in functio.

Example:
    Create a program which can return the sum of all the numbers.
*/

package main

import (
  "fmt"
)

var (

)
func sumAll(n ...int) int {
   total := 0
   for _, num := range n {
     total += num
   }
   return total
}
func main () {
  // call function to add all
  result := sumAll(12,4,5,34,11,755)
fmt.Println("Total:=", result)
}

/*
Explaination:
As usual program starts executing from main function. To understand any program start reading from main function and create a mind map of logic.
A variable "result" is declared and value assigned to it would be the return value of function sumAll. While assigning this value, program control
goes to function sumAll.

In sumAll function:
In function signature, some format is used so that all arguments of same data type can be injected to function. This function is called as variadic function.
only one variable is declared and then three dots (...) are used followed by data type of the variable. While calling this function from main function,
arguments provided as a slice gets injected to function sumAll and all these arguments are referred by single variable "n" of type int.

Now n contains the slice of data. so to get the sum of all numbers , we should take one number at a time and add it "total" , repeat it for second number and add this
to new value of "total". This process will go into the loop until we get EOD. To inplement this logic for loop has been used in the program which does the same thing.

At final "total" will store a value which is a sum of all numbers. And this "total" is returned by the function. Since this function is returning single value,
at the starting of function single data type (int) is typed, also only single value is passed to caller from main function.

In this way "result" displays "total" value. 
a variable is declared "total" and zero value is assigned to it.
*/
