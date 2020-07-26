package main

import (
	"fmt"
)

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	// Just enclose a character in single quotes 'A'
	// If no type is specified, Go will infer a rune.
	// The following three lines are equivalent:
	// grade := 'A'
	// var grade = 'A'
	// var grade rune = 'A'

	var star byte = '*'
	fmt.Printf("%c %[1]v\n", star)

	smile := '☺' // alt + 1
	fmt.Printf("%c %[1]v\n", smile)
}
