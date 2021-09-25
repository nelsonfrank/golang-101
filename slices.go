package main

import "fmt"

func main() {
	fmt.Println("slices")

	// this is array type and has fixed value length
	fixed := [...]int{1, 2, 3}
	fmt.Println(fixed)
	// You can't assign it with array with length grater 
	// than origin value
	// fixed = [...]int{1, 2, 3, 4}

	// this is slices
	a := []int{1, 2, 3, 5}
	fmt.Println(a)
	// you can assign array with value more than initial slice
	a = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a)

	a = append(a, 8, 9, 10)
	fmt.Println(a)

	b := make([]int, 5)
	fmt.Println(b)


	// a slice of a slice
	fmt.Println(a[0:7]) // [1 2 3 4 5 6 7]
	fmt.Println(a[0:])  // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(a[:7])  // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(a[0:12])// [1 2 3 4 5 6 7 8 9 10 0 0] 

	// declaration
	var c []int
	fmt.Println(c == nil)
}