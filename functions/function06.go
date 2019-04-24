/*
Functions in Go

Recursive functions
    Recursive functions calls same functions from within the function until certain condition is satisfied.

Example:
    Write a program using a recursive function which cound the number of vowels and list of all those vowels

Pseudo Code:

package
import
var
func main () {
  Store a string into a variable
  pass this variable to a function
  get number of vowels used and list of those
}
func vowel (argument string) (count int, vowel []rune) {
  get first letter from string match it with a,e,i,o,u
    if match is ok
      add 1 count to count
      store the vowel in an array
      increament the pointer
      call the same function
    else
      exit from function
  return count , and array

}
*/

package main

import (
  "fmt"
)

var (
  word string
)

func vowelFinder(inputWord string) (count int, vowel []rune) {
  count = 0
  for _, c := range inputWord {
    fmt.Printf("%c\n", c)
    // if c == 'a' || 'A' || 'e' || 'E' || 'i' || 'I' || 'o' || 'O' || 'u' || 'U' {
    //   count += 1
    //   vowel = [i]rune
    // }else {
    //   fmt.Println("No vowels found in the word")
    //   }
  }
  return count, vowel
}

func main(){
  word = "gopher"
  numOfVowels, listOfVowels := vowelFinder(word)
  fmt.Printf("Number of vowels appeared in the word \"%s\" is %d and those are %s\n", word, numOfVowels, listOfVowels)
}

















/*

*/
