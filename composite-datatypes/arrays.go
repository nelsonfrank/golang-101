package main

import "fmt"


func main(){

	fmt.Println("Hello, world")

	var arr [5] int

	arr[0] = 1
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(len(arr))

	arrCopy := arr
	arrCopy[0] = 2

	fmt.Println(arr) // [1, 0, 0, 0, 0]
	fmt.Println(arrCopy) // [2, 0, 0, 0, 0]
	fmt.Println(arrCopy == arr) // false

	arrNum := [4]int{1, 2, 3, 4}
	fmt.Println(arrNum) // [1, 2, 3, 4]

	// Since array are iterational DS we can use with for loop
	for _, v:= range arrNum{
		fmt.Println(v)
	}
	fmt.Println()

	// specify array without specify size
	array := [...]int{1, 2, 3, 4}
	fmt.Println(array) // [1, 2, 3, 4]

	// assign value to specific index, the unspecified will be 0
	arrValue := [...]int{1:3, 4:5}
	fmt.Println(arrValue) // [0, 1, 0, 0, 5]

	array1 := [...]int{0:1}
	fmt.Println(array1) // [1]

	// string array
	array2 := [...]string{"Hello", "World"}
	fmt.Println(array2) // ["Hello", "World"]
	 

	array2d := [2][2]int{{1, 2} ,{3, 4}}
	fmt.Println(array2d) // [[1 2] [3 4]]

	array3d := [2][2][2]int{array2d , array2d}
	fmt.Println(array3d) // [[[1 2] [3 4]] [[1 2] [3 4]]]
}