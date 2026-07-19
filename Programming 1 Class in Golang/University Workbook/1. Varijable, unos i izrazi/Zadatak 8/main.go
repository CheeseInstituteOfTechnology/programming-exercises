package main

/*
	Using the values measured during an oblique shot, it is possible to calculate the acceleration of the
	Earth's gravity g using the formula:

	g = (v0^2 / R) * sin(2 * alfa)

	Where v0 is the initial velocity, R is the range, and alfa is the launch angle. Write a program that asks the user
	to enter the values of v0, R, and alfa in that order, and prints out the value of Earth's gravity via the formula
	written above.
*/

import (
	"fmt"
	"math"
)

func main() {
	var v0, R, alfa float64
	fmt.Println("Enter the initial velocity:")
	fmt.Scan(&v0)

	fmt.Println("Enter the range:")
	fmt.Scan(&R)

	fmt.Println("Enter the launch angle:")
	fmt.Scan(&alfa)

	radiansAlfa := alfa * (math.Pi / 180.0)

	g := (math.Pow(v0, 2) / R) * math.Sin(2*radiansAlfa)

	fmt.Println(g)
}
