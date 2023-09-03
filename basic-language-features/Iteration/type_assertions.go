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
		ns, ok := weapon.(ninjaStar) // type assertion applied here
		if ok {
			ns.pickUp()
		}
	}
}