/*
Program to find vowels and return the number of appearance from the provided string.
Program outline:
=> takes an arguments from the program.
  -> check only one argument is provided, error out if 2 arguments are provided
=> pass string to a function
  -> take first letter from the string and match with vowels
  -> if match found, store that chatacter in a slice increase the counter else move on to second character
  -> return slice and count to main function
=> display the result.
*/
package main

import (
	"fmt"
	"os"
)

// VowelsFinder finds appeared vowels,return them and their count
func VowelsFinder(word string ) (foundVowels []rune, count int) {
  fmt.Println("Your provided word is", word)
  // make an array of all vowels
  EngVowels := [5]rune{'a','e','i','o','u'}
  // take first letter from the word and compare it any one from EngVowels.
  count  = 0
  for i,_ := range word {
    for j, _ := range EngVowels {
      if rune(word[i]) == EngVowels[j] {
        foundVowels = append(foundVowels, rune(word[i]))
        count ++
      }
    }
  }
  return foundVowels,count
}

func main() {
	// getting argument from the user
	wordTaken := os.Args
	// check the length of the arguments
  if len(wordTaken) == 1 {
		fmt.Println("Provide a word to find the vowels")
		os.Exit(1)
	}
	if len(wordTaken) > 2 {
		fmt.Println("Only single argument is allowed")
		os.Exit(1)
	}

	// checks done: pass string to function
  Vowels, NumOfVowels :=	VowelsFinder(wordTaken[1])
  fmt.Println("Total number of vowels apreared in this word is", NumOfVowels)
  fmt.Printf("and that/those is/are[%s]\n", string(Vowels))
}
