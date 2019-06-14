/*
Objective: Write a Program to find the sum of all numeric command line arguments:
Objective: Write a Go program that finds the average value of all of its float command-line arguments

Exercise question from: Mastering Go book
*/
package main

import (
  "fmt"
  "os"
  "strconv"
)
var (
  numInt []float64
)
func Sum(num []float64) float64 {
  val := 0.0
  for i := 0; i < len(num); i++ {
    val = val + num[i]
  }
  return val
}
func Average(sum float64, num []float64) float64 {
  div := float64(len(num))
  average := sum / div
  return average
}
func main(){
  argumentNumber := os.Args
  if len(argumentNumber) == 1 {
    fmt.Println("Provide me a number or series of numbers")
    os.Exit(1)
  }
  for i := 1; i < len(argumentNumber); i++ {
    k, _ := strconv.ParseFloat(argumentNumber[i], 64)
    numInt = append(numInt, k)
  }
  returnSum := Sum(numInt)
  fmt.Println("Sum of the provided numbers is", returnSum)
  returnAverage := Average(returnSum, numInt)
  fmt.Println("Average of the provided number is", returnAverage)
}
