package main

import "fmt"

func main() {
	fmt.Println("There is a cavern entrance here and a path to east.")

	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You need further up the mountain.")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
}