package main

import (
	"fmt"
	"strconv"
)

func main() {
	result, err := returnError(true) // change this to true to get error
	if err != nil {
		fmt.Println(err)
	} else{
		fmt.Println(result)
	}
}

// custom error handling method 
type specialError struct {
	errorMessage string
	errorCode int
}

func (s specialError) Error() string {
	return s.errorMessage + " " + strconv.Itoa(s.errorCode)
}

func returnError (returnError bool) (string, error){
	if returnError {
		return "", specialError{errorMessage: "Special Error", errorCode:  123}
	}

	return "Random result", nil
}