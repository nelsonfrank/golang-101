package main

import "fmt"

func main(){
	fmt.Println(" Hello, Wolrd")

	// Struct
	// Declaration
	type ninja struct {
		name string
		weapons []string
		level int
	}

 
	wallace := ninja { name: "Wallace"}

	wallace = ninja {
		name: "wallace",
		weapons: []string{"Ninja star"},
		level: 1,
	}
	fmt.Println(wallace)
	fmt.Println(wallace.name)
	fmt.Println(wallace.weapons)
	fmt.Println(wallace.level)
	wallace.level++
	wallace.weapons = append(wallace.weapons, "Ninja Sword")

	fmt.Println(wallace)
	fmt.Println(wallace.weapons)


	type dojo struct {
		name string
		ninja ninja
	}

	golonaDojo := dojo{
		name: "Wallace",
		ninja: wallace,
	}

	fmt.Println(golonaDojo)

	type addressDojo struct {
		name string
		ninja *ninja
	}

	addressed := addressDojo{
		name: "Address Dojo", 
	    ninja: &wallace,
    }

	fmt.Println(addressed)
	fmt.Println(*addressed.ninja)
	addressed.ninja.level = 4
	fmt.Println(addressed.ninja.level)
	fmt.Println(wallace.level)

	// other way to assign new struct
	johnPointer := new(ninja)
	fmt.Println(johnPointer)
	fmt.Println(*johnPointer)
	johnPointer.name = "Johny"
	johnPointer.weapons = []string{"Ninja star"}
	johnPointer.level = 1

	fmt.Println(*johnPointer)

	intern := ninjaIntern {name: "Intern", level: 1}
	fmt.Println(intern)
	intern.sayHello()
	intern.sayName()

	} 

	// you can declare struct outside main func and call it into main func
	type ninjaIntern struct {
		name string
		level int
	}

	// Add function in struct
	func (ninjaIntern) sayHello()  {
		fmt.Println("Hello, Ninjas")
	}

	// Add function in struct and access struct properties
	func (n ninjaIntern)sayName()  {
		fmt.Println(n.name)
	}