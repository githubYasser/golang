package main

import (
	"errors"
	"fmt"
)

func main() {

	//create an array of strings
	names := []string{"Gladys", "Samantha", "Darrin"}
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Hello" + "," + message)

	}

}

// craete Hello function returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return name, nil
}

// craete a generic function that sum float and ints m
func Sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
