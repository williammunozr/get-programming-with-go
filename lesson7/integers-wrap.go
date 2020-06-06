package main

import "fmt"

// Wrapping applies to 16-bit, 32-bit, and 64-bit integers
func main() {
	var red uint8 = 255
	red += 2
	fmt.Println(red) // Prints 1

	var number int8 = 127
	number += 3
	fmt.Println(number) // Prints -126

	red = 0
	red--
	fmt.Println(red) // Prints 255

	number = -128
	number--
	fmt.Println(number) // Prints 127

	var green uint16 = 65535
	green++
	fmt.Println(green) // Prints 0
}