package helper

import "fmt"

func CausePanic() {
	fmt.Println("Panic!")
	panic("Something went wrong")
}
