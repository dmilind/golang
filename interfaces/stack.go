/*

Using interfaces in go: Pushing and popping colored cubes from a stack

*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	bar []string
)

type Stack interface {
	push()
	pop()
}

type Redcube struct {
	cubenum int
	color   string
}

type Bluecube struct {
	cubenum int
	color   string
}

// method for pushing redcube
func (r Redcube) push() {
	fmt.Println("Pushing", r.cubenum, r.color, "colored cubes on to a stack")
	for i := 0; i < r.cubenum; i++ {
		bar = append(bar, r.color)
	}
	fmt.Println("Current cubes in stack:", bar)
}

// method for popping redcubes
func (r Redcube) pop() {
	// current state of the stack
	fmt.Println("Current cubes in the stack: -->", bar)
	fmt.Println("Popping", r.cubenum, r.color, "colored cubes from a stack")
}

// method for pushing blue cubes
func (r Bluecube) push() {
	fmt.Println("Pushing", r.cubenum, r.color, "colored cubes on to a stack")
	for i := 0; i < r.cubenum; i++ {
		bar = append(bar, r.color)
	}
	fmt.Println("Current cubes in stack:", bar)
}

// method for popping blue cube

// func Activity(s Stack, opt string) {
// 	// pushing activity
// 	switch opt {
// 	case "push":
// 		fmt.Println("------------------------------------")
// 		fmt.Println("Pushing operation on stack")
// 		s.push()
// 	case "pop":
// 		fmt.Println("------------------------------------")
// 		fmt.Println("Popping operation on stack")
// 		s.pop()
// 	default:
// 		fmt.Println("------------------------------------")
// 		fmt.Println("Choose either push or pop")
// 		os.Exit(1)
// 	}

// }

func main() {
	// get an empty stack
	fmt.Println("------------------------------------")
	fmt.Println("Current state of the stack is -->", bar)
	fmt.Println("------------------------------------")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose stack operation from push or pop")
	choice, _ := reader.ReadString('\n')
	switch choice {
	case "push":
		fmt.Println("this is push")
	default:
		return
	}
	// if choice == "push" {
	// 	fmt.Println("Enter color cube")
	// 	cubecolor, _ := reader.ReadString('\n')
	// 	fmt.Println("Enter number of", cubecolor, "cube")
	// 	num, _ := reader.ReadByte()
	// 	rc := Redcube{int(num), cubecolor}
	// 	Activity(rc, choice)
	// } else if choice == "pop" {
	// 	fmt.Println("Enter color cube")
	// 	cubecolor, _ := reader.ReadString('\n')
	// 	fmt.Println("Enter number of", cubecolor, "cube")
	// 	num, _ := reader.ReadByte()
	// 	rc := Redcube{int(num), cubecolor}
	// 	Activity(rc, choice)
	// }

	//bc := Bluecube{5, "BLUE"}

}
