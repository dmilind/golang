/*
Go program to find the area of geometry using pointer receiver to change the value.
*/

package main

import (
	"fmt"
)

const (
	pi float64 = 3.14
)

// Area struct collecting dimensions of geometry
type Area struct {
	length float64
}

// AreaofSquare method calculates area of square --> using value receiver
func (a Area) AreaofSquare() {
	area := a.length * a.length
	fmt.Println("Area of a square of side", a.length, "is:", area)
}

// PointerMethodArea method calculates area of square --> using pointer receiver
func (a *Area) PointerMethodArea(newLength float64) {
	fmt.Println("Old value of side of a square is", a.length)
	fmt.Println("New value of side of square is", newLength)
	a.length = newLength
	area := newLength * newLength
	fmt.Println("Area of square of side", newLength, "is", area)
}
func main() {
	squareOne := Area{
		length: 4,
	}
	squareOne.AreaofSquare()
	squareOne.PointerMethodArea(7)
}
