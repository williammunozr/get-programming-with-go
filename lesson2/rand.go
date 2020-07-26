package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1 // add 1 because it is calculating number from 0-9
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
