package main

/*
	An ellipsoid is a closed central-symmetric surface of the second order. Three intersect in the centre of the
	ellipsoid mutually perpendicular axes (main axes) of symmetry. The ellipsoid can be parameterized in several ways,
	the most common way is:

	x = a * sin(theta) * cos(phi)
	y = b * cos(theta) * sin(phi)
	z = c * sin(theta)

	Write a program that asks the user for the values of a, b, c, theta, and phi, and therefore the program calculates
	and prints out the values of the x, y, and z coordinates.
*/

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, theta, phi float64
	fmt.Println("Please enter the next values of the variables:")
	fmt.Println("a:")
	fmt.Scan(&a)
	fmt.Println("b:")
	fmt.Scan(&b)
	fmt.Println("c:")
	fmt.Scan(&c)
	fmt.Println("Theta:")
	fmt.Scan(&theta)
	fmt.Println("Phi:")
	fmt.Scan(&phi)

	thetaRadians := theta * (math.Pi / 180)
	phiRadians := phi * (math.Pi / 180)

	x := a * math.Sin(thetaRadians) * math.Cos(phiRadians)
	y := b * math.Cos(thetaRadians) * math.Sin(phiRadians)
	z := c * math.Sin(thetaRadians)

	fmt.Printf("x: %f\ny: %f\nz: %f", x, y, z)
}
