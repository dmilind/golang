/*
Implementing bubble sort algorithm in go.
Algorithm:
In a nutshell bubble sort algorithm compares two adjacent elements from a list and swaps them if they are not in correct order. This comparison continues repeatedly until all
element are sorted properly.
Pseudo code:
procedure buublesort(dataset: Array of unsorted elements)
	for i= 0 to i= length(A) do
		for i = 0 to i < length(A)-1 do
			if A[i] > A[i+1]
				swap A[i] and A[i+1]
			endif
		endfor
	endfor
endprocedure
*/
package main

import (
	"fmt"
)

var (
	dataset = []int{406, 140, 727, 96, 719, 563, 132, 510, 698, 196}
)

func bubblesort(data []int) {
	pass := 0
	swap := 0
	for i := 0; i < len(dataset); i++ {
		for j := 0; j < len(dataset)-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swap++
			} else {
				pass++
			}

		}
	}
	fmt.Println("Total Swaps -->", swap)
	fmt.Println("Total pass -->", pass)
}
func main() {
	fmt.Println("Original unsorted data is -->", dataset)
	fmt.Println("-----------------------------------------")
	bubblesort(dataset)
	fmt.Println("-----------------------------------------")
	fmt.Println("Sorted data set is -->", dataset)
}
