/*
Go program to show different data types used in golang. In this program different data types has been used and information provided on it.
Golang is a strongly typed language means the data type must be specified for the variable used in the program otherwise compiler will throw error.
Agenda in this program:
1. Use data types --> int , float , bool , string , array , slice
2. Check type of data
3. Convert between types of data.
Program is explained at the end.
*/

package main

import (
  "fmt"
  "reflect"
)
var (
  firstName string
  middleName rune
  lastName string
  age int
  height float64
  ownHome bool
  randNumber = [3]int{122, 319, 892}
  randString = []string{"This", "Is", "An", "Example", "1", "From", "Slice"}
)
func main() {
  // all variable have been defined but not assigned to value so start assigning it
  firstName = "Milind"
  lastName = "Dhoke"
  middleName = 'A'
  age = 32
  height = 5.9
  ownHome = false

  // print out all values
  fmt.Println("Hello! My name is", firstName, string(middleName), lastName) // Watch string(middleName)
  fmt.Println("My age is", age)
  fmt.Println("My height is", height)
  fmt.Println("Do I owe home here ? So it is", ownHome)
  fmt.Println("Printing Array:")
  for i := 0; i < len(randNumber); i++ {
    fmt.Println(randNumber[i])
  }
  fmt.Println("Printing Slice:")
  for j := 0; j < len(randString); j++ {
    fmt.Println(randString[j])
  }
  // check type of Data
  fmt.Println("firstName is of type", reflect.TypeOf(firstName))
  fmt.Println("middleName is of type", reflect.TypeOf(middleName))
  fmt.Println("lastName is of type", reflect.TypeOf(lastName))
  fmt.Println("age is of type", reflect.TypeOf(age))
  fmt.Println("height is of type", reflect.TypeOf(height))
  fmt.Println("ownHome is of type", reflect.TypeOf(ownHome))
  fmt.Println("randNumber is of type", reflect.TypeOf(randNumber))
  fmt.Println("randString is of type", reflect.TypeOf(randString))

}

/*
Explaination::
pakage main:
Every go executable program starts with main package. If program is not executable then it would be a part of golang library. In that case package main is
not necessary.

import :
import keyword is used to import the standard package from the go library. Here package fmt is being used since we are using some functions from this package like Println
function. If this package is not used then compiler will error out like undefined fmt. Try removing this line from the above code and rerun the program.

var:
Varible are decalred in this block. The scope of all these variable would be package level. That means these variable can be called at any part of the code
within this package. Varible assignment can be done here also. But in this program it is done in main function. Since golang is strongly typed language, variable must
be specified with data type otherwise compiler will error out.

Data Types:
Program needs data for porcessing. To understand this data correctly and to perform the equivalent operation on this data , compiler should know the type of data.
Data types used in golang are int, fload32 or float64, bool, array, slice.
int ==> Used to store integer data type. Depending on the processor it could be int32 or int64. But mostly int is used.
float32 or float64 ==> Data with decimal point can be stored either in fload32 or float64. Again it is dependent on the system architecture.
bool ==> True or false value can be referenced by bool data type.
array ==> fixed list of same type of data can be reference in array data type.
slice ==> List of different type of data can be referenced in slice data type

variables can be declared in other ways also. Below format goes only into main function , in main function compiler assigns data type during compilation.
firstName := "Milind"
middleName := 'A'
lastName := "Dhoke"
age := 32
height := 5.9
ownHome := false
randNumber := [3]int{122, 319, 892}
randString := []string{"This", "Is", "An", "Example", "1", "From", "Slice"}

if you run this program it would get below output if string(middleName) method is not used.
Hello! My name is Milind 65 Dhoke
My age is 32
My height is 5.9
Do I owe home here ? So it is false

If you see first line from output. Instead A it printed 65. Why ? Because A rune is a type meant to represent a Unicode code point. ASCII defines 128 characters,
identified by the code points 0–127. It covers English letters, Latin numbers, and a few other characters.

Whenever a rune is assined then its asosciated ASCII value will be stored to that variable. To convert this value to actual value then string() method can be used
which will convert it to string value.

Type Checking:
To find out the type of data we used in this program. Function called TypeOf is being used which finds the data type. But if you observe it closely. This
function is being called from a package named reflect. If this package is not imported then compiler will panic and throw an error.

This is enough information on data types as of now. Loops can not be explained since it will be explained in other section.
*/
