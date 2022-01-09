package main

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil

}
func main() {

	msg, err := Hello("Venu")
	if err == nil {
		fmt.Println(msg)
	} else {
		fmt.Println(err)
	}

}
