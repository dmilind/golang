/*
Using veriadic function
*/

package main

import (
  "fmt"
  "os"
  "reflect"
  //"strconv"
)

func Add(num ...int) int {
  fmt.Println(reflect.TypeOf(num))
  Sum := 0
  for _, v := range num {
    Sum = Sum + int(num[i])
  }
  return Sum

//  return num
}

func main(){
  result := Add(os.Args)
  fmt.Println(result)


}
