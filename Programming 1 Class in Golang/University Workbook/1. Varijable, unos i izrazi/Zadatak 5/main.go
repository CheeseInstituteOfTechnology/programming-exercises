package main

/*
	Write a program that asks the user to enter the length of two sides of a triangle (b & c) and the angle between them
	(alfa).The program has to calculate and print out the length of the remaining, third side a. You can use the formula:
	a^2 = b^2 + c^2 - 2bccos(alfa)
*/

import (
	"fmt"
	"math"
)

func main() {
	var b, c, alfa float64
	fmt.Println("Enter the length of side b:")
	fmt.Scan(&b)
	fmt.Println("Enter the length of side c:")
	fmt.Scan(&c)
	fmt.Println("Enter the angle alfa:")
	fmt.Scan(&alfa)

	radiansAlfa := alfa * (math.Pi / 180)
	a := math.Sqrt(b*b + c*c - 2*b*c*math.Cos(radiansAlfa))
	fmt.Println(a)
}
