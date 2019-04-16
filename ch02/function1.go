/*
Functions in go
Understanding functions:
Functions are small code blocks, or chunks of code which can be reused any where in the upcoming program. Consider you are writing a program for something.
In this program a block of code might get repetate, you observe that same block of code appears saveral times. Writing same line of code at different location
might be time consuming also it increases a length of the program. To avoid this bad practice, we should hava some entity where repetative or block of code can
be written at one point and we can refer this entity in main program. So this entity would be called as a function.

Functions can be called as building blocks of a programming language. Code maintainance becomes easy by using functions.
In golang functions can be written as below.

func example (arguments_name argument_data_type) function_return_data_type {
  all function logic
  ...
  ...
  ...
  return
}

func keyword tells that this is function. Functions can be injected with arguments. Basically if functions is called by some code, mostly some arguments
will be passed to be processed in the function. For that reason (arguments_name argument_data_type) is used. Since golang is strongly typed language so
these arguments must be declared with data type.

Every function will process the data and returns some value to the caller. So what type of data will be returned , that has to be declared like function_return_data_type
If data type is not matching with the return value then compilation error occures.

At the end , return keyword is used to return the value to the caller.

Now lets try to understand below code.
*/

package main

import (
  "fmt"
)

func addNum(x int, y int) int {
  fmt.Println("I am in function addNum")
  var result int  = x + y
  fmt.Println("Addition is performed and output is stored in result which is of type int, passing the result to the main function")
  return result
}
func main() {
  fmt.Println("Using functions! I am in main function")
  fmt.Println("Creating a function to add provided numbers and get the result value")
  numOne := 23
  numTwo := 11
  finalValue := addNum(numOne, numTwo) // Calling function addNum. Function addNum is returning a single value that value finalValue is used. And value returned
                                      // by it is of type int so finalValue 's data type would be int.
  fmt.Println("I am in main function")
  fmt.Println("Final result is", finalValue)
}

/*
Expected output:
Using functions! I am in main function
Creating a function to add provided numbers and get the result value
I am in function addNum
Addition is performed and output is stored in result which is of type int, passing the result to the main function
I am in main function
Final result is 34
*/
