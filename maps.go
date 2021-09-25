package main

import "fmt"


func main(){

	// Maps
	var m map [string] string

	fmt.Println(m == nil)
	var n = map [string]string{} //instatiation
	fmt.Println(n == nil) //false

	fmt.Println(m) // map[]
	fmt.Println(len(m))

	// instantiation(2nd method)
	k := make(map[string]string, 5)
	fmt.Println(len(k))

	k = map[string]string{ "name" : "wallace", "location" : "tz" }
	fmt.Println(k)
	k["skill"] = " Programmer"
	fmt.Println(k)

	for key, value := range k {
		fmt.Println(key, value)
	}

	title, ok := k["skill"]

	if ok {
		fmt.Println(title)
	} else{
		fmt.Println("skill doesn't exit")
	}
  
	name, _ := k["name"]
	fmt.Println(name)

}