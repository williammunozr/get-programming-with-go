package main

import (
	"fmt"
)

func main() {
	message := "shalom"

	c := message[5]
	fmt.Printf("%c\n", c)

	// exercise
	for i := 0; i < len(message); i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}
}
