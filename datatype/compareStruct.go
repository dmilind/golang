/*
Struct : Data structure in Go
Comparing two structs:
  Instances of the structs can be compared for equality or inequality.
*/
package main

import (
  "fmt"
  "reflect"
)

type Drinks struct {
  Name string
  Ice bool
}

type HardDrinks struct {
  Name string
  Ice bool
}
func main(){
  a := Drinks {
    Name: "Coke",
    Ice: true,
  }
  b := Drinks {
    Name: "Pepsi",
    Ice: false,
  }
  c := Drinks {
    Name: "Coke",
    Ice: true,
  }
  d := HardDrinks {
    Name: "Whisky",
    Ice: true,
  }
  // Compare instances
  if a == b {
    fmt.Println("Instaces a and b are identical")
  } else if a == c {
    fmt.Println("Instances a and c are indentical")
  } else if b == c {
    fmt.Println("Instaces b and c are identical")
  } else {
    fmt.Println("Looks like nothing is matching")
  }
  // checking type of the instances using reflect package
  fmt.Println("Type of instance a is", reflect.TypeOf(a))
  fmt.Println("Type of instance a is", reflect.TypeOf(b))
  fmt.Println("Type of instance a is", reflect.TypeOf(c))
  fmt.Println("Type of instance a is", reflect.TypeOf(d))

  // comparing
  if a ==b {
    fmt.Println("Instaces a and d are same")
  } else {
    fmt.Println("Instances a and d are not identical")
  }
}
