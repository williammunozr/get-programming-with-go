package main

import "fmt"

const hoursPerDay = 24

var (
	distance = 56000000 // km
	days     = 28
)

func main() {
	var speed = distance / hoursPerDay / days

	fmt.Println("Speed needed to travel in 28 days is", speed, "km/h")

	apendixSolution()
}

func apendixSolution() {
	fmt.Println(distance / (days * hoursPerDay))
}
