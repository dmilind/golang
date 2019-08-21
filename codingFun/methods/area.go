/*
Go program to find the area of geometry
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
	length  float64
	breadth float64
	height  float64
	radius  float64
	base    float64
}

// AreaofSquare method calculates area of square
func (a Area) AreaofSquare() {
	area := a.length * a.length
	fmt.Println("Area of a square of side", a.length, "is:", area)
}

// AreaofRectangle method calculates area of rectangle
func (a Area) AreaofRectangle() {
	area := a.length * a.breadth
	fmt.Println("Area of a rectangle of sides", a.length, a.breadth, "is:", area)
}

// AreaofTriagle method calculates area of triagle
func (a Area) AreaofTriagle() {
	area := 0.5 * (a.base * a.height)
	fmt.Println("Area of a triangle of base of height", a.base, a.height, "is:", area)
}

// AreaofCircle method calculates area of circle
func (a Area) AreaofCircle() {
	area := pi * a.radius * a.radius
	fmt.Println("Area of a circle with radius", a.radius, "is:", area)

}

func main() {
	squareOne := Area{
		length: 4,
	}

	rectangleOne := Area{
		length:  5,
		breadth: 7,
	}

	triagleOne := Area{
		base:   3.5,
		height: 2.9,
	}

	circleOne := Area{
		radius: 7.2,
	}

	// area of square
	squareOne.AreaofSquare()

	// area of a rectagle
	rectangleOne.AreaofRectangle()

	// area of triangle
	triagleOne.AreaofTriagle()

	// area of circle
	circleOne.AreaofCircle()
}
