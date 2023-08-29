package main

import (
	"fmt"

	"example/helper"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start")

	// call the causePanic function from the same package
	causePanic()

	// Call the CausePanic function from the helper package
	helper.CausePanic()

	fmt.Println("End")

}
