package main

import "fmt"

func causePanic() {
	fmt.Println("Panic!")
	panic("Something went wrong")
}
