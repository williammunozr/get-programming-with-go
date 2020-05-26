package main

import "fmt"

func main() {
	var room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough // used to execute the body of the next case
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}
}