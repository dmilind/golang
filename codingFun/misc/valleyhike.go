/*

 */

package main

import (
	"fmt"
)

var (
	currentPosition, downElevationCount, upElevationCount int
)

func findposition(elevation []rune) ([]rune, int, int) {
	var hikeTracked []rune
	for i, _ := range elevation {
		switch elevation[i] {
		case 'd':
			fmt.Println("Current elevation --> Down", "After", i+1, "hike")
			currentPosition -= 1
			downElevationCount += 1
			hikeTracked = append(hikeTracked, elevation[i])

		case 'u':
			fmt.Println("Current elevation --> Up", "After", i+1, "hike")
			currentPosition += 1
			upElevationCount += 1
			hikeTracked = append(hikeTracked, elevation[i])
		}
		if downElevationCount == upElevationCount {
			fmt.Println("Zero elevation after", i+1, "th hike")
		}
	}
	return hikeTracked, downElevationCount, upElevationCount

}

func main() {
	// initialize a slice of elevations of the hike
	hikeelevation := []rune{'d', 'd', 'u', 'd', 'd', 'u', 'd', 'u', 'u', 'u', 'u', 'd'}
	fmt.Println("Total hike with all elevations:", string(hikeelevation))
	ht, downs, ups := findposition(hikeelevation)
	fmt.Println("Total covered elevation: -->", string(ht))
	fmt.Println("Total [ downs: --> ", downs, "] [ ups: -->", ups, "] along the hike")
	if string(hikeelevation) == string(ht) {
		fmt.Println("Hiker hiked on correct trail")
	} else {
		fmt.Println("Hiker misguied on the trail")
	}
}
