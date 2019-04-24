/*
Functions in Go
returning multiple values from the function
Sometime in programming we want function to return multiple values which will be used by the calling code block. To do so function can return multiple
values provided all values data types has been typed already in function signature.
Syntax for such function would be like:
func func_name (arguments data_types) (return_data_type, return_data_type) {
  ...
  ...
  ...
  return value1, value2
}

lets take an example:
objective:
Customer A went to a shop, he bought some fruits like, apples, oranges, kiwis etc. As a shop owner you should have a record stating what quatity of fruits has
been bought by customer A. What is the price of single type fruits and what is total sell value. We are not creating any record but just displaying on the terminal.
*/

package main

import (
  "fmt"
)

var (
  apples int
  oranges int
  kiwis int

  oneApple float64
  oneOrange float64
  oneKiwi float64

  appleCost, orangeCost, kiwiCost float64
)

func fruitCalculator(unoApple,unoOrange,unoKiwi float64, numApples, numOranges, numKiwis int)(float64, float64, float64) {
  appleCost = unoApple * float64(numApples)
  orangeCost = unoOrange * float64(numOranges)
  kiwiCost = unoKiwi * float64(numKiwis)
  //
  fmt.Printf("Customer has bought %d apples, %d oranges, and %d kiwis\n", numApples, numOranges, numKiwis)
  return appleCost, orangeCost, kiwiCost

}

func main(){
  // Assign per unit price to each apple, orange and kiwi
  oneApple = 1.75
  oneOrange = 1.24
  oneKiwi = 0.79

  apples = 7
  oranges = 3
  kiwis = 16

  fmt.Println("Finding the unit prices of all fruits, number of quatities bought by customer and final billing for all fruits")
  a, o, k := fruitCalculator(oneApple, oneOrange, oneKiwi, apples, oranges, kiwis)
  fmt.Printf("Total price of apples is %f, total price of oranges %f , total price of kiwis %f\n", a, o, k)
  totalSale := a + o + k
  fmt.Println("Total bill of the customer is:", totalSale)
}

/*
Expected output:
Finding the unit prices of all fruits, number of quatities bought by customer and final billing for all fruits
Customer has bought 7 apples, 3 oranges, and 16 kiwis
Total price of apples is 12.250000, total price of oranges 3.720000 , total price of kiwis 12.640000
Total bill of the customer is: 28.61

Unserstanding Program and concepts
program executions alwways starts from function main. So lets try to understand it from main only.
in var block varaible have been declared , and assignement happened in main. In main variables have been assingned to thire values.
next function is called and return value of this function is assigned to 3 variables a, o, and k.
==> Why assinged to 3 variables ?
    The function fruitCalculator returns 3 values to the caller block. if we assign less than or greater than 3 values then compiler will error out.
    assignment totally depends on the number of return values from the function.

When execution comes to line 59 , it takes controll to function fruitCalculator. execution starts in that function now.
Now we are in function fruitCalculator.
This functions take input arguments and process the logic to find the total cost of each fruit. Total cost for each type of fruit is stored in appleCost
orangeCost and kiwiCost.
Since we need these values to be passed to main function to calculate total sale. So this function should return these values. For that purpose, function is declared
in such a way that 3 float values are being return.

If you observe the code minutely. You can see addition statement where type conversion has been took place.
==> Why ?
    we cannot add, subtract, compare or perform any kind of operation on two different types in golang.
    Unlike other statically typed languages like C, C++, and Java, Go doesn’t provide any implicit type conversion. To make this work you will have explicitely cast
    to same data type so that operations can be performed on same data type variables.
*/
