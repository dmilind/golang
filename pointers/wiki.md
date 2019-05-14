#### Understanding Call by Value and Call by Reference

When a function is called , some argument will be injected to that function. By using this function performs some tasks and returns the final output. Now when arguments are passed to function, those arguments can be direct substitution of variable or indirect substitution of variable. Direct substitution means actual face value will be passed to the function, indirect substitution means address of that variable is passed to the function. This behavior is called as call by value and call by reference resp.

Call By Value:

Consider below snippet of the code:

func example (var int) int {
  var = 3
  return var
}

func main(){
  param := 1
  faceValue := example(param)
}

This is an example of call by value function. In func main , param value is set to 1, so actual face value of this variable is 1. When assignment is done memory is allocated to param, for example lets say the memory address allocated to this variable is 0001. Now this value is being passed to func example, this value is loaded
into variable var. Since variable is declared here so memory will be allocated at this step also. This memory address will be different than that of param. Lets say variable var got memory address of 1010. In function example, var value is assigned to 3, that operation performed on memory address of var. Address of param is untouched. At this address param value is still saved. When functino example return this value and this value saves into faceValue , so value would be 3. 
