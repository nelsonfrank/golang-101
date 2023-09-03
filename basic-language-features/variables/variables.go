// Variable and Types Declaration
package main

import "fmt"

func main() {
   // var name type = expression
   var cars int = 1
   fmt.Println(cars)

   // declare multiple variable in single line
   // 1st 
   var num_of_people, num_of_room int
   fmt.Println(num_of_people, num_of_room)

	// 2nd
   var num_of_car, num_of_tractor, location = 1, 2, "Tanzania"
   fmt.Println(num_of_car, num_of_tractor, location)


   // Without assigning value to variable
   // it return the following
   // 0 if variable type is int
   // "" if varible type is string
   // false if variable type is boolean (bool)
   var houses int 
   fmt.Println(houses)

   // Short declaration
   // name := expression
   isPresent := false
   isLogin := true
   fmt.Println(isPresent, isLogin)


   // ## POINTER ##
   // pointer declaration

   // short declaration
   x := 1
   p := &x

   fmt.Println(p) // return address of variable x
   fmt.Println(*p) // return content of variable x

   // full declaration with types
   var i int = 2
   var h *int = &i

   fmt.Println(h)
   fmt.Println(*h)


   // ## Type Declaration ##
   type fahrenheit int
   type celsius int

   
   var f fahrenheit = 32
   var c celsius = 0
   
   // Now f is not equivalent to c 
   // since they have different type
   // example, f = c, this will lead to error below
   // Error: cannot use c (variable of type celsius) 
   // as fahrenheit value in assignment. compiler(IncompatibleAssign)

   fmt.Println(f, c)

}