/*
This is a comment block
*/
// This is also a comment but in a single line

// every golang executable program start with package main
package main

// import the package from the stardard library
import (
  "fmt"
)

// program start executing from function main
func main() {
  fmt.Println ("Hello ! Gopher")
}

/*
save this program as hello.go and execute go run hello.go
Expected output:
Hello ! Gopher
*/
