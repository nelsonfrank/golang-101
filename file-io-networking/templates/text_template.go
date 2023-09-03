package main

import (
	"fmt"
	"os"
	"text/template"
)

type secret struct {
	Name string
	Secret string
}

func main() {
	

	secret := secret{Name: "Nelson", Secret: "I'm a programmer"}

	templatePath := "/home/nelson/code_forest/studies/golang-101/template.txt"

	t, err := template.New("template.txt").ParseFiles(templatePath)

	if err != nil {
		fmt.Println(err)
	}

	err = t.Execute(os.Stdout, secret)

	if err != nil {
		fmt.Println(err)
	}

	//Read more about Text template
	// https://golang.com/pkg/text/template
}