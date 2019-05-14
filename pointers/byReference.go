/*
Call By Reference
*/

package main

import (
  "fmt"
)

func ByRef (ptr *int) *int  {
  fmt.Println("============================ FROM Function =====================================")
  fmt.Println("Value passed into function is", ptr, "and value at this address is", *ptr)
  fmt.Println("===============================================================================")
  *ptr = 99
  fmt.Println("============================ FROM Function =====================================")
  fmt.Println("After assigning new value to ptr")
  fmt.Println("New value after assignment is", *ptr, "and memory address is", ptr)
  fmt.Println("======================================================================")
  return ptr

}


func main() {
  fmt.Println("============================ Base Value ==============================")
  BaseValue := 11
  fmt.Println("Value from main function is", BaseValue, "and memory address is", &BaseValue)
  fmt.Println()
  mem := ByRef(&BaseValue) // passing memory address instead of actual face value. so that operations would be performed on memory address indirectly.
  fmt.Println("Value from main function is", BaseValue, "and memory address is", mem)
  fmt.Println("======================================================================")
}
