/*
program to understand the concept of pointers and ampersands.
Also can be called as call by value and call by reference.
Find the explaination of concept at the end of program.
*/
package main

import (
  "fmt"
)

var (
  value string
  reference string
)

func main (){
  value = "ptr"
  reference = "ref"

  // Printing the actual values : call by value
  fmt.Println("Under call by value, value is:",value)
  fmt.Println("Under call by value, value is", reference)

 // printing the value through memory reference: call by reference
 fmt.Println("Under call by reference, value is:", &value)
 fmt.Println("Under call by reference, value is:",&reference)

 // Understanding pointer
 var p *string// declaring pointer variable
 randomText := "Understanding pointer"
 p = &randomText
 fmt.Println("randomText is located at", p) // This will return the memory address location
 fmt.Println("value of randomText from memory location is", *p) // this will return the value stored at memory address p

 // try to change the value of variable
 randomText = "Changed Value"
 fmt.Println("Value of changed variable from memory location is", *p) // Value go changed from memory.

}

/*
Explaination:
1.When variable are declared or defined, memory address has been allocated to these variables. So this storage space can be referenced by memory address.
And in that memory address the value of that variable is stored.

When we use the variable directly using declared variable names , that means we are calling the variable using call by value concept.
When we derive the value of object from storage location , this concept is called as call by reference.

What is pointer ?
concept of pointer comes from c language. A pointer points to the memory address of a variable. Consider a pointer as an entity which stores the memory address
value of the variable.

Why do we need pointers in programming ?
Some times in programming you dont find a data type or data structure to retrieve the value , in that case you can use pointers.

In above program. from line 29 onwards.
pointers can be declared as var variable *type. Pointer is a special data type which is declared by variable name followed *data_type
Here we initialized a variable randomText with a string value "Understanding pointer"
Think that this value is located in memory cell. Now this memory cell has some address. To get this address we need to use reference character which is &.
in next line , the memory address value is stored in the pointer. So pointer p is now pointing to the memory address of the randomText.
If we print p , then memory address value will be printed.

Now how can we get the value from the location. So here * acts as dereference character also. dereference parameter will find the value from the location.
In next line we tried to get the value from the location by using dereference character *.

In next line we changed the value of randomText variable. 
*/
