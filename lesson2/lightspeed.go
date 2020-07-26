// How long does it take to get to Mars?

package main

import (
	"fmt"
	"math/rand"
)

var (
	distance = 56000000 // km
	speed    = 100800   // km/h
)

func main() {
	const lightSpeed = 299792 // km/2

	fmt.Println(distance/lightSpeed, "seconds")

	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	daysToMars()
	randomDistanceToMars()
}

func daysToMars() {
	const hoursPerDay = 24
	var distance = 96300000 // km

	fmt.Println(distance/speed/hoursPerDay, "days")
}

func randomDistanceToMars() {
	var distance = rand.Intn(345000001) + 56000000
	fmt.Println("Random distance to Mars", distance)
}
