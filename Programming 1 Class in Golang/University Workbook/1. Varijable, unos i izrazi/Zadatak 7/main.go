package main

/*
	Cartesian coordinates x & y can be converted to polar coordinates r and phi via:

	r = sqrt(x^2 + y^2)
	phi = arctan(y / x)

	Write a program that requires the user to enter the values of the x & y Cartesian coordinates, and then print out
	the given values of r and phi of the polar coordinate system. The value of phi must not be printed out in the value
	of radians, but in degrees.
*/

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	fmt.Println("Enter the value of x:")
	fmt.Scan(&x)
	fmt.Println("Enter the value of y:")
	fmt.Scan(&y)

	r := math.Sqrt(x*x + y*y)
	phi := math.Atan2(y, x)

	degreesPhi := (180.0 * phi) / math.Pi
	fmt.Printf("r: %f\nphi: %f", r, degreesPhi)
}
