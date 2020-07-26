package main

import "fmt"

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n", third)
	fmt.Printf("%f\n", third)
	fmt.Printf("%.3f\n", third)
	fmt.Printf("%4.2f\n", third)
	fmt.Printf("%05.2f\n", third)
}

// Prints
// 0.3333333333333333
// 0.3333333333333333
// 0.333333
// 0.333
// 0.33
// 00.33