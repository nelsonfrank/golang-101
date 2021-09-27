package main

import "fmt"


type ninjaStar struct{
	owner string
}

func(ninjaStar) attack(){
	fmt.Println("Throwing ninja star")
}



type ninjaSword struct{
	owner string
}

func(ninjaSword) attack(){
	fmt.Println("Swinging ninja sword")
}


type ninjaWeapon interface {
	attack()
}

func attack(weapon ninjaWeapon){
	weapon.attack()
}

func main() {
	stars := []ninjaStar{{owner: "wallace"}, {owner: "wallace"}}

	for _, star := range stars {
		star.attack()
	}

	fmt.Println()

	swords := []ninjaSword{{owner: "wallace"}, {owner: "wallace"}}

	for _, sword := range swords {
		sword.attack()
	}

	fmt.Println()

	weapons := []ninjaWeapon{ninjaStar{owner: "wallace"}, ninjaSword{owner: "wallace"}}

	for _, weapon := range weapons {
		attack(weapon)
	}
}