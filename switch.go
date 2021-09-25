package main

import "fmt"

func main() {

	venues := []string{"Home", "Work", "School", "Gym"}

	for _, venue := range venues{
		switch venue {
		case "Home":
			greeting("Mom, I'm hungry")
		case "Work", "School":
			greeting("Weather is great today")
		default:
			greeting("sup, bois")
		}
	}
	
}

func greeting(greeting string){
	fmt.Println(greeting) 
}