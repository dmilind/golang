/*
Call By Value
*/
package main

import (
  "fmt"
)

func ByVal (digit int) int {
  fmt.Println("============================ OLD =====================================")
  fmt.Println("Value passed into function is", digit, "and memory address is", &digit)
  fmt.Println("=====================================================================")
  // now changing the value
  digit = 99
  fmt.Println("============================ NEW =====================================")
  fmt.Println("After assigning new value to digit")
  fmt.Println("Value passed into function is", digit, "and memory address is", &digit)
  fmt.Println("======================================================================")
  return digit
}

func main () {
  fmt.Println("============================ Base Value ==============================")
  BaseValue := 11
  fmt.Println("Value from main function is", BaseValue, "and memory address is", &BaseValue)
  BaseValue = ByVal(BaseValue)
  fmt.Println("Value from main function is", BaseValue, "and memory address is", &BaseValue)
  fmt.Println("======================================================================")

}
