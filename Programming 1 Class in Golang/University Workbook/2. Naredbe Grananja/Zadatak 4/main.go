package main

/*
	Write a program that asks the user to enter the x and y coordinates of a circle's centre and the circle's radius.
	The program then requests the coordinates of point A. It checks whether point A lies inside the circle, on the
	circle's circumference, or outside the circle, and prints the corresponding message:
	"The point lies inside the ricle",
	"The point lies on the circle",
	"The point lies outside the circle"
*/

import (
	"fmt"
	"math"
)

func main() {
	var x, y, R, xa, ya float64

	fmt.Println("Enter the coordinates of (x, y):")
	fmt.Scan(&x)
	fmt.Scan(&y)

	fmt.Println("Enter the radius of the circle:")
	fmt.Scan(&R)

	fmt.Println("Enter the coordinates of (xa, ya)")
	fmt.Scan(&xa)
	fmt.Scan(&ya)

	D := math.Sqrt(math.Pow(xa-x, 2) + math.Pow(ya-y, 2))

	if D < R*R {
		fmt.Println("The point lies inside the circle.")
	} else if D > R*R {
		fmt.Println("The point lies outside the cirlce")
	} else {
		fmt.Println("The point lies on the circle")
	}
}
