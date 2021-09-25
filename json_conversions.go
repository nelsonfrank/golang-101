package main

import (
	"encoding/json"
	"fmt"
)

type Ninja struct{
	Name string `json:"full_name"` // you can specify name json should refers to
	Weapon Weapon
	Level int
}

type Weapon struct {
	Name string `json:"weapon_name"`
	Level string `json:"weapon_level"`
}
func main() {
	
	sFrom := `{"full_name": "Wallace", "weapon": {"weapon_name": "Ninja Star", "weapon_level":1}, "level": 1}`
	fmt.Println(sFrom)

	var wallace Ninja
	fmt.Println(wallace)
	err := json.Unmarshal([]byte(sFrom), &wallace)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(wallace)

	sTo, err := json.Marshal(wallace)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sTo)
	fmt.Printf("%T\n", sTo)
	fmt.Printf("%s\n", sTo)
}