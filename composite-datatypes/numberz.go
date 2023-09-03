package main

import (
	"fmt"
	"math"
)


func main() {

	var x int = 3
	var p int = 1
	var x32 int32 = 2
	var x64 int64 = 3
	fmt.Println(x, x32, x64)

	var y float64 = 3.14
	fmt.Println(y)
	fmt.Printf("%T", y)

	var z float32 = 3.14
	fmt.Println(z)
	fmt.Printf("%T", z)

	// Arthmetics
	fmt.Println(x+p)
	fmt.Println(x-p)
	fmt.Println(x*p)
	fmt.Println(x/p)
	fmt.Println(x%p)
	fmt.Println(x+int(y))

	// MATH LIBRARY
	fmt.Println(math.Ceil(3.00001))
	fmt.Println(math.Floor(3.00001))
	fmt.Println(math.Min(3.00001, 6.012))
	fmt.Println(math.Max(3.00001, 6.012))
	fmt.Println(math.Abs(-1))
	fmt.Println(math.Sqrt(100))
	fmt.Println(math.Pow(3, 3)+1)

	// * Complex number
	fmt.Println(complex(1, 2))

}

