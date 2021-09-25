package main

import (
	"errors"
	"fmt"
)

func main() {
	greeting() // "Hello, World!"

	variadic(1, 2, 3, 4) // 1 2 3 4

	lists := []int{1, 2, 3, 4}
	variadic(lists...)  // 1 2 3 4

	var sum int = sum(1, 2)
	fmt.Println(sum)

	 name, size := nameLength("nelson")
	 fmt.Println(name, size)

	result, error := sayHi("Nelson")
	if error != nil {
		fmt.Println(error)
	} else{
		fmt.Println(result)
	}

	result, ok := sayHiOk("Nelson")
	if ok {
		fmt.Println(result)
	} 

	// reccussion func
	var fibonacci int = fibonacci(10)
	fmt.Println(fibonacci)

	// variables are passed by value not by reference
	i := 1
	fmt.Println(i)

	// This will not change the value of i 
	// since variable passed as a copy of value of i
	withoutPointer(i)
	fmt.Println(i)

	// Way around update previous defined value using pointer
	// alternative to variable by reference
	withPointer(&i)
	fmt.Println(i)

}


func greeting(){
	fmt.Println("Hello, World!")
}

// function overload is not valid in go
//
// Below function doesn't going to work, 
// since function should have different name, 
// since golang doesn't support function overload
//
// func greetings(name string){
// 	fmt.Println("Hello, ", name)
// }

func variadic (values ...int){
	for _, value := range values {
		fmt.Println(value)
	}
}

func sum(a int, b int) int {
	return a + b
}
// We can also do this
// func sum(a , b int) int {}


// return multiple data type
func nameLength(name string)(string, int){
	return name, len(name)
}

// Error handling in function
func sayHi(name string)(string, error){
	if len(name) ==0 {
		return "", errors.New("Empty name!, This is not good")
	}

	return "Hi " + name, nil
}

func sayHiOk(name string)(string, bool){
	if len(name) ==0 {
		return "", false
	}

	return "Hi " + name, true 
}

// recursive function

func fibonacci(i int)int{
	if i == 0 {
		return 0
	} else if i == 1 {
		return 1
	} else {
		return fibonacci(i-1) + fibonacci(i-2)
	}
}

func withoutPointer(i int){
	i = i +1
}

func withPointer(i *int){
	*i = *i + 1
}