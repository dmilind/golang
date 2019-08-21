/*
[WIP]
Program to find min and max from the provided command line argument
code-logic:
1. Pass command line arguments while running the code.
   check: If no argument is passed then error out also check minimum 2 numbers required
2. Numbers are stored in Arguments. Pass argumets from index position 1 to MinMaxFinder function.
3.
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

var (
	min int
)

func MinMaxFinder(numberPassed []string) {
	var numList []int
	for i, _ := range numberPassed {
		num, err := strconv.Atoi(numberPassed[i])
		if err != nil {
			log.Fatal(err)
		}
		// append num to slice
		numList = append(numList, num)
	}
	// all good till this point
	// now compare each indexed number with other indexed number
	for i, _ := range numList {
		for j := 0; j < len(numList); j++ {
			if numList[i] < numList[j] {
				min = numList[i]
			}
		}
	}
	// if numOne < numTwo {
	//   fmt.Println("Minimum number is", numOne)
	// } else {
	//   fmt.Println("Minimum number is", numTwo)
	// }
	fmt.Println("Minimum number is", min)
}
func main() {
	//1. Take arguments from command line
	Arguments := os.Args
	// 1. Check
	if len(Arguments) == 1 {
		fmt.Println("No number have been provided! Provide numbers for calculations")
		os.Exit(1)
	} else if len(Arguments) < 3 {
		fmt.Println("You provided single number, but for finding minimum and maximum provided atleast 2 numbers")
		os.Exit(1)
	}
	MinMaxFinder(Arguments[1:])

}
