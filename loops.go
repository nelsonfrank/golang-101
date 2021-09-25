package main

import "fmt"

func main(){
	fmt.Println("Hello, World!")

	// For loop
	// 1
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//2
	for i, j := 0, 1; i < 12 && j<10; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	fmt.Println()
	// Foreach loop
	s:= "Hello, world"

	for i, c := range s {
		fmt.Printf("%d %c", i, c)
		fmt.Println()
	}

	// break
	for _, c := range s {
		if c == ' '{
			break
		}
		fmt.Printf(" %c",  c)
	}



	// continue
	for _, c := range s {
		if c == ' '{
			continue
		}
		fmt.Printf(" %c",  c)
	}
	fmt.Println()

	//label
	outerForLoop:
	for i := 0; i < 5; i++{
		for j := 0; j< 5; j++{
			if j == 3 {
				break outerForLoop
			}
			fmt.Printf("%d%d ", i, j)
		}
		fmt.Println()
	}
}