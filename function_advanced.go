package main

import "fmt"

func main() {
	fmt.Println("Function Advanced")
	fmt.Printf("%T \n", returnParams) // func(int) int 
	fmt.Printf("%T \n", returnParams(2)) // int
	fmt.Printf("%d \n", returnFuntion(2, returnParams)) // 2


	defer first() // this function will be executed lastly in main func.
	second()
	second()

	aFunction :=  func ()  {
		fmt.Println("aFunction")
	}

	aFunction()
	

	func ()  {
		fmt.Println("aFunction 2")
	}()

	aFunction = first

	aFunction()

	// aFunction = returnParams // this will not work since returnParams pass params but aFunction doesn't

	crazyishFunction := func(f func()) func() {
		return f
	}

	crazyFunction(crazyishFunction)(second)()
}

func crazyFunction(f func(func()) func()) func(func()) func() {
	return f
}

func returnParams(i int) int{
	return i
}

func returnFuntion(i int , f func(int)int) int{
	return f(i)
}

func first()  {
	fmt.Println("first")
}

func second()  {
	fmt.Println("second")
}