/*
Functions in GO
Defer functions

What is defer statement for functions.
    defer statement is a feature from golang which tells the function to be executed later once surrounded function got executed and returned.
    For exmaple suppose we have 2 functions
    function one ()
    function two()
    now we want function one to be executed after function two. Either we can write to call function two first in main function. But this time we need to
    make sure that proper flow has been maintained.
    Other way we can use defer statement for function one() like defer function one()

    function main(){
    defere function one ()
    function two()
  }

  By using this logic, function two() will get exectuted first then function one()
*/

package main

import (
  "fmt"
)
var (

)
func hello() {
  fmt.Printf("Hello")
}
func world() {
  fmt.Printf(" World !\n")
}

func main () {
  // calling functions hello and world in sequence
  hello()
  world()

  // Calling function world and hello in sequence
  world()
  hello()

  // Now calling world and hello in sequence but using defer in order to get hello world output.
  defer world()
  fmt.Println("second")
  hello()
}
