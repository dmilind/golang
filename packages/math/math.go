/*
Custom math package in go
*/

package math // package has name at this line , we dont use main package since this is not executable program

import (
  "fmt"
)
// function to find the average of two numbers .. Note the cap letter used in function name which makes this function available out of scope.
func Average(arg1 float64, arg2 float64) float64 { // take arguments from the main go program and returns average value in float64
  fmt.Println("Calculating Average:")
  avg := (arg1 + arg2)/2
  return avg
}
func Sum(s1, s2 float64)(total float64){
  total = (s1 + s2)
  return
}
func Subtract(arg1, arg2 float64) float64 {
  subtract := (arg1 - arg2)
  return subtract
}
// if you run this program , compiler error out as
// go run: cannot run non-main package

/*
Now this package will provide functions. For example Average function will be imported to main package.go file.
To use this package , we need to create a object file in the go path. To do so run "go install" this command will generate a object file under $GOPATH/pkg/.../.../math.a
this file will be imported ,means indirectly go functionality is imported.

*/
