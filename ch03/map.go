/*
Map: Data Structure in Go
What is map ?
  We have seen array and slice before. In both of those data structure, values are retrived by the index value starting with 0. what if we want some key to
  the values we store in the array or slice. To do so go has data structure called map. map can be called as assosciative array or dictionary or hash table or hashmap.
  In the bottom map is a data structure to store key value pair like key = value

Syntax:
  var fruits = map[string] int {"string":"int", ...}
We can use make function to initialize map.
  var fruits  = make(map[string]int)

*/
package main

import (
  "fmt"
)

func main(){
  // initialize map for a team of barcelona with the position of the players
  var barcelonaTeam = map[string]string{"Marc":"GK", "Nelson":"DF", "Gerald":"DF", "Ivan":"MF", "Sergio":"MF", "John":"DF", "Philip":"MF", "Messi":"FW", "Louise":"Fw"}

  // initialize map using make function: map for Madrid team
  var madridTeam = make(map[string]string)
  madridTeam["Keylor"] = "GF"
  madridTeam["Dani"] = "DF"
  madridTeam["Jesus"] = "DF"
  madridTeam["Marianno"] = "FW"
  madridTeam["Toni"] = "MF"
  madridTeam["Karim"] = "FW"
  madridTeam["Luka"] = "MF"
  madridTeam["Mercelo"] = "MF"
  madridTeam["Lucas"] = "FW"

  //Now printing the map for barcelona team
  fmt.Println("Messi's game position is", barcelonaTeam["Messi"])
  fmt.Println("Luka's game position is", madridTeam["Luka"])

  fmt.Println("Barcelone team")
  // finding all team members with position from barcelona team
  for p, gp := range barcelonaTeam {
    fmt.Printf("%s => %s\n", p, gp)
  }
  fmt.Println("Madrid Team")
  // finding all team members with position from madrid team
  for p, gp := range madridTeam {
    fmt.Printf("%s => %s\n", p, gp)
  }


}
