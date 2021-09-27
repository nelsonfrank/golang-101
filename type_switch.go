package main

import "fmt"


type ninjaStar struct{
	owner string
}

type ninjaSword struct{
	owner string
}


func(ninjaStar) attack(){
	fmt.Println("Throwing ninja star")
}

func(ninjaStar) pickUp(){
	fmt.Println("Picking back up ninja star")
}
func(ninjaSword) attack(){
	fmt.Println("Swinging ninja sword")
}


type ninjaWeapon interface {
	attack()
}
 
func main() {
	weapons := []ninjaWeapon{ninjaStar{owner: "wallace"}, ninjaSword{owner: "wallace"}}

	for _, weapon := range weapons {
		weapon.attack()

		switch weapon.(type) {
		case ninjaStar:
			ns := weapon.(ninjaStar)
			ns.pickUp()
		case ninjaSword:
			fmt.Println("No need to pick it back")
		}
	}
}