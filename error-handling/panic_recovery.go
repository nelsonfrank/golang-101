package main

import "fmt"

// Panic is built-in func  throw unexpected in go
// Recover is build-in func catch panic exceptions
//
// Read more
// https://go.dev/blog/defer-panic-and-recover
func main() {
	defer safeExit()

	const one = 2

	if one != 1 {
		panic("One isn't 1, This isn't good")
	}

	
}

func  safeExit(){
	if r := recover(); r!= nil {
		fmt.Println("Panic is recovered!")
	}
}
