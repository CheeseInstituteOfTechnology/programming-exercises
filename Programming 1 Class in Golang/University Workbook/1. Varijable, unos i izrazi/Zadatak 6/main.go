package main

/*
	Write a program that calculates the volume of a torus:

	V = 2 * PI^2 * R * r^2

	where R is the distance, and r the radius. The program requires the user to enter these two values.
*/

import (
	"fmt"
	"math"
)

func main() {
	var R, r float64

	fmt.Println("Enter the value of the distance:")
	fmt.Scan(&R)
	fmt.Println("Enter the value of the radius:")
	fmt.Scan(&r)
	V := 2 * math.Pi * math.Pi * R * r * r
	fmt.Println(V)
}
